package termination

import (
	"context"
	"io"
	"os"
	"testing"
	"time"
)

func TestListen(t *testing.T) {
	CleanFunc(func() {})

	doneCh := make(chan struct{})
	osExit = func(code int) { doneCh <- struct{}{} }

	_, clean := Listen(context.Background(), io.Discard)
	defer clean()

	proc, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatal(err)
	}

	err = proc.Signal(os.Interrupt)
	if err != nil {
		t.Fatal(err)
	}

	select {
	case <-doneCh:
		return
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timeout")
	}
}
