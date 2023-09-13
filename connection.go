package min

import "net"

type connection struct {
	id   int64
	conn net.Conn
}

func newConnection(id int64, conn net.Conn) *connection {
	return &connection{
		id:   id,
		conn: conn,
	}
}

func (c *connection) Close() error {
	return c.conn.Close()
}
