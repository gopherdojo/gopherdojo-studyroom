package convert_test

import (
    "testing"
	"reflect"
	"convert/convert"
)

func TestFileSearch(t *testing.T) {
        cases := []struct{
            name string
            from string
            dir string
            output []string
            }{
            {name : "jpeg, current_dir" , from : "jpeg", dir : "../test_images", output : []string{"../test_images/test_image1.jpeg", "../test_images/test_image2.jpeg","../test_images/test_subimages/test_image1.jpeg","../test_images/test_subimages/test_image2.jpeg"}},
            {name : "png, current_dir" , from : "png", dir : "../test_images", output : []string{"../test_images/test_image1.png", "../test_images/test_image2.png","../test_images/test_subimages/test_image1.png","../test_images/test_subimages/test_image2.png"}},
            {name : "jpeg, sub_dir" , from : "jpeg", dir : "../test_images/test_subimages", output : []string{"../test_images/test_subimages/test_image1.jpeg","../test_images/test_subimages/test_image2.jpeg"}},
            {name : "png, sub_dir" , from : "png", dir : "../test_images/test_subimages", output : []string{"../test_images/test_subimages/test_image1.png","../test_images/test_subimages/test_image2.png"}},
            }

        for _, c := range cases {
            t.Run(c.name, func(t *testing.T) {
                conv := &convert.Conv{c.from, "", c.dir}
                paths, err := conv.FileSearch(c.dir, c.from);
                if err != nil {
                    t.Error(err)
                }
                if !reflect.DeepEqual(paths, c.output) {
                    t.Errorf("invalid result. testCase:%#v, actual:%v", c.output, paths)
                }
            })
        }
}

func CheckErr(t *testing.T, err error) bool {
  t.Helper()
  if err !=nil {
	t.Errorf("%s",err)
	return true
  }
  return false
}

func TestConvert (t *testing.T) {
        cases := []struct{
            name string
            path string
            to string
            output error
            }{
            {name : "jpeg, png" , path : "../test_images/test_image1.jpeg", to : "png",  output : nil},
            {name : "jpeg, png" , path : "../test_images/test_image2.jpeg", to : "png",  output : nil},
            {name : "png, jpeg" , path : "../test_images/test_image1.png", to : "jpeg",  output : nil},
            {name : "png, jpeg" , path : "../test_images/test_image2.png", to : "jpeg",  output : nil},
            }

        conv := &convert.Conv{"", "", ""}
        for _, c := range cases {
            err := conv.Convert(c.path, c.to)
            ch := CheckErr(t, err)
            if !ch && err != c.output {
                t.Errorf("invalid result. testCase:%#v, actual:%v", c.output, err)
            }
        }
}
