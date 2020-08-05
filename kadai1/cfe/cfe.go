// cfe package converts file extention
// - mandatory
//   - set root directory
//     - default setting is directory executed this command
//   - execute recursively
// - optional
//   - arguments
//     - `-f`
//       - file extention before executing
//       - default setting is jpg
//     - `-t`
//       - file extention after executing
//       - default setting is png

package cfe

import (
    "os"
    "flag"
)

const (
    version = "0.0.1"
    msg = "cfe v" + version + ", converts file extention\n"
)

// Cfe structs
type Cfe struct {
    RootDir string
    From string
    To   string
}

// New for cfe package
func New() *Cfe {
    return &Cfe{
        rootDir: ".",
        from: "jpg",
        to: "png",
    }
}

// Run converts file extension(from -> to).
func (cfe *Cfe) Run() error {
    if err := cfe.SetParameters(); err != nil {
        return err
    }
}

// SetParameters sets parameters for execution.
func (cfe *Cfe) SetParameters() error {
    // parse information
    fr := flag.String("f", "jpg", "file extention before executing")
    to := flag.String("t", "png", "file extention after executing")
    
    flag.Parse()
    args := flag.Args()

    // set information
    if *fr != "jpg" {
        cfe.From = *fr
    }
    if *to != "png" {
        cfe.To = *to
    }

    // use default setting
    if len(args) == 0 {
        return nil
    }

    if args[0] != "." {
        if _, err := os.Stat(args[0]); err != nil {
            return ErrInvalidPath.WithDebug(err.Error())
        }
        cfe.RootDir = args[0]
    }

    return nil
}
