package interrupt_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/kimuson13/gopherdojo-studyroom/kimuson13/interrupt"
)

func TestInterrupt(t *testing.T) {
	ctx, cancel := interrupt.Listen(context.Background())
	defer cancel()

	proc, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatalf("err: %d", err)
	}

	err = proc.Signal(os.Interrupt)
	if err != nil {
		t.Fatalf("err: %d", err)
	}

	select {
	case <-ctx.Done():
		return
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timeout")
	}
}
