package min

import (
	"log"
	"testing"
)

// TestStartStop tests the Start function.
func TestListen(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := Listen(); err != nil {
		t.Error(err)
	}
}

// TestListenWithAddr9999 tests the ListenWithAddr function.
// func TestListenWithAddr9999(t *testing.T) {
// 	s := NewServer()
//
// 	log.SetFlags(log.LstdFlags | log.Lshortfile)
//
// 	s.OnReady(func(s *server) {
// 		if s.Options().addr != ":9999" {
// 			t.Errorf("Expected :9999, got %s", s.Options().addr)
// 		}
// 		time.Sleep(1000 * time.Millisecond)
// 		if conn, err := net.Dial("tcp", ":9999"); err != nil {
// 			t.Error(err)
// 		} else {
// 			conn.Close()
// 		}
//
// 		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
// 	})
//
// 	if err := s.ListenWithAddr(":9999"); err != nil {
// 		t.Error(err)
// 	}
// }

// TestListenWithAddrFoo tests the ListenWithAddr function.
// func TestListenWithAddrFoo(t *testing.T) {
// 	s := NewServer()
//
// 	if err := s.ListenWithAddr("foo"); err != nil {
// 		t.Log(err)
// 	} else {
// 		t.Error("Expected error, got nil")
// 	}
// }

// TestListenWithAddrWithWebsocket tests the ListenWithAddr function.
// func TestListenWithAddrWithWebsocket(t *testing.T) {
// 	s := NewServer()
//
// 	s.OnReady(func(s *server) {
// 		if s.Options().addr != ":9999" {
// 			t.Errorf("Expected :9999, got %s", s.Options().addr)
// 		}
// 		if s.Options().isWebsocket != true {
// 			t.Errorf("Expected true, got %t", s.Options().isWebsocket)
// 		}
// 		time.Sleep(1000 * time.Millisecond)
// 		if conn, err := net.Dial("tcp", ":9999"); err != nil {
// 			t.Error(err)
// 		} else {
// 			conn.Close()
// 		}
//
// 		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
// 	})
//
// 	if err := s.ListenWithAddr(":9999", WithWebsocket()); err != nil {
// 		t.Error(err)
// 	}
// }
