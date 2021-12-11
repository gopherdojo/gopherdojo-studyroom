package termination

import (
	"context"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestTermination_Listen(t *testing.T) {
	CleanFunc(func() {})

	doneCh := make(chan struct{})
	osExit = func(code int) { doneCh <- struct{}{} }

	_, clean := Listen(context.Background(), ioutil.Discard)
	defer clean()

	proc, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatalf("err %s", err)
	}

	err = proc.Signal(os.Interrupt)
	if err != nil {
		t.Fatalf("err %s", err)
	}

	select {
	case <-doneCh:
		return
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timeout")
	}
}
