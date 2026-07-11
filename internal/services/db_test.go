package services

import (
	"context"
	"fmt"
	"testing"
	"time"

	_ "modernc.org/sqlite" // Import for side-effects
)

func TestDatabaseServiceMultiQueryWithContext(t *testing.T) {

	dbc, err := CreateDatabaseService("sqlite", "/home/pling77/Desktop/REAL_PROJECTS/grpm/test.db")

	if err != nil {
		t.Error("Failed to create database service", "error", err)
		return
	}

	var id int
	var name string

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	rows, err := dbc.QueryRowsWithContext(ctx, "SELECT * FROM users")

	if err != nil {
		t.Error("Failed to query the table", "error", err)
		return
	}

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			t.Error("Failed to fetch a row from the table", "error", err)
			return
		}
		fmt.Println(id, name)
	}

}

func TestDatabaseServiceMultiQuery(t *testing.T) {
	dbc, err := CreateDatabaseService("sqlite", "/home/pling77/Desktop/REAL_PROJECTS/grpm/test.db")

	if err != nil {
		t.Error("Failed to create database service", "error", err)
		return
	}

	var id int
	var name string

	rows, err := dbc.QueryRows("SELECT * FROM users")

	if err != nil {
		t.Error("Failed to query the table", "error", err)
		return
	}

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			t.Error("Failed to fetch a row from the table", "error", err)
			return
		}
		fmt.Println(id, name)
	}
}

func TestDatabaseServiceContextQuery(t *testing.T) {
	dbc, err := CreateDatabaseService("sqlite", "/home/pling77/Desktop/REAL_PROJECTS/grpm/test.db")

	if err != nil {
		t.Error("Failed to create database service", "error", err)
		return
	}

	var id int
	var name string

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err = dbc.QueryRowWithContext(ctx, "SELECT * FROM users WHERE id=$1", []any{2}, &id, &name)

	if err != nil {
		t.Error("Failed to query the table", "error", err)
		return
	}

	fmt.Println(id, name)

}

func TestDatabaseServiceQuery(t *testing.T) {
	dbc, err := CreateDatabaseService("sqlite", "/home/pling77/Desktop/REAL_PROJECTS/grpm/test.db")

	if err != nil {
		t.Error("Failed to create database service")
		return
	}

	var id int
	var name string

	err = dbc.QueryRow("SELECT id,name FROM users WHERE id=$1", []any{1}, &id, &name)

	if err != nil {
		t.Error("Failed to query table users", err)
		return
	}

	fmt.Println("User ID: ", id)
	fmt.Println("User Name: ", name)

}
