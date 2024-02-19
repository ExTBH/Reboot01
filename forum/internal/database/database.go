package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"sandbox/internal/structs"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Connects to the database, if an error happens exists with status 1
func Connect(dbPath string) error {
	// sql.Open wont error if file not found
	fi, err := os.Stat(dbPath)
	if err != nil || fi.IsDir() {
		return errors.New("database file not found")
	}
	dsn := fmt.Sprintf("file:%s?cache=shared&mode=rwc", dbPath)
	ldb, err := sql.Open("sqlite3", dsn)
	if err != nil {
		msg := fmt.Sprintf("can't connect to database: %s", err.Error())
		return errors.New(msg)
	}
	db = ldb
	return nil
}

// Called to close the connection and finish up
func Close() error {
	return db.Close()
}

// checks if a value exists on a certain table
func CheckExistance(tablename, columnname, value string) (bool, error) {
	// Prepare the SQL statement with a placeholder
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ?", tablename, columnname)

	stmt, err := db.Prepare(query)

	if err != nil {
		return false, err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the count
	var count int
	err = stmt.QueryRow(value).Scan(&count)
	if err != nil {
		return false, err
	}

	if count > 0 {
		// value exists in the database
		return true, nil
	}

	// value does not exist in the database
	return false, nil
}

// SearchContent searches for posts that contain the given keyword in the title or message.
func SearchContent(keyword string) ([]structs.Post, error) {
	var posts []structs.Post

	// Prepare the SQL statement for searching posts
	stmt, err := db.Prepare(`SELECT id, user_id, parent_id, title, message, image_id, time, like_count, 
							dislike_count, love_count, haha_count, skull_count
							FROM Post 
							WHERE title LIKE ? OR message LIKE ?`)
	if err != nil {
		return posts, err
	}
	defer stmt.Close()

	// Execute the SQL statement to search posts
	rows, err := stmt.Query("%"+keyword+"%", "%"+keyword+"%")
	if err != nil {
		return posts, err
	}
	defer rows.Close()

	return getPostsHelper(rows)
}
