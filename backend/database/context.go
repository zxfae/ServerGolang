package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"real-time-backend/backend/modals"

	"golang.org/x/crypto/bcrypt"
)

func CreatePost(post *modals.Post) error {
	query := `INSERT INTO Posts (id, userId, username, created_at, title, description, categoryname) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := Db.Exec(query, post.Id, post.UserId, post.Username, post.Creation, post.Title, post.Description, post.Name)
	return err
}

// Login function, query necessary options,
// Handling error, more details to server and simple message to client
func LoginUser(ctx context.Context, nickname, password string) error {
	var hashedPassword string
	query := "SELECT password FROM Users WHERE nickname = ? OR email = ?"

	err := Db.QueryRowContext(ctx, query, nickname, nickname).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("authentification failed")
		}
		log.Printf("Failed to query user: %v", err)
		return fmt.Errorf("authentification failed")
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Printf("Incorrect password for user %v", nickname)
		return fmt.Errorf("authentification failed")
	}

	log.Printf("User %s logged in successfully", nickname)
	return nil
}

// Hashing password, && UUID generation for id
// Using context
func RegisterUser(ctx context.Context, nickname, age, gender, firstname, lastname, email, password string) error {

	//Check nickname && email doubles
	existsNickname, err := CheckNicknameExists(ctx, nickname)
	if err != nil {
		log.Printf("Username : %v already exists", err)
		return fmt.Errorf("failed to check username: %w", err)
	}
	if existsNickname {
		return fmt.Errorf("nickname already exists")
	}

	existsEmail, err := CheckEmailExists(ctx, email)
	if err != nil {
		log.Printf("Email : %v already exists", err)
		return fmt.Errorf("failed to check email: %w", err)

	}
	if existsEmail {
		return fmt.Errorf("email already exists")
	}

	if len(password)-1 < 6 {
		return fmt.Errorf("password too short")
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		log.Printf("Failed to hash password %v", err)
		return fmt.Errorf("failed to registration")
	}
	query := "INSERT INTO Users (id, nickname, age, gender, firstname, lastname, email, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	//New transaction
	tx, err := Db.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//Cancel transaction if not terminated
	defer tx.Rollback()
	//Context execution
	_, err = tx.ExecContext(ctx, query, GenerateUUID(), nickname, age, gender, firstname, lastname, email, hashedPassword)
	if err != nil {
		return nil
	}
	//Transaction validation
	err = tx.Commit()
	if err != nil {
		return nil
	}
	fmt.Println("User registered successfully")

	return nil
}
