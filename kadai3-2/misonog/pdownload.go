package main

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"runtime"
	"time"

	"github.com/misonog/gopherdojo-studyroom/kadai3-2/misonog/termination"
)

// Pdownload structs
type Pdownload struct {
	Utils
	URL       string
	TargetDir string
	Procs     int
	timeout   time.Duration
	useragent string
	referer   string
}

func New() *Pdownload {
	return &Pdownload{
		Utils:   &Data{},
		Procs:   runtime.NumCPU(), // default
		timeout: timeout,
	}
}

func (pdownload *Pdownload) Run(ctx context.Context, args []string, targetDir string, timeout time.Duration) error {
	var cancel context.CancelFunc

	ctx, clean := termination.Listen(ctx, os.Stdout)
	defer clean()

	if err := pdownload.Ready(args, targetDir, timeout); err != nil {
		return err
	}

	dir, err := os.MkdirTemp(pdownload.TargetDir, "")
	if err != nil {
		return err
	}
	clean = func() { os.RemoveAll(dir) }
	defer clean()
	termination.CleanFunc(clean)

	ctx, cancel = context.WithTimeout(ctx, pdownload.timeout)
	defer cancel()

	err = pdownload.Check(ctx, dir)
	if err != nil {
		return err
	}

	if err := pdownload.Download(ctx); err != nil {
		return err
	}

	if err := pdownload.Utils.MergeFiles(pdownload.Procs); err != nil {
		return err
	}

	return nil
}

func (pdownload *Pdownload) Ready(args []string, targetDir string, timeout time.Duration) error {
	if err := pdownload.parseURL(args); err != nil {
		return err
	}

	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		return fmt.Errorf("target directory is not exist: %w", err)
	}
	pdownload.TargetDir = targetDir
	pdownload.timeout = timeout

	return nil
}

func (pdownload *Pdownload) parseURL(args []string) error {
	if len(args) > 1 {
		return errors.New("URL must be a single")
	}
	if len(args) < 1 {
		return errors.New("urls not found in the arguments passed")
	}

	for _, arg := range args {
		_, err := url.ParseRequestURI(arg)
		if err != nil {
			return err
		}
		pdownload.URL = arg
	}

	return nil
}
