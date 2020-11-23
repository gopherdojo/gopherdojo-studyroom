package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func Test_validate_ng(t *testing.T) {
	c := Convertor{}
	msg, ok := c.validate()
	if msg != "-inフラグの入力は必須です" {
		t.Errorf("msg=%s, expected=[-inフラグの入力は必須です]\n", msg)
	}
	if ok {
		t.Errorf("validate() %T, expected false.\n", ok)
	}
}
