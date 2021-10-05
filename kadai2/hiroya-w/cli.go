package imgconv

import (
	"flag"
	"fmt"
	"io"
	"os"
)

// CLI is the command line interface
type CLI struct {
	OutStream, ErrStream io.Writer
}

// validateType validates the type of the image
func validateType(t string) error {
	switch t {
	case "jpg", "jpeg", "png", "gif":
		return nil
	default:
		return fmt.Errorf("invalid type: %s", t)
	}
}

// Run parses the command line arguments and runs the imgConv
func (cli *CLI) Run() int {
	config := &Config{}
	fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	fs.StringVar(&config.InputType, "input-type", "jpg", "input type[jpg|jpeg|png|gif]")
	fs.StringVar(&config.OutputType, "output-type", "png", "output type[jpg|jpeg|png|gif]")
	fs.SetOutput(cli.ErrStream)
	fs.Usage = func() {
		fmt.Fprintf(cli.ErrStream, "Usage: %s [options] DIRECTORY\n", "imgconv")
		fs.PrintDefaults()
	}

	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Fprintf(cli.ErrStream, "Error parsing arguments: %s\n", err)
		return 1
	}

	if err := validateType(config.InputType); err != nil {
		fmt.Fprintf(cli.ErrStream, "invalid input type: %s\n", err)
		return 1
	}

	if err := validateType(config.OutputType); err != nil {
		fmt.Fprintf(cli.ErrStream, "invalid output type: %s\n", err)
		return 1
	}

	if config.InputType == config.OutputType {
		fmt.Fprintf(cli.ErrStream, "input type and output type must be different\n")
		return 1
	}

	if fs.Arg(0) == "" {
		fmt.Fprintf(cli.ErrStream, "directory is required\n")
		return 1
	}

	config.Directory = fs.Arg(0)

	dec, err := NewDecoder(config.InputType)
	if err != nil {
		fmt.Fprintf(cli.ErrStream, "failed to create decoder: %s\n", err)
		return 1
	}
	enc, err := NewEncoder(config.OutputType)
	if err != nil {
		fmt.Fprintf(cli.ErrStream, "failed to create encoder: %s\n", err)
		return 1
	}
	imgConv := &ImgConv{
		OutStream: cli.OutStream,
		Decoder:   dec,
		Encoder:   enc,
		TargetDir: config.Directory,
	}
	convertedFiles, err := imgConv.Run()
	if err != nil {
		fmt.Fprintf(cli.ErrStream, "failed to convert images: %s\n", err)
		return 1
	}

	fmt.Fprintf(cli.OutStream, "converted %d files\n", len(convertedFiles))
	for _, f := range convertedFiles {
		fmt.Fprintf(cli.OutStream, "%s\n", f)
	}

	return 0
}
