package main

import "runtime"

// Pdownload structs
type Pdownload struct {
	Utils
	URL       string
	TargetDir string
	Procs     int
	useragent string
	referer   string
}

func New() *Pdownload {
	return &Pdownload{
		Utils: &Data{},
		Procs: runtime.NumCPU(), // default
	}
}
