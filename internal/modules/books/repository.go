package books

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

// Repository is a struct that contains a MySQL database for operations
type Repository struct {
	DB *sql.DB
}

// NewRepository creates a new Repository with the provided MySQL database
func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

// GetAuthorsByBook is a method to fetches the authors of a book title from the MySQL repository.
func (r *Repository) GetAuthorsByBook(ctx context.Context, title string) ([]*Author, error) {
	rows, err := r.DB.
		QueryContext(
			ctx,
			`SELECT a.name
			FROM authors a
			JOIN author_book ab ON ab.author_id = a.id
			JOIN books b ON ab.book_id = b.id
			WHERE b.title = ?
			ORDER BY a.name ASC`,
			title,
		)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Scan the results into a slice of authors
	var authors []*Author
	for rows.Next() {
		var author Author
		if err := rows.Scan(&author.Name); err != nil {
			return nil, err
		}
		authors = append(authors, &author)
	}

	// If no authors were found, return an error.
	if authors == nil {
		return nil, fmt.Errorf("no authors found for the book: %s", title)
	}

	return authors, nil
}

// GetBooksByAuthor is a method to fetches the books by an author name from the MySQL repository.
func (r *Repository) GetBooksByAuthor(ctx context.Context, name string) ([]*Book, error) {
	rows, err := r.DB.
		QueryContext(
			ctx,
			`SELECT b.title
			FROM books b
			JOIN author_book ab ON ab.book_id = b.id
			JOIN authors a ON ab.author_id = a.id
			WHERE a.name = ?
			ORDER BY b.title ASC`,
			name,
		)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Scan the results into a slice of books
	var books []*Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.Title); err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	// If no books were found, return an error.
	if books == nil {
		return nil, fmt.Errorf("no books found for the author: %s", name)
	}

	return books, nil
}
