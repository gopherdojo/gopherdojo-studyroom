package starter

import (
	"testing"
	"time"

	"github.com/kotaaaa/gopherdojo-studyroom/kadai3-1/kotaaaa/questions"
)

func TestSolve(t *testing.T) {
	vc := questions.Vocab{
		Words:    []string{"apple", "red fruit"},
		Meanings: []string{"orange", "orange fruit"},
	}
	timeCh := time.NewTimer(time.Duration(0) * time.Second)
	ret := Solve(&vc, timeCh)
	if ret != 2 {
		t.Errorf("Return is illegal. Result: %v, Expected: 2(Fail) ", ret)
	}
}
