package persistance

import (
	"database/sql"
	"encoding/hex"
	"hish22/grpm/internal/config"
	"hish22/grpm/internal/serialization"
	"hish22/grpm/internal/util"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	charmlog "github.com/charmbracelet/log"
	"lukechampine.com/blake3"
	_ "modernc.org/sqlite" // Import for side-effects
)

const (
	CacheRfootLocation = "cache"
)

func ChacheRootLocation(append string) string {
	home, err := util.HomeDir()
	if err != nil {
		charmlog.Error("Failed to fetch home directory", "error", err)
	}
	switch runtime.GOOS {
	case "linux":
		return filepath.Join(home, ".cache", append)
	case "windows":
		return filepath.Join(home, ".cache", append)
	default:
		return ""
	}
}

func MetadataDbLocation() string {
	return filepath.Join(config.LocalConfigDirPath().String(), "metadata.db")
}

func NewCache(link string, response any) {
	hash := blake3.Sum256([]byte(link))
	blakeHexVersion := hex.EncodeToString(hash[:])
	blob := blob{
		HashedLink: []byte(blakeHexVersion),
		Location:   filepath.Join(ChacheRootLocation("grpm"), blakeHexVersion),
		Expire:     time.Now().AddDate(0, 0, 1),
	}
	chunk := serialization.JsonSerialization(response)
	metedataEntry(&blob)
	storeChunk(&blob, chunk)
}

func DeleteCache(link *[]byte) {
	// Delete metadata entry from db
	db := OpenMetadataDB()
	defer db.Close()
	query := "DELETE FROM cache WHERE hashedlink=?"
	_, err := db.Exec(query, link)
	if err != nil {
		log.Fatal("Cant delete cache of ", link, ", ", err)
	}
	// delete file from cache folder
	cacheFilePath := filepath.Join(ChacheRootLocation("grpm"), string(*link)+".json")
	err = os.Remove(cacheFilePath)
	if err != nil {
		log.Fatal("Can't remove cache .json file, ", err)
	}
}

func FetchFromCache(response any, link string) bool {
	hashed := blake3.Sum256([]byte(link))
	BlakeHexVersion := hex.EncodeToString(hashed[:])
	text_hashed := []byte(BlakeHexVersion)
	blob, exists := FetchBlob(&text_hashed)
	if exists {
		if blob.Expire.After(time.Now()) {
			buf := ReadBlob(&blob.Location)
			serialization.JsonUnserialization(buf, &response)
			return true
		} else {
			charmlog.Info("clearing cache")
			DeleteCache(&text_hashed)
		}
	}
	return false
}

func OpenMetadataDB() *sql.DB {
	db, err := sql.Open("sqlite", MetadataDbLocation())
	if err != nil {
		log.Fatal("Can't open/create metadata.db, ", err)
	}
	return db
}

func metedataEntry(blob *blob) {
	db := OpenMetadataDB()
	defer db.Close()
	ddlQuery := "CREATE TABLE IF NOT EXISTS cache (id INT PRIMARY KEY,hashedlink TEXT UNIQUE,location TEXT,expire DATE);"
	_, err := db.Exec(ddlQuery)
	if err != nil {
		charmlog.Error("Failed to create cache table, ", err)
	}
	dmlQuery := "INSERT INTO cache(hashedlink,location,expire) VALUES (?,?,?)"
	_, err = db.Exec(dmlQuery, blob.HashedLink, blob.Location, blob.Expire)
	if err != nil {
		charmlog.Error("Failed to insert cache metadata into cache table, ", err)
	}
}

func storeChunk(blob *blob, chunk []byte) {
	// Create the cache folder
	if os.MkdirAll(ChacheRootLocation("grpm"), 0755) != nil {
		charmlog.Error("Failed to create cache directory")
	}
	// Write json blob into .json file
	if err := os.WriteFile(blob.Location+".json", chunk, 0644); err != nil {
		charmlog.Error("Failed to create a cache blob, ", err)
	}
}
