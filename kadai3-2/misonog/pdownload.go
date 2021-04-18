package main

// Pdownload structs
type Pdownload struct {
	Utils
	URL       string
	TargetDir string
}

func New() *Pdownload {
	return &Pdownload{
		Utils: &Data{},
	}
}
