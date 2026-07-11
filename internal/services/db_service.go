package services

/*
 * DB service.
 *
 * Provide an api to deal with different
 * database providers such as sqlite.
 */

import (
	"context"
	"database/sql"

	_ "modernc.org/sqlite" // Import for side-effects
)

type DatabaseService struct {
	connection *sql.DB
	driver     string
	file       string
}

// Create new database service by providing driver name and path to the (*.db) file.
func CreateDatabaseService(driver string, file string) (*DatabaseService, error) {
	conn, err := sql.Open(driver, file)
	if err != nil {
		return nil, err
	}
	return &DatabaseService{
		connection: conn,
		driver:     driver,
		file:       file,
	}, nil
}

// QueryRow run a query to fetch a row from a database file.
func (db DatabaseService) QueryRow(q string, args []any, vals ...any) error {
	row := db.connection.QueryRow(q, args...)
	err := row.Scan(vals...)
	return err
}

// QueryRowWithContext run a query to fetch a row from a database file, with specifying the context.
func (db DatabaseService) QueryRowWithContext(ctx context.Context, q string, args []any, vals ...any) error {
	row := db.connection.QueryRowContext(ctx, q, args...)
	err := row.Scan(vals...)
	return err
}

// QueryRows run a query to fetch multiple rows.
func (db DatabaseService) QueryRows(q string, args ...any) (*sql.Rows, error) {
	rows, err := db.connection.Query(q, args...)
	return rows, err
}

// QueryRowsWithContext run a query to fetch multiple rows, with specifying the context.
func (db DatabaseService) QueryRowsWithContext(ctx context.Context, q string, args ...any) (*sql.Rows, error) {
	rows, err := db.connection.QueryContext(ctx, q, args...)
	return rows, err
}
