package services

/*
 * Cache service.
 *
 * Provide ability to retain data
 * for quick access in interval time.
 */

import (
	"encoding/gob"
	"errors"
	"time"
)

type CacheService struct {
	fingerprint [32]byte
	data        []byte
	stream      *fileService
}

type CacheEntry struct {
	Fingerprint [32]byte
	Date        time.Time
	Data        []byte
}

func CreateCacheEntry(fingerprint [32]byte, date time.Time, data []byte) *CacheEntry {
	return &CacheEntry{
		fingerprint,
		date,
		data,
	}
}

// Create new CacheService. CacheService expects a fingerprint, fileService, and data in bytes to retain.
func CreateCacheService(fingerprint [32]byte, file *fileService, data []byte) *CacheService {
	return &CacheService{
		fingerprint,
		data,
		file,
	}
}

// Apply new cache to indexed file stream.
func (c CacheService) Dispatch() error {
	mdf, err := c.stream.NewRWFile()
	if err != nil {
		return err
	}
	gobe := gob.NewEncoder(mdf)
	err = gobe.Encode(CreateCacheEntry(
		c.fingerprint,
		time.Now(),
		c.data,
	))
	if err != nil {
		return err
	}
	return nil
}

/* -------------------------Related Cache Functions------------------------------ */

// Fetch applied cache content by fingerprint. FetchCacheByFingerprint expect
// cache streams container to have the requested cache entries and fileIndex.
func FetchCacheByFingerprint(container *fileService) (*CacheEntry, error) {
	if container == nil {
		return nil, errors.New("Failed to fetch cahce stream, container is nil.")
	}

	var fetchedCacheEntry *CacheEntry

	cacheFileData, err := container.OpenROFile()

	if err != nil {
		return nil, err
	}

	defer cacheFileData.Close()

	gobd := gob.NewDecoder(cacheFileData)
	err = gobd.Decode(&fetchedCacheEntry)

	if err != nil {
		return nil, err
	}

	return fetchedCacheEntry, nil
}
