package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Converter is a image converter
type Converter struct {
	Directory string
	// The extension to convert images from
	FromExt string
	// The extension to convert images to
	ToExt string
}

var supported = map[string]struct{}{
	"jpg":  true,
	"jpeg": true,
	"png":  true,
	"gif":  true,
}

// New allocates a new image converter
func New(directory, from, to string) (*Converter, error) {
	if err := validate(directory, strings.ToLower(from), strings.ToLower(to)); err != nil {
		return nil, err
	}
	return &Converter{directory, from, to}, nil
}

func validate(directory, from, to string) error {
	src := filepath.Clean(directory)
	info, err := os.Stat(src)

	if os.IsNotExist(err) {
		err = fmt.Errorf("%s is not exists: %w", src, err)
		return err
	}

	if !info.IsDir() {
		return fmt.Errorf("%s must be directory:", src)
	}

	if _, ok := supported[from]; !ok {
		return fmt.Errorf("unsupported format %s", from)
	}

	if _, ok := supported[to]; !ok {
		return fmt.Errorf("unsupported format %s", from)
	}

	if from == to {
		return fmt.Errorf("%s should be different from %s", from, to)
	}

	return nil
}

// Walk converts all images in the directory recursively.
func (c *Converter) Walk() error {
	return filepath.Walk(c.Directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		ext := normalizedExt(path)
		if ext != c.FromExt {
			return nil
		}
		reader, rErr := os.Open(path)
		if rErr != nil {
			return fmt.Errorf("failed to open %s: %v", path, rErr)
		}
		defer reader.Close()

		dest := strings.TrimSuffix(path, filepath.Ext(path)) + "." + c.ToExt
		writer, wErr := os.Create(dest)
		if wErr != nil {
			return fmt.Errorf("failed to write to %s: %v", dest, wErr)
		}
		defer writer.Close()

		return c.convert(writer, reader)
	})
}
func (c *Converter) convert(w io.Writer, r io.Reader) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}

	switch c.ToExt {
	case "jpg", "jpeg":
		fallthrough
	case "jpeg":
		return jpeg.Encode(w, img, nil)
	case "png":
		return png.Encode(w, img)
	case "gif":
		return gif.Encode(w, img, nil)
	default:
		return fmt.Errorf("unknown format %s", c.ToExt)
	}
}
func normalizedExt(path string) string {
	ext := strings.TrimLeft(filepath.Ext(path), ".")
	ext = strings.ToLower(ext)

	switch ext {
	case "jpeg":
		fallthrough
	case "jpg":
		return "jpg"
	default:
		return ext
	}
}
