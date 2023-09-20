package rest

type Interface interface {
  GetRateLimiter() flowcontrol.RateLimiter
  Verb(verb string) *Request
  Post() *Request
  Put() *Request
  Patch(pt types.PatchType) *Request
  Get() *Request
  Delete() *Request
  APIVersion() schema.GroupVersion
}

type RESTClient struct {
  base *url.URL

  versionedAPIPath string

  content ClientContentConfig

  createBackoffMgr func() BackoffManager

  rateLimiter flowcontrol.RateLimiter

  waringHandler WaringHandler

  Client *http.Client
}
