package internals

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

func Get_Posts(config Config) ([]Post, error) {
	db, err := sqlx.Connect("postgres", config.DatabaseConnectionString())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	defer db.Close()

	var posts []Post
	err = db.Select(&posts, "SELECT post_id, title, content, author_id, name, email FROM posts, users WHERE user_id = author_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		return nil, err
	}

	return posts, nil
}
