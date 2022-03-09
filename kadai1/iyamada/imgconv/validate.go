package imgconv

import (
	"errors"
	"flag"
)

var inputExtFlag string
var outputExtFlag string

func init() {
	flag.StringVar(&inputExtFlag, "i", "jpg", "input file extension")
	flag.StringVar(&outputExtFlag, "o", "png", "output file extension")
}

func validateFlag() (from, to string, err error) {
	switch inputExtFlag {
	case "jpg", "jpeg", "png", "gif":
		break
	default:
		return "", "", errors.New("error: invalid extension")
	}
	switch outputExtFlag {
	case "jpg", "jpeg", "png", "gif":
		break
	default:
		return "", "", errors.New("error: invalid extension")
	}
	return inputExtFlag, outputExtFlag, nil
}

// ValidateArgs is a function that retrieves the directory containing the
// image file and the extension of the image file from a command line argument.
// It returns an error if the arguments are not appropriate.
func ValidateArgs(args []string) (dirs []string, from, to string, err error) {
	if len(args) == 1 {
		return nil, "", "", errors.New("error: invalid argument\n")
	}
	flag.CommandLine.Parse(args[1:])
	from, to, err = validateFlag()
	if err != nil {
		return nil, "", "", err
	}
	return flag.Args(), from, to, nil
}
