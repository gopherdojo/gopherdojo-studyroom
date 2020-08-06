package eimg

import (
    "fmt"
    "os"
    "testing"
)

func TestEimg(t *testing.T) {
    TSetParameters(t)

    // TGetFilePathsRec(t, eimg)
    // TEncodeFile(t)
    // TConvertExtension(t)
}

func TSetParameters(t *testing.T) {
    eimg := New()
    
    cases := []struct {
        name string
        rootDir string
        fromExt string
        toExt string
        expected []string
    } {
            {name: "set RootDir only", rootDir: "test/documents", fromExt: "", toExt: "", expected: []string{"test/documents", "jpeg", "png"}},
            {name: "set RootDir and FromExt", rootDir: "test/img", fromExt: "gif", toExt: "", expected: []string{"test/img", "gif", "png"}},
            {name: "set RootDir and ToExt", rootDir: "test/img", fromExt: "", toExt: "gif" , expected: []string{"test/img", "jpeg", "gif"}},
            {name: "set all arguments", rootDir: "test/img", fromExt: "gif", toExt: "jpeg", expected: []string{"test/img", "gif", "jpeg"}},
        }
    
    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            fmt.Printf("[TEST] %s begins\n", c.name)
            
            os.Args = []string{
                map[bool]string{true: "", false: "-f="+c.fromExt}[c.fromExt == ""],
                map[bool]string{true: "", false: "-t="+c.toExt}[c.toExt == ""],
                map[bool]string{true: "", false: c.rootDir}[c.rootDir == ""],
            }

	   fmt.Println(os.Args)
            
            if err := eimg.SetParameters(); err != nil {
                t.Errorf("failed to set parameter: %#v", err)
            }
            
            if eimg.RootDir != c.expected[0] {
                t.Errorf("RootDir=> Actual: %s, Expected: %s", eimg.RootDir, c.expected[0])

            }
            if eimg.FromExt != c.expected[1] {
                t.Errorf("FromExt=> Actual: %s, Expected: %s", eimg.FromExt, c.expected[1])

            }

            if eimg.ToExt != c.expected[2] {
                t.Errorf("ToExt=> Actual: %s, Expected: %s", eimg.ToExt, c.expected[2])

            }

            os.Args = []string{}
        })
        
    }
}
