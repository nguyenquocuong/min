package min

type options struct {
	addr        string
	isWebsocket bool
}

func Listen() error {
	s := newServer()

	if err := s.listen(); err != nil {
		return err
	}
	return nil
}

func ListenWithAddr(addr string, opts ...Option) error {
	s := newServer()

	s.opts.addr = addr

	for _, opt := range opts {
		opt(&s.opts)
	}

	if err := s.listen(); err != nil {
		return err
	}
	return nil
}
