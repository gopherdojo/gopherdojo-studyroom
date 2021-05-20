package picconvert

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

// PicConverter is User-defined type for converting
// a picture file format has root path, pre-conversion format,
// and after-conversion format.
type PicConverter struct {
  Path        string
  PreFormat   []string
  AfterFormat string
}

// Conv converts picture file format.
func (p *PicConverter) Conv() {
  files, err := glob(p.Path, p.PreFormat)

  if err != nil {
    log.Fatal("could not find the file path.");
  }

  // prosessing for each files
  for _, file := range files {
    // fmt.Println("from:", file)
    f, err := os.Open(file)
    if err != nil {
      fmt.Fprintln(os.Stderr, " the file could not open.");
      os.Exit(1);
    }
    defer f.Close()

    // reading the image file.
    img, _, err := image.Decode(f)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Failed to convert image files.")
      os.Exit(1)
    }

    // creating filepath for output.
    output, err := os.Create(baseName(file) + 
                            "_converted" +
                            "." + 
                            p.AfterFormat)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Failed to create the file.")
      os.Exit(1)
    }

    // converting the file.
    if p.AfterFormat == "png" {
      err = png.Encode(output, img)
    } else if p.AfterFormat == "jpg" || p.AfterFormat == "jpeg" {
      err = jpeg.Encode(output, img, nil)
    } else if p.AfterFormat == "gif" {
      err = gif.Encode(output, img, nil)
    }

    if err != nil {
      fmt.Fprintln(os.Stderr, err)
      continue
    }
    
    // fmt.Println("to:", output.Name())
  }
}

// NewPicConverter is the constructor for a PicConverter.
func NewPicConverter(Path string, PreFormat string, AfterFormat string) *PicConverter {
  res := new(PicConverter)
  res.Path = Path

  if PreFormat == "jpeg" || PreFormat == "jpg" {
    res.PreFormat = []string{"jpeg", "jpg"}
  } else {
    res.PreFormat = []string{PreFormat}
  }

  res.AfterFormat = AfterFormat
  return res
}

// glob returns a slice of file path where format.
func glob(path string, format []string) ([]string, error) {
  var res []string

  var err error 
  for _, f := range format {
    err = filepath.Walk(path,
      func(path string, info os.FileInfo, err error) error {
        if filepath.Ext(path) == "." + f {
          res = append(res, path)
        }
        return nil
      })
  }

  return res, err
}

// baseName returns the filepath without a extension.
func baseName(filePath string) string {
  ext := filepath.Ext(filePath)
  return filePath[:len(filePath)-len(ext)]
}
