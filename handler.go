package min

import (
	"errors"
	"io"
	"log"

	"github.com/gorilla/websocket"
)

type handler struct{}

func newHandler() *handler {
	return &handler{}
}

func (h *handler) handle(c *connection) {
	log.Printf("new connection: id=%d\n", c.id)

	go func() {
		defer c.Close()

		buf := make([]byte, 1024)

		for {
			buf := buf[:cap(buf)]
			n, err := c.conn.Read(buf)
			if err != nil {
				if errors.Is(err, io.EOF) {
					log.Println("Connection closed")
					return
				}
				log.Println("Error reading:", err.Error())
				return
			}
			log.Println("Received:", string(buf[:n-1]))
		}
	}()
}

func (h *handler) handleWebsocket(conn *websocket.Conn) {
	log.Println("new websocket connection:", conn.LocalAddr())
}
