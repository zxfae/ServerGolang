package database

import (
	"database/sql"
)

// createTableCategories creates Categories table
func CreateTableCategories(db *sql.DB) error {
	table := `
	CREATE TABLE IF NOT EXISTS Categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		description TEXT NOT NULL
	);
	`
	_, err := db.Exec(table)
	return err
}

// createTablePosts creates the Posts table
func CreateTablePosts(db *sql.DB) error {
	table := `
	CREATE TABLE IF NOT EXISTS Posts (
		id TEXT NOT NULL PRIMARY KEY,
		username TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		categoryname TEXT NOT NULL,
		FOREIGN KEY(username) REFERENCES Users(nickname)
		FOREIGN KEY(categoryname) REFERENCES Categories(name)
	);`
	_, err := db.Exec(table)
	return err
}

func CreateTableComments(db *sql.DB) error {
	table := `
	CREATE TABLE IF NOT EXISTS "Comments" (
		id TEXT PRIMARY KEY, 
		postId TEXT NOT NULL,
		userId TEXT NOT NULL,
		date TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP),
		content TEXT NOT NULL,
		FOREIGN KEY (postId) REFERENCES Posts(id),
		FOREIGN KEY (userId) REFERENCES Users(id)
	);`
	_, err := db.Exec(table)
	return err
}

func CreateTableLikesComments(db *sql.DB) error {
	table := `
	CREATE TABLE IF NOT EXISTS "LikesComments" (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		userId TEXT NOT NULL,
		commentsId TEXT NOT NULL,
		date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		sentiment TEXT NOT NULL CHECK (sentiment IN ('love', 'hate')),
		FOREIGN KEY (commentsId) REFERENCES Comments(id),
		FOREIGN KEY (userId) REFERENCES Users(id),
		UNIQUE (userId, commentsId)
	);`
	_, err := db.Exec(table)
	return err
}

func CreateTablePostsCategories(db *sql.DB) error {
	table := `
	CREATE TABLE IF NOT EXISTS "PostCategories" (
		postId TEXT,
		categoryId INTEGER,
		PRIMARY KEY (postId, categoryId),
		FOREIGN KEY (postId) REFERENCES Posts(id),
		FOREIGN KEY (categoryId) REFERENCES Categories(id)
	);`
	_, err := db.Exec(table)
	return err
}

func CreateTablePostsLikes(db *sql.DB) error {
	table := `
	CREATE TABLE IF NOT EXISTS "PostsLike" (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		userId TEXT NOT NULL,
		postId TEXT NOT NULL,
		date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		sentiment TEXT NOT NULL CHECK (sentiment IN ('love', 'hate')),
		FOREIGN KEY (postId) REFERENCES Posts(id),
		FOREIGN KEY (userId) REFERENCES Users(id),
		UNIQUE (userId, postId)
	);`
	_, err := db.Exec(table)
	return err
}

// createTableUsers creates the Users table
func CreateTableUsers(db *sql.DB) error {
	table := `
 CREATE TABLE IF NOT EXISTS Users (
        id TEXT NOT NULL PRIMARY KEY,
        nickname TEXT NOT NULL UNIQUE,
        age INTEGER,
        gender TEXT,
        firstname TEXT,
        lastname TEXT,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
        created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        image TEXT NOT NULL DEFAULT 'backend/database/img/img.png'
    );
    `
	_, err := db.Exec(table)
	return err
}
