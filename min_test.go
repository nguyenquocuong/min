package min

import (
	"syscall"
	"testing"
)

// TestNewServer tests the NewServer function.
func TestNewServer(t *testing.T) {
	s := NewServer()
	if s == nil {
		t.Error("NewServer returned nil")
	}
}

// TestStartStop tests the Start function.
func TestListen(t *testing.T) {
	s := NewServer()

	s.OnServerStarted(func() {
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	})

	if err := s.Listen(); err != nil {
		t.Error(err)
	}
}

// TestListenWithAddr9999 tests the ListenWithAddr function.
func TestListenWithAddr9999(t *testing.T) {
	s := NewServer()

	s.OnServerStarted(func() {
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	})

	if err := s.ListenWithAddr(":9999"); err != nil {
		t.Error(err)
	}
}

// TestListenWithAddrFoo tests the ListenWithAddr function.
func TestListenWithAddrFoo(t *testing.T) {
	s := NewServer()

	if err := s.ListenWithAddr("foo"); err != nil {
		t.Log(err)
	} else {
		t.Error("Expected error, got nil")
	}
}
