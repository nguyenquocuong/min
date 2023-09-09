package min

import "net"

type OnServerStartedFunc func()

type Server interface {
	Listen() error
	ListenWithAddr(addr string) error
	Deaf() error

	// hooks
	OnServerStarted(OnServerStartedFunc)

	// private
	listen() error
	handle(conn net.Conn) error
}
