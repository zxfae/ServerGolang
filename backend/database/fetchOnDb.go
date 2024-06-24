package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type QueryExecutor func(*sql.Rows) (interface{}, error)

func FetchDb(query string, executor QueryExecutor) ([]interface{}, error) {
	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []interface{}

	for rows.Next() {
		result, err := executor(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func ConvertResults[T any](results []interface{}) ([]T, error) {
	converted := make([]T, len(results))
	for i, result := range results {
		item, ok := result.(T)
		if !ok {
			return nil, fmt.Errorf("failed to convert result to type %T", result)
		}
		converted[i] = item
	}
	return converted, nil
}
