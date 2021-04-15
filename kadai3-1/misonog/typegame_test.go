package main

import (
	"bytes"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	t.Parallel()
	in := make(chan string)
	go func() {
		time.Sleep(100 * time.Millisecond)
		in <- "rust"
		time.Sleep(1 * time.Second)
	}()
	out := new(bytes.Buffer)
	run(in, out)

	expected := "rust"
	actual := strings.TrimRight(out.String(), "r\n")
	if actual != expected {
		t.Errorf(`expected="%s" actual="%s"`, expected, actual)
	}
}

func TestInput(t *testing.T) {
	t.Parallel()
	buf := bytes.NewBufferString("test")
	expected := "test"
	actual := <-input(buf)
	if actual != expected {
		t.Errorf(`expected="%s" actual="%s"`, expected, actual)
	}
}

func TestImportWords(t *testing.T) {
	t.Helper()

	// CreateTempを試しに利用してみる
	f, err := os.CreateTemp("", "word_list_*.txt")
	if err != nil {
		log.Fatalf("err %s", err)
	}
	defer os.Remove(f.Name())

	if _, err := f.Write([]byte("go")); err != nil {
		log.Fatal(err)
	}

	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{name: "import txt", input: "word_list.txt", expected: []string{"go", "rust", "r", "python", "ruby"}},
		{name: "temp txt", input: f.Name(), expected: []string{"go"}},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			actual, err := importWords(c.input)
			if err != nil {
				t.Fatalf("err %s", err)
			}

			if !reflect.DeepEqual(c.expected, actual) {
				t.Errorf("want importWords(%v) = %v, got %v",
					c.input, c.expected, actual)
			}
		})
	}
}
