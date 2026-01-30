package cache

import (
	"database/sql"
	"encoding/hex"
	"hish22/grpm/internal/config"
	"hish22/grpm/internal/serialization"
	"log"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3" // Import for side-effects
	"lukechampine.com/blake3"
)

const (
	CacheRootLocation = "cache"
)

func MetadataDbLocation() string {
	return filepath.Join(config.LocalConfigDirPath(), "metadata.db")
}

func NewCache(link string, response any) {
	hash := blake3.Sum256([]byte(link))
	blakeHexVersion := hex.EncodeToString(hash[:])
	blob := blob{
		HashedLink: []byte(blakeHexVersion),
		Location:   filepath.Join(CacheRootLocation, blakeHexVersion),
		Expire:     time.Now().AddDate(0, 0, 1),
	}
	chunk := serialization.JsonSerialization(response)
	metedataEntry(&blob)
	storeChunk(&blob, &chunk)
}

func openMetadataDB() *sql.DB {
	db, err := sql.Open("sqlite3", MetadataDbLocation())
	if err != nil {
		log.Fatal("Can't open/create metadata.db, ", err)
	}
	return db
}

func metedataEntry(blob *blob) {
	db := openMetadataDB()
	defer db.Close()
	ddlQuery := "CREATE TABLE IF NOT EXISTS cache (id INT PRIMARY KEY,hashedlink TEXT UNIQUE,location TEXT,expire DATE);"
	_, err := db.Exec(ddlQuery)
	if err != nil {
		log.Fatal("Can't create cache table, ", err)
	}
	dmlQuery := "INSERT INTO cache(hashedlink,location,expire) VALUES (?,?,?)"
	_, err = db.Exec(dmlQuery, blob.HashedLink, blob.Location, blob.Expire)
	if err != nil {
		log.Fatal("Can't insert entry into cache table, ", err)
	}
}

func storeChunk(blob *blob, chunk *[]byte) {
	if err := os.WriteFile(blob.Location+".json", *chunk, 0644); err != nil {
		log.Fatal("Can't create a blob cache, ", err)
	}
}
