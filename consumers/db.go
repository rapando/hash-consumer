package consumers

import (
	"database/sql"
	"fmt"
)

func save(protocol, rawPassword, hashedPassword string, db *sql.DB) (err error) {
	query := fmt.Sprintf("INSERT INTO %s (raw_password, hashed_password) VALUES (?, ?)", protocol)
	_, err = db.Exec(query, rawPassword, hashedPassword)
	return err
}
