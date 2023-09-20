package kubernetes

import (
  corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type Interface interface {
  CoreV1() corev1.CoreV1Interface
}

type Clientset struct {
  corev1 *coreV1.CoreV1Client
}

func (c *Clientset) CoreV1() corev1.CoreV1Interface {
  return c.coreV1
}
