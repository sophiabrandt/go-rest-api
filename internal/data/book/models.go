package book

import "strings"

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

type Infos []Info

// NewBook is the required data for creating a new book.
type NewBook struct {
	AuthorName    string `json:"author_name" validate:"required,alpha_space"`
	Title         string `json:"title" validate:"required"`
	PublishedDate string `json:"published_date" validate:"required,date"`
	ImageUrl      string `json:"image_url" validate:"required,url,max=255"`
	Description   string `json:"description" validate:"required,max=255"`
}

// ToDto formats the database model for client consumption.
func (b Info) ToDto() Info {
	return Info{
		ID:            b.ID,
		AuthorID:      b.AuthorID,
		AuthorName:    properTitle(b.AuthorName),
		Title:         properTitle(b.Title),
		PublishedDate: b.PublishedDate,
		ImageUrl:      b.ImageUrl,
		Description:   b.Description,
	}
}

// ToDto formats the database model for client consumption.
func (bs Infos) ToDto() []Info {
	books := make([]Info, len(bs))
	for i, b := range bs {
		books[i] = b.ToDto()
	}
	return books
}

// titleCase the input.
func properTitle(input string) string {
	words := strings.Fields(input)
	smallwords := " a an on the of to "
	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") && word != words[0] {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}
