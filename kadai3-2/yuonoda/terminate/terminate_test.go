package terminate_test

import (
	"github.com/yuonoda/gopherdojo-studyroom/kadai3-2/yuonoda/terminate"
	"os"
	"syscall"
	"testing"
	"time"
)

func TestListen(t *testing.T) {
	ctx := terminate.Listen()

	// 中断シグナルを出す
	proc, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatal(err)
	}
	err = proc.Signal(syscall.SIGTERM)
	if err != nil {
		t.Fatal(err)
	}

	// 時間内にコンテクストがクローズするか
	select {
	case <-ctx.Done():
		return
	case <-time.After(100 * time.Millisecond):
		t.Fatal("termination timeout")

	}
}
