package middlewares

import (
	"database/sql"
	"hish22/grpm/internal/config"
	"path/filepath"

	charmlog "github.com/charmbracelet/log"
)

type DB struct {
	Driver string
	File   string
}

/*  Opens a database */
func (db DB) Conn() *sql.DB {
	conn, err := sql.Open(db.Driver, db.File)
	if err != nil {
		charmlog.Fatal("Failed to open/create database", "error", err, "File", db.File)
	}
	return conn
}

/* Opens metadata database */
func MetadataDBConn() *sql.DB {
	db := DB{
		Driver: "sqlite",
		File:   filepath.Join(config.LocalConfigDirPath().String(), "metadata.db"),
	}
	return db.Conn()
}
