package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initializes the connection to the database
func InitDB() (*sql.DB, error) {
	dbFilePath := "./backend/database/database.db"

	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		return nil, err
	}

	// Configure database connection settings
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(0)

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// InitMainDB initializes the main database
func InitMainDB() {
	var err error

	Db, err = InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	//defer db.Close()

	if err = createTableUsers(Db); err != nil {
		log.Fatalf("Failed to create Users table: %v", err)
	}

	if err = createTableCategories(Db); err != nil {
		log.Fatalf("Failed to create Categories table : %v", err)
	}

	if err = createTablePosts(Db); err != nil {
		log.Fatalf("Failed to create Posts table: %v", err)
	}

	if err = createTableComments(Db); err != nil {
		log.Fatalf("Failed to create Posts table: %v", err)
	}

	if err = createTableLikesComments(Db); err != nil {
		log.Fatal("Failed to create LikesComments table")
	}

	if err = createTablePostsCategories(Db); err != nil {
		log.Fatal("Failed to create PostCategories table")
	}

	if err = createTablePostsLikes(Db); err != nil {
		log.Fatal("Failed to create PostLikes table")
	}
	log.Println("Database initialized, test user and test post inserted successfully")
}
