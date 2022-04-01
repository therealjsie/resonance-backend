package internals

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

type Author struct {
	Id    int `db:"author_id"`
	Name  string
	Email string
}

type Post struct {
	Id      int `db:"post_id"`
	Title   string
	Content string
	Author
}

func Hello_world() ([]Post, error) {
	config := ReadConfig()

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
