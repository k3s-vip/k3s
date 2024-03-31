package spegel

import (
	"context"
	"errors"
	"io"
	"sync"

	"github.com/opencontainers/go-digest"
	"github.com/spegel-org/spegel/pkg/oci"
	"github.com/spegel-org/spegel/pkg/oci/containerd"
	"github.com/spegel-org/spegel/pkg/store"
)

var errStoreNotStarted = errors.New("deferred OCI Store not started")

type backend interface {
	oci.ImageLister
	store.Provider
	store.Watcher
	io.Closer
}

// DeferredStore extends `backend` by allowing for delayed connection
// to the backend. Functions are expected to return errors until the store
// has been started.
type DeferredStore interface {
	backend
	Start() error
}

// explicit interface check
var _ DeferredStore = &deferredStore{}

type deferredStore struct {
	mu     sync.RWMutex
	store  backend
	create func() (backend, error)
}

func (ds *deferredStore) get() backend {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	return ds.store
}

func (ds *deferredStore) Name() string {
	if b := ds.get(); b != nil {
		return b.Name()
	}
	return "deferred"
}

func (ds *deferredStore) ListImages(ctx context.Context) ([]oci.Image, error) {
	if b := ds.get(); b != nil {
		return b.ListImages(ctx)
	}
	return nil, errStoreNotStarted
}

func (ds *deferredStore) Resolve(ctx context.Context, ref string) (digest.Digest, error) {
	if b := ds.get(); b != nil {
		return b.Resolve(ctx, ref)
	}
	return "", errStoreNotStarted
}

func (ds *deferredStore) Descriptor(ctx context.Context, dgst digest.Digest) (store.Descriptor, error) {
	if b := ds.get(); b != nil {
		return b.Descriptor(ctx, dgst)
	}
	return store.Descriptor{}, errStoreNotStarted
}

func (ds *deferredStore) Open(ctx context.Context, dgst digest.Digest) (io.ReadSeekCloser, error) {
	if b := ds.get(); b != nil {
		return b.Open(ctx, dgst)
	}
	return nil, errStoreNotStarted
}

func (ds *deferredStore) Watch(ctx context.Context) ([]store.Event, <-chan store.Event, error) {
	if b := ds.get(); b != nil {
		return b.Watch(ctx)
	}
	return nil, nil, errStoreNotStarted
}

// Close is not part of the Store interface, but probably should be. Containerd implements it.
func (ds *deferredStore) Close() error {
	ds.mu.Lock()
	b := ds.store
	ds.store = nil
	ds.mu.Unlock()
	if b == nil {
		return nil
	}
	return b.Close()
}

// Start is called to actuall make a connection to the store backend
func (ds *deferredStore) Start() error {
	if ds.get() != nil {
		return nil
	}
	b, err := ds.create()
	if err != nil {
		return err
	}
	ds.mu.Lock()
	defer ds.mu.Unlock()
	if ds.store != nil {
		return b.Close()
	}
	ds.store = b
	return nil
}

// NewDeferredContainerd creates a deferred store that creates a new Containerd store when Start is called.
func NewDeferredContainerd(ctx context.Context, sock, namespace string, opts ...containerd.ContainerdOption) (DeferredStore, error) {
	return &deferredStore{
		create: func() (backend, error) {
			return containerd.NewContainerd(ctx, sock, namespace, opts...)
		},
	}, nil
}
