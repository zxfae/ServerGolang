package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(databaseName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", databaseName)
	if err != nil {
		log.Printf("Error opening database: %v", err)
		return nil, err
	}
	return db, nil
}

// CreateTableUser for register
func CreateTableUser(db *sql.DB) error {
	createUserTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		nickname TEXT,
		age INTEGER,
		gender TEXT,
		firsname TEXT,
		LASTNAME TEXT,
		email TEXT,
		password TEXT
	);
	CREATE TABLE IF NOT EXISTS POSTS(
		userid INTEGER NOT NULL,
		postid INTEGER PRIMARY KEY AUTOINCREMENT,
		img TEXT DEFAULT NULL,
		commentid INTEGER DEFAULT NULL,
		username TEXT NOT NULL,
		role TEXT NOT NULL,
		category TEXT DEFAULT NULL,
		categoryB TEXT DEFAULT NULL,
		userpfp TEXT NOT NULL,
		content TEXT NOT NULL,
		postdate DATE NOT NULL,
		FOREIGN KEY(category) REFERENCES CATEGORIES(name),
		FOREIGN KEY(userid) REFERENCES USERS(id),
		FOREIGN KEY(commentid) REFERENCES POSTS(id)
		);
	CREATE TABLE IF NOT EXISTS CATEGORIES(
		name TEXT UNIQUE,
		posts INTEGER
	);
	CREATE TABLE IF NOT EXISTS LikesDislikes (
		id INTEGER PRIMARY KEY,
		userid INTEGER NOT NULL,
		postid INTEGER NOT NULL,
		likestatus TEXT NOT NULL CHECK(likestatus IN ('like', 'dislike')),
		FOREIGN KEY(userid) REFERENCES USERS(id),
		FOREIGN KEY(postid) REFERENCES POSTS(id)
	 );
	`

	_, err := db.Exec(createUserTable)
	if err != nil {
		log.Printf("Error creating table: %v", err)
	}
	return err
}

// func InsertUser(db *sql.DB, username string, password string) error {
// 	statement, err := db.Prepare("INSERT INTO users (nickname, age, gender, firstname, lastname, email, password) VALUES (?, ?, ?, ?, ?, ?, ?)")
// 	if err != nil {
// 		log.Printf("Error preparing insert: %v", err)
// 		return err
// 	}
// 	defer statement.Close()

// 	_, err = statement.Exec(username, password)
// 	if err != nil {
// 		log.Printf("Error inserting user: %v", err)
// 	}
// 	return err
// }
