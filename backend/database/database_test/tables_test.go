package databasetest

import (
	"database/sql"
	"real-time-backend/backend/database"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func OpenTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open test database: %s", err)
	}
	return db
}

func TestCreateTableCategories(t *testing.T) {
	db := OpenTestDB(t)
	defer db.Close()

	if err := database.CreateTableCategories(db); err != nil {
		t.Errorf("error was not expected while creating table: %s", err)
	}
}

func TestCreateTablePosts(t *testing.T) {
	db := OpenTestDB(t)
	defer db.Close()

	if err := database.CreateTablePosts(db); err != nil {
		t.Errorf("error was not expected while creating table: %s", err)
	}
}

func TestCreateTableComments(t *testing.T) {
	db := OpenTestDB(t)
	defer db.Close()

	if err := database.CreateTableComments(db); err != nil {
		t.Errorf("error was not expected while creating table: %s", err)
	}
}

func TestCreateTableLikesComments(t *testing.T) {
	db := OpenTestDB(t)
	defer db.Close()

	if err := database.CreateTableLikesComments(db); err != nil {
		t.Errorf("error was not expected while creating table: %s", err)
	}
}

func TestCreateTablePostsCategories(t *testing.T) {
	db := OpenTestDB(t)
	defer db.Close()

	if err := database.CreateTablePostsCategories(db); err != nil {
		t.Errorf("error was not expected while creating table: %s", err)
	}
}

func TestCreateTablePostsLikes(t *testing.T) {
	db := OpenTestDB(t)
	defer db.Close()

	if err := database.CreateTablePostsLikes(db); err != nil {
		t.Errorf("error was not expected while creating table: %s", err)
	}
}

func TestCreateTableUsers(t *testing.T) {
	db := OpenTestDB(t)
	defer db.Close()

	if err := database.CreateTableUsers(db); err != nil {
		t.Errorf("error was not expected while creating table: %s", err)
	}
}
