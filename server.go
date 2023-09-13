package min

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/gorilla/websocket"
)

type server struct {
	listener net.Listener
	srv      *http.Server
	err      chan error
	opts     options
	handler  *handler

	mu          sync.RWMutex
	connections map[int64]*connection
}

func newServer() *server {
	return &server{
		err: make(chan error),
		opts: options{
			addr:        ":9898",
			isWebsocket: false,
		},
		handler:     newHandler(),
		connections: map[int64]*connection{},
	}
}

func (s *server) listen() error {

	go func() {
		if !s.opts.isWebsocket {
			s.quit(s.serveTCP())
		} else {
			s.quit(s.serveWebsocket())
		}
	}()

	go s.listenSignals()

	<-s.err
	log.Println("Server stopped")
	return nil
}

func (s *server) serveTCP() error {
	listener, err := net.Listen("tcp", s.opts.addr)
	if err != nil {
		return err
	}

	s.listener = listener

	log.Println("Waiting for TCP connections on", s.opts.addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				log.Println("Connection closed by server")
				return nil
			}
			log.Println("Error accepting connection:", err)
			continue
		}

		c := s.addNewConnection(conn)
		go s.handler.handle(c)
	}
}

func (s *server) serveWebsocket() error {
	log.Println("Waiting for websocket connections...")

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	s.srv = &http.Server{
		Addr: s.opts.addr,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		go s.handler.handleWebsocket(conn)
	})

	if err := s.srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (s *server) listenSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	log.Println("Waiting for signal...")
	for sig := range c {
		log.Println("Received signal:", sig)
		s.close()
		return
	}
}

func (s *server) quit(err error) {
	s.err <- err
}

func (s *server) close() error {
	log.Println("Stopping server")
	if s.listener != nil {
		log.Println("Closing listener")
		s.listener.Close()
	}
	if s.srv != nil {
		log.Println("Shutting down server")
		s.srv.Shutdown(context.TODO())
	}
	s.err <- nil
	return nil
}

func (s *server) uniqueID() int64 {
	for {
		id := rand.Int63()
		if _, exists := s.connections[id]; !exists {
			return id
		}
	}
}

func (s *server) addNewConnection(conn net.Conn) *connection {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.uniqueID()
	c := newConnection(id, conn)
	s.connections[id] = c

	return c
}
