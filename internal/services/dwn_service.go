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
)

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
