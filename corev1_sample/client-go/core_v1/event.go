package v1

import (
  rest "k8s.io/client-go/rest"
)

type EventsGetter interface {
  Events(namespace string) EventInterface
}

type EventInterface interface {
  Create(ctx context.Context, event *v1.Event, opts metav1.CreateOptions) (*v1.Event, error)
}

type events struct {
  client rest.Interface
  ns string
}

func newEvents(c *CoreV1Client, namespace string) *events {
  return &events{
    client: c.RESTClient(),
    ns: namespace,
  }
}
