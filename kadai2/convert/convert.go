package convert

import (
    "io/ioutil"
    "path/filepath"
    "os"
    "image"
    "image/jpeg"
    "image/png"
    "errors"
)

// user defined type
type Conv struct {
    From string
    To   string
    Dir  string
}

// create Conv
func NewConv(from string, to string, dir string)(*Conv, error){
    return &Conv{from, to, dir}, nil
}

// search directory
func (conv *Conv)FileSearch(dir string, from string)([]string, error){
        var paths []string
        files, err := ioutil.ReadDir(dir)
        if err != nil {
            return nil, err
        }

        for _, file := range files {
            if file.IsDir() {
                subpaths, err := conv.FileSearch(filepath.Join(dir, file.Name()), from)
                paths = append(paths, subpaths...)
                if err != nil {
                    return nil, err
                }
                continue
            }
            ext := filepath.Ext(file.Name())

            if ext !="" && ext[1:] == from {
                fullpath := filepath.Join(dir, file.Name())
                paths = append(paths, fullpath)
            }
        }
        return paths, err
}

// replace filepath
func (conv *Conv) Convert(path string, to string) (error) {
    inputFile, err := os.Open(path)
    if err != nil {
        return err
    }
    defer inputFile.Close()

    img, _, err := image.Decode(inputFile)
    if err != nil {
        return err
    }

    out_file, err := os.Create(path[:len(path)-len(filepath.Ext(path))+1] + to)
    if err != nil {
        return err
    }
    defer out_file.Close()

    switch to {
    case "jpg", "jpeg" :
            err := jpeg.Encode(out_file, img, &jpeg.Options{})
            if err != nil {
                return err
            }
    case "png" :
            err := png.Encode(out_file, img)
            if err != nil {
                return err
            }
    default:
            errors.New("wrong after extension")
    }

    err = os.Remove(path)
    if err != nil {
        return err
    }

    return err
}
