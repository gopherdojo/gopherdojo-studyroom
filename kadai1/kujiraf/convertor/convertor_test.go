package imgconv

import (
	"testing"
)

func Test_validate_ng(t *testing.T) {
	c := Convertor{}
	err := c.Validate()
	if err == nil {
		t.Errorf("validation error is expected. src=%s", c.Src)
	}
	if err.Error() != "-in is required" {
		t.Errorf("msg=%s, expected=[-in is required]\n", err.Error())
	}
}
