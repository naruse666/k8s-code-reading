package v1

import (
  rest "k8s.io/client-go/rest"
)

type CoreV1Interface interface {
  RESTClient() rest.Interface
}

type CoreV1Client struct {
  restClient rest.Interface
}

func (c *CoreV1Client) Events(namespace string) EventInterface {
  return newEvents(c, namespace)
}

func (c *CoreV1Client) RESTClient() rest.Interface {
  if c == nil {
    return nil
  }
  return c.restClient
}
