package k8s

import (
	"context"
	"io"
	"time"

	"github.com/docker/distribution/reference"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"

	"github.com/windmilleng/tilt/internal/container"
	"github.com/windmilleng/tilt/pkg/model"
)

var _ Client = &explodingClient{}

type explodingClient struct {
	err error
}

func (ec *explodingClient) Upsert(ctx context.Context, entities []K8sEntity) ([]K8sEntity, error) {
	return nil, errors.Wrap(ec.err, "could not set up k8s client")
}

func (ec *explodingClient) Delete(ctx context.Context, entities []K8sEntity) error {
	return errors.Wrap(ec.err, "could not set up k8s client")
}

func (ec *explodingClient) GetByReference(ctx context.Context, ref v1.ObjectReference) (K8sEntity, error) {
	return K8sEntity{}, errors.Wrap(ec.err, "could not set up k8s client")
}

func (ec *explodingClient) PodsWithImage(ctx context.Context, image reference.NamedTagged, n Namespace, lp []model.LabelPair) ([]v1.Pod, error) {
	return nil, errors.Wrap(ec.err, "could not set up k8s client")
}

func (ec *explodingClient) PollForPodsWithImage(ctx context.Context, image reference.NamedTagged, n Namespace, lp []model.LabelPair, timeout time.Duration) ([]v1.Pod, error) {
	return nil, errors.Wrap(ec.err, "could not set up k8s client")
}

func (ec *explodingClient) PodByID(ctx context.Context, podID PodID, n Namespace) (*v1.Pod, error) {
	return nil, errors.Wrap(ec.err, "could not set up k8s client")
}

func (ec *explodingClient) WatchPod(ctx context.Context, pod *v1.Pod) (watch.Interface, error) {
	return nil, errors.Wrap(ec.err, "could not set up k8s client")
}

func (ec *explodingClient) ContainerLogs(ctx context.Context, podID PodID, cName container.Name, n Namespace, startTime time.Time) (io.ReadCloser, error) {
	return nil, errors.Wrap(ec.err, "could not set up k8s client")
}

func (ec *explodingClient) CreatePortForwarder(ctx context.Context, namespace Namespace, podID PodID, optionalLocalPort, remotePort int, host string) (PortForwarder, error) {
	return nil, errors.Wrap(ec.err, "could not set up k8s client")
}

func (ec *explodingClient) WatchResource(ctx context.Context, gvr schema.GroupVersionResource, lps labels.Selector) (<-chan interface{}, error) {
	return nil, errors.Wrap(ec.err, "could not set up k8s client")
}

func (ec *explodingClient) ConnectedToCluster(ctx context.Context) error {
	return errors.Wrap(ec.err, "could not set up k8s client")
}

func (ec *explodingClient) ContainerRuntime(ctx context.Context) container.Runtime {
	return container.RuntimeUnknown
}

func (ec *explodingClient) PrivateRegistry(ctx context.Context) container.Registry {
	return container.Registry{}
}

func (ec *explodingClient) Exec(ctx context.Context, podID PodID, cName container.Name, n Namespace, cmd []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	return errors.Wrap(ec.err, "could not set up k8s client")
}
