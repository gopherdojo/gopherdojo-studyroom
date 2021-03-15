package lib

import "testing"

func TestExistDir(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected bool
	}{
		{name: "Dir exsit", input: "../testdata/png", expected: true},
		{name: "Dir not exist", input: "../testdata/notexist", expected: false},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if actual := ExistDir(c.input); c.expected != actual {
				t.Errorf("want ExitDir(%v) = %v, got %v",
					c.input, c.expected, actual)
			}
		})
	}
}
