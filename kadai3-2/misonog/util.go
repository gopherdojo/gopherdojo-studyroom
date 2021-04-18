package main

import (
	"fmt"
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
	// like setter
	SetFileName(string)
	SetFileSize(uint)
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

// SetFullFileName set to Data structs member
func (d *Data) SetFullFileName(dir, filename string) {
	if dir == "" {
		d.fullfilename = filename
	} else {
		d.fullfilename = fmt.Sprintf("%s/%s", dir, filename)
	}
}
