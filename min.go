package min

import (
	"fmt"
	"net"
	"os"
	"os/signal"
)

type options struct {
	addr string
}

type server struct {
	exit            chan struct{}
	opts            options
	connections     map[int64]net.Conn
	onServerStarted OnServerStartedFunc
}

func NewServer() Server {
	return &server{
		exit: make(chan struct{}, 1),
		opts: options{
			addr: ":9898",
		},
		connections: map[int64]net.Conn{},
	}
}

func (s *server) Listen() error {
	return s.listen()
}

func (s *server) ListenWithAddr(addr string) error {
	s.opts.addr = addr

	return s.listen()
}

func (s *server) Deaf() error {
	fmt.Println("Stopping server")
	s.exit <- struct{}{}
	return nil
}

func (s *server) OnServerStarted(handler OnServerStartedFunc) {
	s.onServerStarted = handler
}

func (s *server) listen() error {
	listener, err := net.Listen("tcp", s.opts.addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		fmt.Println("Waiting for connections...")
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			go s.handle(conn)
		}
	}()

	go func() {
		fmt.Println("Waiting for signal...")
		for sig := range c {
			fmt.Println("Received signal:", sig)
			s.Deaf()
			return
		}
	}()

	fmt.Println("Server listening on", s.opts.addr)

	// Calling hooks
	if s.onServerStarted != nil {
		s.onServerStarted()
	}

	<-s.exit
	fmt.Println("Server stopped")
	return nil
}

func (s *server) handle(conn net.Conn) error {
	fmt.Println("Handling connection: ", conn.LocalAddr())
	return nil
}
