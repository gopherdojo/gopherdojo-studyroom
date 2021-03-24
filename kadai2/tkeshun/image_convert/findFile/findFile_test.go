package findFile_test
import (
	"testing"
    "image_convert/findFile"
	"reflect"
)
func TestSearch(t *testing.T){
  cases := []struct{name string;path string; input_fmt string;expected []string}{
	  {name : "maindir, jpeg", path: "../testdir",input_fmt:"jpeg",expected :[]string{"../testdir/subdir/test_image1.jpeg", "../testdir/subdir/test_image2.jpeg", "../testdir/test_image1.jpeg", "../testdir/test_image2.jpeg"}},
	  {name : "maindir, png", path: "../testdir",input_fmt:"png",expected: []string{"../testdir/subdir/test_image1.png", "../testdir/subdir/test_image2.png", "../testdir/test_image1.png", "../testdir/test_image2.png"}},
	  {name : "subdir, jpeg", path: "../testdir/subdir",input_fmt:"jpeg", expected: []string{"../testdir/subdir/test_image1.jpeg", "../testdir/subdir/test_image2.jpeg"}},
	  {name : "subdir, png", path: "../testdir/subdir",input_fmt:"png", expected:[]string{"../testdir/subdir/test_image1.png", "../testdir/subdir/test_image2.png"}},
  }
  for _ , c := range cases {
    t.Run(c.name, func(t *testing.T) {
		if actual := findFile.Search(c.path,c.input_fmt); reflect.DeepEqual(actual,c.expected) != true{
           t.Errorf("want findFile(%s,%s) = %s, got %s",c.path,c.input_fmt,c.expected,actual)
		}
	})
  }
}

