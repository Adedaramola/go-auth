package database

import (
	"time"
)

type User struct {
	ID        uint64     `db:"id"`
	Name      string     `db:"name"`
	Email     string     `db:"email"`
	Password  string     `db:"password"`
	CreatedAt *time.Time `db:"created_at"`
}

// func (db *DB) InsertUser(ctx context.Context) {
// 	query := `
// 		INSERT INTO users (name, email, password)
// 		VALUES ($1, $2, $3)
// 		RETURNING id`

// 	var id int

// 	err := db.GetContext(ctx, &id, query)
// }
