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

	// like setter
	SetFileName(string)
	SetFileSize(uint)
	SetDirName(string)
	SetFullFileName(string, string)

	// like getter
	FileName() string
	FileSize() uint
	DirName() string
	FullFileName() string
}

// FileName get from Data structs member
func (d Data) FileName() string {
	return d.filename
}

// FileSize get from Data structs member
func (d Data) FileSize() uint {
	return d.filesize
}

// DirName get from Data structs member
func (d Data) DirName() string {
	return d.dirname
}

// FullFileName get from Data structs member
func (d Data) FullFileName() string {
	return d.fullfilename
}

// SetFileName set to Data structs member
func (d *Data) SetFileName(filename string) {
	d.filename = filename
}

// SetFileSize set to Data structs member
func (d *Data) SetFileSize(size uint) {
	d.filesize = size
}

// SetDirName set to Data structs member
func (d *Data) SetDirName(dir string) {
	d.dirname = dir
}

// SetFullFileName set to Data structs member
func (d *Data) SetFullFileName(dir, filename string) {
	if dir == "" {
		d.fullfilename = filename
	} else {
		d.fullfilename = fmt.Sprintf("%s/%s", dir, filename)
	}
}

// MakeRange will return Range struct to download function
func (d *Data) MakeRange(i, split, procs uint) Range {
	low := split * i
	high := low + split - 1
	if i == procs-1 {
		high = d.FileSize()
	}

	return Range{
		low:   low,
		high:  high,
		woker: i,
	}
}

// MergeFiles function merege file after split download
func (d *Data) MergeFiles(procs int) error {
	filename := d.filename
	dirname := d.dirname

	mergefile, err := os.Create(d.fullfilename)
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

		io.Copy(mergefile, subfp)
		subfp.Close()

		if err := os.Remove(f); err != nil {
			return err
		}
	}

	return nil
}
