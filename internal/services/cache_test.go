package services

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"lukechampine.com/blake3"
)

func TestCacheServiceDispatch(t *testing.T) {
	data := "This is test data to cache."

	hash := blake3.Sum256([]byte(data))

	file, err := CreateFileService(
		CreatePathService(
			[]string{"tmp", "test"},
			[]string{strings.Join([]string{hex.EncodeToString(hash[:]), "gmd"}, ".")},
		),
		0755,
	)

	if err != nil {
		t.Error(err)
		return
	}

	cache := CreateCacheService(hash, file, []byte(data))

	err = cache.Dispatch(0)

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("Cache dispatched")

}

func TestCacheServiceFetchEntry(t *testing.T) {
	data := "This is test data to cache."

	hash := blake3.Sum256([]byte(data))

	entries, err := CreateFileService(
		CreatePathService(
			[]string{"tmp", "test"},
			[]string{strings.Join([]string{hex.EncodeToString(hash[:]), "gmd"}, ".")},
		),
		0755,
	)

	if err != nil {
		t.Error(err)
		return
	}

	cache, err := FetchCacheByFingerprint(entries, 0)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(cache.Data))
	fmt.Println(cache.Date)
	fmt.Println(hex.EncodeToString(cache.Fingerprint[:]))

}
