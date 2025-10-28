package main

import (
	"os"

	v1 "github.com/k3s-io/k3s/pkg/apis/k3s.cattle.io/v1"
	controllergen "github.com/rancher/wrangler/v3/pkg/controller-gen"
	"github.com/rancher/wrangler/v3/pkg/controller-gen/args"
)

var (
	basePackage = "github.com/k3s-io/k3s/types"
)

func main() {
	os.Unsetenv("GOPATH")
	controllergen.Run(args.Options{
		OutputPackage: "github.com/k3s-io/k3s/pkg/generated",
		Boilerplate:   "scripts/boilerplate.go.txt",
		Groups: map[string]args.Group{
			"k3s.cattle.io": {
				Types: []interface{}{
					v1.Addon{},
					v1.ETCDSnapshotFile{},
				},
				GenerateTypes:   true,
				GenerateClients: true,
			},
		},
	})
}
