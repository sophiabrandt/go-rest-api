package book

// Info is the book model.
type Info struct {
	ID            string `db:"book_id" json:"book_id"`
	AuthorID      string `db:"author_id" json:"author_id,omitempty"`
	AuthorName    string `db:"author_name" json:"author_name"`
	Title         string `db:"title" json:"title"`
	PublishedDate string `db:"published_date" json:"published_date"`
	ImageUrl      string `db:"image_url" json:"image_url"`
	Description   string `db:"description" json:"description"`
}

// NewBook is the required data for creating a new book.
type NewBook struct {
	AuthorName    string `json:"author_name"`
	Title         string `json:"title"`
	PublishedDate string `json:"published_date"`
	ImageUrl      string `json:"image_url"`
	Description   string `json:"description"`
}
