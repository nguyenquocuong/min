package min

func WithWebsocket() Option {
  return func(o *options) {
    o.isWebsocket = true
  }
}
