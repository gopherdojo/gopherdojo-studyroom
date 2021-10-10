package file_searcher_test

import (
	"imgconv/file_searcher"
	"reflect"
	"testing"
)

func TestDo(t *testing.T) {
	cases := []struct{
		name string
		dir string
		ext string
		output []string
	}{
		{name : "jpeg, current_dir", dir : "../testdata/dst", ext : "jpeg", output : []string{"../testdata/dst/jpeg/1.jpeg", "../testdata/dst/jpeg/sub/1.jpeg", "../testdata/dst/jpeg/sub/sub/1.jpeg"}},
		{name : "png, current_dir", dir : "../testdata/dst", ext : "png", output : []string{"../testdata/dst/png/1.png", "../testdata/dst/png/sub/1.png", "../testdata/dst/png/sub/sub/1.png"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			searcher := &file_searcher.FileSearcher{Dir: c.dir, Ext: c.ext}
			paths, err := searcher.Do()
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(paths, c.output) {
				t.Errorf("invalid result. testCase:%v, actual:%v", c.output, paths)
			}
		})
	}
}
