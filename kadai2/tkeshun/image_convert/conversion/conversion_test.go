package conversion_test

import (
	"testing"
    "image_convert/conversion"    
)

func CheckErr(t *testing.T, err error) bool {
  t.Helper()
  if err !=nil {
	t.Errorf("%s",err)
	return true
  }
  return false
}


func TestConvert(t *testing.T){
	//errMsg := errors.New("指定した拡張子がまちがっている可能性があります" )
	cases := []struct {name string; srcPath string ;output_fmt string;expected error}{
		{name:"png to jpeg", srcPath:"../testdir/test_image1.png", output_fmt:"jpeg", expected:nil},
		{name:"png to jpeg", srcPath:"../testdir/test_image2.png", output_fmt:"jpeg", expected:nil},
		{name:"jpeg to png", srcPath:"../testdir/test_image1.jpeg", output_fmt:"png", expected:nil },
		{name:"jpeg to png", srcPath:"../testdir/test_image2.jpeg", output_fmt:"png", expected:nil },
		
	    
	}
	for _,c := range cases {
	    actual := conversion.Convert(c.srcPath,c.output_fmt)
		ch := CheckErr(t,actual)
		if  !ch && actual != c.expected {
			t.Errorf("want Convert(%s,%s) = %s, got %s",c.srcPath, c.output_fmt, c.expected, actual)
		}
		
	}
}