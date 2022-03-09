package imgconv_test

import (
	"reflect"
	"sync"
	"testing"

	"example.com/ex01/imgconv"
)

var mtx sync.Mutex

func TestValidateArgs(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantDirs []string
		wantFrom string
		wantTo   string
		wantErr  bool
	}{
		{"default", []string{"cmd", "/some/directory"}, []string{"/some/directory"}, "jpg", "png", false},
		{"assign format", []string{"cmd", "-i=png", "-o=jpg", "/some/directory"}, []string{"/some/directory"}, "png", "jpg", false},
		{"assign format", []string{"cmd", "-i=png", "-o=jpeg", "/some/directory"}, []string{"/some/directory"}, "png", "jpeg", false},
		{"assign format", []string{"cmd", "-i=png", "-o=gif", "/some/directory"}, []string{"/some/directory"}, "png", "gif", false},
		{"assign format", []string{"cmd", "-i=png", "-o=png", "/some/directory"}, []string{"/some/directory"}, "png", "png", false},
		{"assign format", []string{"cmd", "-i=jpg", "-o=png", "/some/directory"}, []string{"/some/directory"}, "jpg", "png", false},
		{"assign format", []string{"cmd", "-i=jpg", "-o=jpeg", "/some/directory"}, []string{"/some/directory"}, "jpg", "jpeg", false},
		{"assign format", []string{"cmd", "-i=jpg", "-o=gif", "/some/directory"}, []string{"/some/directory"}, "jpg", "gif", false},
		{"assign format", []string{"cmd", "-i=jpg", "-o=jpg", "/some/directory"}, []string{"/some/directory"}, "jpg", "jpg", false},
		{"assign format", []string{"cmd", "-i=jpeg", "-o=gif", "/some/directory"}, []string{"/some/directory"}, "jpeg", "gif", false},
		{"assign format", []string{"cmd", "-i=jpeg", "-o=png", "/some/directory"}, []string{"/some/directory"}, "jpeg", "png", false},
		{"assign format", []string{"cmd", "-i=jpeg", "-o=jpg", "/some/directory"}, []string{"/some/directory"}, "jpeg", "jpg", false},
		{"assign format", []string{"cmd", "-i=jpeg", "-o=jpeg", "/some/directory"}, []string{"/some/directory"}, "jpeg", "jpeg", false},
		{"assign format", []string{"cmd", "-i=gif", "-o=jpg", "/some/directory"}, []string{"/some/directory"}, "gif", "jpg", false},
		{"assign format", []string{"cmd", "-i=gif", "-o=png", "/some/directory"}, []string{"/some/directory"}, "gif", "png", false},
		{"assign format", []string{"cmd", "-i=gif", "-o=jpeg", "/some/directory"}, []string{"/some/directory"}, "gif", "jpeg", false},
		{"assign format", []string{"cmd", "-i=gif", "-o=gif", "/some/directory"}, []string{"/some/directory"}, "gif", "gif", false},
		{"no args", []string{"cmd"}, nil, "", "", true},
		{"invalid format", []string{"cmd", "-i=txt", "-o=bmp", "/some/directory"}, nil, "", "", true},
		{"invalid format", []string{"cmd", "-i=jpg", "-o=txt", "/some/directory"}, nil, "", "", true},
		{"invalid format", []string{"cmd", "-i=mp3", "-o=mp4", "/some/directory"}, nil, "", "", true},
		{"multi flag", []string{"cmd", "-i=jpg", "-i=jpeg", "-i=png", "-o=gif", "/some/directory"}, []string{"/some/directory"}, "png", "gif", false},
		{"multi flag", []string{"cmd", "-i=jpg", "-o=jpeg", "-o=png", "-o=gif", "/some/directory"}, []string{"/some/directory"}, "jpg", "gif", false},
		// {"assign error", []string{"cmd", "/some/directory", "-i=gif", "-o=jpg"}, []string{"/some/directory"}, "gif", "jpg", false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mtx.Lock()
			gotDirs, gotFrom, gotTo, err := imgconv.ValidateArgs(tt.args)
			mtx.Unlock()
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDirs, tt.wantDirs) {
				t.Errorf("ValidateArgs() gotDirs = %v, want %v", gotDirs, tt.wantDirs)
			}
			if gotFrom != tt.wantFrom {
				t.Errorf("ValidateArgs() gotFrom = %v, want %v", gotFrom, tt.wantFrom)
			}
			if gotTo != tt.wantTo {
				t.Errorf("ValidateArgs() gotTo = %v, want %v", gotTo, tt.wantTo)
			}
		})
	}
}
