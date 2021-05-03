package main

import (
	"fmt"
	"io"
	"os"
)

// Data struct has file of relational data
type Data struct {
	filename     string
	filesize     uint
	dirname      string
	fullfilename string
}

// Utils interface indicate function
type Utils interface {
	MakeRange(uint, uint, uint) Range
	MergeFiles(int) error
}

// mergeFiles function merege file after split download
func mergeFiles(procs int, filename, dirname, fullfilename string) error {
	mergefile, err := os.Create(fullfilename)
	if err != nil {
		return err
	}
	defer mergefile.Close()

	var f string
	for i := 0; i < procs; i++ {
		f = fmt.Sprintf("%s/%s.%d.%d", dirname, filename, procs, i)
		subfp, err := os.Open(f)
		if err != nil {
			return err
		}

		if _, err := io.Copy(mergefile, subfp); err != nil {
			return err
		}
		subfp.Close()

		if err := os.Remove(f); err != nil {
			return err
		}
	}

	return nil
}
