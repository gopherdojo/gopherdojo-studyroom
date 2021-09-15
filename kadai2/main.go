package main

import (
    "convert/convert"
    "flag"
    "fmt"
)

// initialize flag
var (
    from = flag.String("from", "jpeg", "from extension")
    to = flag.String("to", "png", "after extension")
    directory = flag.String("directory", "images", "directory path")
)

func init() {
    flag.Usage = func() {
        fmt.Printf(`Usage: -from FROM_FORMAT -to TO_FORMAT -dir DIRECTORY
        Use: convert image files.
        Default: from jpeg to png.
        `)
        flag.PrintDefaults()
    }
}

func main() {
    flag.Parse()

    conv, err := convert.NewConv(*from, *to, *directory)
    if err != nil {
            fmt.Println(err)
    }

    paths, err := conv.FileSearch(*directory, *from)
    if err != nil {
            fmt.Println(err)
    }
    for _, path := range paths {
        err := conv.Convert(path, *to)
        if err != nil {
            fmt.Println(err)
        }
    }
}
