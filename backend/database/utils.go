package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const DbFilePathEnv = ""

var Db *sql.DB

func GenerateUUID() string {
	return uuid.New().String()
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}
	return string(hashedPassword), nil
}

func CheckNicknameExists(ctx context.Context, nickname string) (bool, error) {
	query := "SELECT COUNT(*) FROM Users WHERE nickname = ?"
	var count int
	err := Db.QueryRowContext(ctx, query, nickname).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to execute query: %w", err)
	}
	return count > 0, nil
}

func CheckEmailExists(ctx context.Context, email string) (bool, error) {
	query := "SELECT COUNT(*) FROM Users WHERE email = ?"
	var count int
	err := Db.QueryRowContext(ctx, query, email).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to execute query: %w", err)
	}
	return count > 0, nil
}
