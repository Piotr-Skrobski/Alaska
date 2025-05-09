package repositories

import (
	"database/sql"
	"fmt"
)

func MigrateReviewTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS reviews (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    movie_id VARCHAR(30) NOT NULL,
    rating INTEGER NOT NULL CHECK (rating BETWEEN 1 AND 10),
    comment TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create review table: %w", err)
	}
	return nil
}
