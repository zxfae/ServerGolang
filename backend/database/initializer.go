package database

import (
	"database/sql"
	"fmt"
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

	tableCreationFuncs := []struct {
		name string
		fn   func(*sql.DB) error
	}{
		{"Categories", CreateTableCategories},
		{"Posts", CreateTablePosts},
		{"Comments", CreateTableComments},
		{"LikesComments", CreateTableLikesComments},
		{"PostCategories", CreateTablePostsCategories},
		{"PostsLikes", CreateTablePostsLikes},
		{"Users", CreateTableUsers},
	}

	for _, tableFunc := range tableCreationFuncs {
		if err := tableFunc.fn(Db); err != nil {
			fmt.Printf("Error creating table %s: %s\n", tableFunc.name, err)
		} else {
			fmt.Printf("Table %s created successfully\n", tableFunc.name)
		}
	}

	//defer db.Close()

	if err = CreateTableUsers(Db); err != nil {
		log.Fatalf("Failed to create Users table: %v", err)
	}

	if err = CreateTableCategories(Db); err != nil {
		log.Fatalf("Failed to create Categories table : %v", err)
	}

	if err = CreateTablePosts(Db); err != nil {
		log.Fatalf("Failed to create Posts table: %v", err)
	}

	if err = CreateTableComments(Db); err != nil {
		log.Fatalf("Failed to create Posts table: %v", err)
	}

	if err = CreateTableLikesComments(Db); err != nil {
		log.Fatal("Failed to create LikesComments table")
	}

	if err = CreateTablePostsCategories(Db); err != nil {
		log.Fatal("Failed to create PostCategories table")
	}

	if err = CreateTablePostsLikes(Db); err != nil {
		log.Fatal("Failed to create PostLikes table")
	}
	log.Println("Database initialized, test user and test post inserted successfully")
}
