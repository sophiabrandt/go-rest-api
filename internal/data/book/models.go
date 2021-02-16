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
	AuthorName    string `json:"author_name" validate:"required,max=255"`
	Title         string `json:"title" validate:"required,max=255"`
	PublishedDate string `json:"published_date" validate:"required,max=255"`
	ImageUrl      string `json:"image_url" validate:"required,url,max=255"`
	Description   string `json:"description" validate:"required,max=255"`
}
