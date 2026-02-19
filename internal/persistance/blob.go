package persistance

import (
	"database/sql"
	"hish22/grpm/internal/middlewares"
	"log"
	"os"
	"time"
)

type blob struct {
	HashedLink []byte    `sql:"hashedlink"`
	Location   string    `sql:"location"`
	Expire     time.Time `sql:"expire"`
}

func rowBlobFromDb(link *[]byte) *sql.Row {
	db := middlewares.MetadataDBConn()
	defer db.Close()
	query := "SELECT hashedlink, location, expire FROM cache WHERE hashedlink=?"
	return db.QueryRow(query, link)
}

func FetchBlob(link *[]byte) (*blob, bool) {
	row := rowBlobFromDb(link)
	b := &blob{} // Result blob
	err := row.Scan(&b.HashedLink, &b.Location, &b.Expire)
	if err != nil {
		return b, false
	}
	return b, true
}

func ReadBlob(location *string) []byte {
	cachedJson, err := os.ReadFile(*location + ".json")
	if err != nil {
		log.Fatal("No such cache stored, ", err)
	}
	return cachedJson
}
