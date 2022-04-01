package internals

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
