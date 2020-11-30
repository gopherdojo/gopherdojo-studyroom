package imgconv_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/takkyuuplayer/gopherdojo-studyroom/kadai1/imgconv"
)

func TestNew(t *testing.T) {
	t.Parallel()

	file, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	type args struct {
		directory string
		from      string
		to        string
	}
	tests := []struct {
		name    string
		args    args
		want    *imgconv.Converter
		wantErr bool
	}{
		{
			name:    "Supports .jpg => .png",
			args:    args{"./", "jpg", "png"},
			want:    &imgconv.Converter{"./", "jpg", "png"},
			wantErr: false,
		},
		{
			name:    "Supports .jpg => .gif",
			args:    args{"./", "jpg", "gif"},
			want:    &imgconv.Converter{"./", "jpg", "gif"},
			wantErr: false,
		},
		{
			name:    "Supports .png => .jpg",
			args:    args{"./", "png", "jpg"},
			want:    &imgconv.Converter{"./", "png", "jpg"},
			wantErr: false,
		},
		{
			name:    "Supports .png => .jpeg",
			args:    args{"./", "png", "jpeg"},
			want:    &imgconv.Converter{"./", "png", "jpeg"},
			wantErr: false,
		},
		{
			name:    "Invalid: .foo => .png",
			args:    args{"./", "foo", "png"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid: .jpg => .foo",
			args:    args{"./", "jpg", "foo"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid: .jpg => .jpg",
			args:    args{"./", "jpg", "jpg"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid: directory does not exists",
			args:    args{"/dir/does/not/exist", "jpg", "jpg"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid: directory is a file",
			args:    args{file.Name(), "jpg", "jpg"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := imgconv.New(tt.args.directory, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConverter_Walk(t *testing.T) {
	t.Parallel()
	type fields struct {
		FromExt string
		ToExt   string
	}
	tests := []struct {
		name              string
		fields            fields
		convertedFilePath string
		wantErr           bool
	}{
		{
			".jpg => .png",
			fields{"jpg", "png"},
			"/moon.png",
			false,
		},
		{
			".png => .jpg",
			fields{"png", "jpg"},
			"/sub/sun.jpg",
			false,
		},
		{
			".jpg => .gif",
			fields{"jpg", "gif"},
			"/moon.gif",
			false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			dir := createTestTempDir(t)
			//defer os.RemoveAll(dir)

			c := &imgconv.Converter{
				Directory: dir,
				FromExt:   tt.fields.FromExt,
				ToExt:     tt.fields.ToExt,
			}
			err := c.Walk()

			if (err != nil) != tt.wantErr {
				t.Errorf("Walk() error = %v, wantErr %v", err, tt.wantErr)
			}
			path := filepath.Join(dir, tt.convertedFilePath)
			if err == nil && !testExist(t, path) {
				t.Errorf("%v has not been created", path)
			}
		})
	}
}

func createTestTempDir(t *testing.T) string {
	t.Helper()

	dir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("cp", "-r", "testdata/", dir)
	err = cmd.Run()
	if err != nil {
		t.Fatal(fmt.Errorf("%v failed: %w", cmd.String(), err))
	}

	return dir
}

func testExist(t *testing.T, path string) bool {
	t.Helper()

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
