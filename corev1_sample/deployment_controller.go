package deployment

import (
  v1core "k8s.io/client-go/kubernetes/typed/core/v1"
)

type DeploymentController struct {
  rsControl controller.RSControlInterface
  client clientset.Interface

  eventBroadcaster record.EventBroadcaster

  dLister applisters.DeploymentLister
  
  dListerSynced cache.InformerSynced

  queue workqueue.RateLimitingInterface
}

func (dc *DeploymentController) Run(ctx context.Context, workers int) {
  // dc.client.CoreV1().Events("")の関係を追っていく
  dc.eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: dc.client.CoreV1().Events("")}) 

  <-ctx.Done()
}
