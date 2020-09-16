package imgconv

import (
	"errors"
	"testing"
)

func TestValidationSuccess(t *testing.T) {
	patterns := []struct {
		title string
		ie    string
		oe    string
	}{
		{"jpg->jpg", "jpg", "jpg"},
		{"jpg->jpeg", "jpg", "jpeg"},
		{"jpg->png", "jpg", "png"},
		{"jpg->gif", "jpg", "gif"},
		{"png->gif", "png", "gif"},
		{"gif->jpg", "gif", "jpg"},
	}

	id := "../../"
	od := "../images/output"
	for _, p := range patterns {
		t.Run(p.title, func(t *testing.T) {
			received := Do(&id, &od, &p.ie, &p.oe)
			if received != nil {
				t.Fatalf("[FAIL] received: %s, expected: nil", received.Error())
			}
		})
	}
}

func TestValidationFail(t *testing.T) {
	patterns := []struct {
		id       string
		od       string
		ie       string
		oe       string
		title    string
		expected error
	}{
		{"../../none", "../images/output", "png", "gif", "invalid target dir", errors.New("invalid input dir")},
		{"", "../images/output", "png", "gif", "invalid target dir", errors.New("invalid input dir")},
		{"../../", "../images/none", "png", "gif", "invalid output dir", errors.New("invalid output dir")},
		{"../../", "", "png", "gif", "invalid output dir", errors.New("invalid output dir")},
		{"../../", "../images/output", "none", "gif", "invalid extension of input file", errors.New("invalid extension of input file")},
		{"../../", "../images/output", "png", "none", "invalid extension of output file", errors.New("invalid extension of output file")},
	}

	for _, p := range patterns {
		t.Run(p.title, func(t *testing.T) {
			received := Do(&p.id, &p.od, &p.ie, &p.oe)
			if received != nil && received.Error() != p.expected.Error() {
				t.Fatalf("[FAIL] received: %s, expected: %s", received.Error(), p.expected.Error())
			}
		})
	}
}
