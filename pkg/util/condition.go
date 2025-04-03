package util

import (
	"context"
	"encoding/json"
	"time"

	"github.com/k3s-io/k3s/pkg/daemons/config"
	"github.com/k3s-io/k3s/pkg/util/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	toolscache "k8s.io/client-go/tools/cache"
	toolswatch "k8s.io/client-go/tools/watch"
)

// GetNodeCondition extracts the provided condition from the given status and returns that.
// Returns nil and -1 if the condition is not present, and the index of the located condition.
func GetNodeCondition(status *corev1.NodeStatus, conditionType corev1.NodeConditionType) (int, *corev1.NodeCondition) {
	if status == nil {
		return -1, nil
	}
	for i := range status.Conditions {
		if status.Conditions[i].Type == conditionType {
			return i, &status.Conditions[i]
		}
	}
	return -1, nil
}

// SetNodeCondition updates specific node condition with patch operation.
func SetNodeCondition(core config.CoreFactory, nodeName string, condition corev1.NodeCondition) error {
	if core == nil {
		return ErrCoreNotReady
	}
	condition.LastHeartbeatTime = metav1.NewTime(time.Now())
	patch, err := json.Marshal(map[string]any{
		"status": map[string]any{
			"conditions": []corev1.NodeCondition{condition},
		},
	})
	if err != nil {
		return err
	}
	_, err = core.Core().V1().Node().Patch(nodeName, types.StrategicMergePatchType, patch, "status")
	return err
}

type NodeConditionFunc func(node *corev1.Node) (bool, error)

func WaitForNode(ctx context.Context, client kubernetes.Interface, nodeName string, condition NodeConditionFunc) error {
	selector := fields.Everything()
	if nodeName != "" {
		selector = fields.OneTermEqualSelector(metav1.ObjectNameField, nodeName)
	}
	lw := toolscache.NewListWatchFromClient(client.CoreV1().RESTClient(), "nodes", metav1.NamespaceNone, selector)
	return waitForNode(ctx, lw, condition)
}

func waitForNode(ctx context.Context, lw toolscache.ListerWatcher, condition NodeConditionFunc) error {
	until := func(ev watch.Event) (bool, error) {
		node, ok := ev.Object.(*corev1.Node)
		if !ok {
			return false, errors.New("event object not of type v1.Node")
		}
		return condition(node)
	}
	_, err := toolswatch.UntilWithSync(ctx, lw, &corev1.Node{}, nil, until)
	return err
}
