package imgconv

import (
	"flag"
)

var inputExtFlag string
var outputExtFlag string

func init() {
	flag.StringVar(&inputExtFlag, "i", "jpg", "input file extension")
	flag.StringVar(&outputExtFlag, "o", "png", "output file extension")
}

func parseFlag(in, out string) (from, to string, err error) {
	if !isValidFileExtent(in) || !isValidFileExtent(out) {
		return "", "", invalidExt
	}
	return in, out, nil
}

func isNoArg(args []string) bool {
	return len(args) == 1
}

// Parse is a function that retrieves the directory containing the
// image file and the extension of the image file from a command line argument.
// It returns an error if the arguments are not appropriate.
func Parse(args []string) (dirs []string, from, to string, err error) {
	if isNoArg(args) {
		return nil, "", "", invalidArg
	}
	flag.CommandLine.Parse(args[1:])
	from, to, err = parseFlag(inputExtFlag, outputExtFlag)
	if err != nil {
		return nil, "", "", err
	}
	return flag.Args(), from, to, nil
}
