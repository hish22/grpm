package services

/*
 * Download Service
 *
 * Provide ability to downstream content,
 * from outsource location into an upstream file.
 */

import (
	"context"
	"errors"
	"io"
	"os"

	"github.com/schollz/progressbar/v3"
)

// IN GENERAL THIS SERVICES + OTHERS NEEDS REFACTORING AND FIXING.
// NEEDS REFACTOR

type DownloadService struct {
	downstream *ApiService
	upstream   *fileService
}

// Create a new DownloadService. DownloadService expects a downstream ApiService and an upstream fileService.
func CreateDownloadService(downstream *ApiService, upstream *fileService) (*DownloadService, error) {
	if downstream == nil || upstream == nil {
		return nil, errors.New("Failed to create new DownloadService, downstream or upstream must not be nil.")
	}
	return &DownloadService{
		downstream,
		upstream,
	}, nil
}

// Operate into downstreaming content into file upstream with functional bar.
func (d *DownloadService) bardown() error {
	ctx, cancel := context.WithTimeout(context.Background(), d.downstream.Timeout)
	defer cancel()
	resp, err := d.downstream.Request(ctx)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	tf, err := d.upstream.NewTmpFile()
	if err != nil {
		return err
	}
	defer tf.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"Downloading",
	)

	mw := io.MultiWriter(tf, bar)
	_, err = io.Copy(mw, resp.Body)
	if err != nil {
		return err
	}

	filePath, err := d.upstream.path.String(d.upstream.index)
	if err != nil {
		return err
	}

	return os.Rename(tf.Name(), filePath)
}

// Operate into downstreaming content into file upstream.
func (d *DownloadService) down() error {
	ctx, cancel := context.WithTimeout(context.Background(), d.downstream.Timeout)
	defer cancel()
	resp, err := d.downstream.Request(ctx)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = d.upstream.ContingencySafeCopy(resp.Body)
	return err
}
