package persistance

import (
	"context"
	"database/sql"
	"hish22/grpm/internal/middlewares"
	"os"
	"time"

	charmlog "github.com/charmbracelet/log"
)

type blob struct {
	HashedLink []byte    `sql:"hashedlink"`
	Location   string    `sql:"location"`
	Expire     time.Time `sql:"expire"`
}

func rowBlobFromDb(link *[]byte) (*sql.Row, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	db := middlewares.MetadataDBConn()
	defer db.Close()
	query := "SELECT hashedlink, location, expire FROM cache WHERE hashedlink=?"
	return db.QueryRowContext(ctx, query, link), cancel
}

func FetchBlob(link *[]byte) (*blob, bool) {
	row, cancel := rowBlobFromDb(link)
	defer cancel()
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
		charmlog.Error("Failed to fetch specified cache", "error", err)
	}
	return cachedJson
}
