package eimg

// import (
//     "testing"
// )

// func TestRun(t *testing.T) {
//     cases := []struct {
//         name string
//         rootDir string
//         fromExt string
//         toExt string
//         expectedPkg string
//     } {
//         {name: "no arguments", rootDir: "", fromExt: "", toExt: "", expectedPkg: "test/noArguments.tar.gz"},
//             {name: "set RootDir only", rootDir: "test/documents", fromExt: "", toExt: "", expectedPkg: "test/setRootDirOnly.tar.gz"},
//             {name: "set RootDir and FromExt", rootDir: "test/img", fromExt: "gif", toExt: "", expectedPkg: "test/setRootDirAndFromExt.tar.gz"},
//             {name: "set RootDir and ToExt", rootDir: "test/img", fromExt: "", toExt: "gif", expectedPkg: "test/setRootDirAndToExt.tar.gz"},
//             {name: "set FromExt and ToExt", rootDir: "", fromExt: "png", toExt: "jpg", expectedPkg: "test/setRootDirAndFromExt.tar.gz"},
//             {name: "set all arguments", rootDir: "test/img", fromExt: "gif", toExt: "jpeg", expectedPkg: "test/setAllArgs.tar.gz"},
//             {name: "set other extensions", rootDir: "", fromExt: "txt", toExt: "rtf", expectedPkg: "test/setOtherExtensions.tar.gz"},
//         }

//     for _, c := range cases {
//         t.Run(c.name, func(t *testing.T) {
//             fmt.Printf("[TEST] %s begins\n", c.name)
//             err := exec.Command("")
//         })
//     }    
// }
