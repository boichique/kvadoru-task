package tests

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/boichique/kvadoru_task/internal/modules/books"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TestBooksServer_GetAuthorsByBook tests the GetAuthorsByBook method of the Books server
func TestBooksServer_GetAuthorsByBook(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) // Create a mock database
	defer db.Close()                                                                  // Ensure the database will be closed after the test finishes

	repo := books.NewRepository(db)        // Create a new repository with the mock database
	server := books.NewBookServer(repo.DB) // Create a new Books server with the repository
	ctx := context.Background()
	bookTitle := "Test Book" // Define a book title for testing

	query := `SELECT a.name 
		FROM authors a 
		JOIN author_book ab ON ab.author_id = a.id 
		JOIN books b ON ab.book_id = b.id 
		WHERE b.title = ? 
		ORDER BY a.name ASC` // Define the SQL query to be mocked

	// Define a subtest for a valid request
	t.Run("valid request", func(t *testing.T) {
		expectedAuthors := []*books.Author{
			{Name: "Author1"},
			{Name: "Author2"},
			{Name: "Author3"},
		}

		rows := sqlmock.
			NewRows([]string{"name"}).
			AddRow(expectedAuthors[0].Name).
			AddRow(expectedAuthors[1].Name).
			AddRow(expectedAuthors[2].Name)

		mock.ExpectQuery(query).
			WithArgs(bookTitle).
			WillReturnRows(rows)

		req := &books.GetAuthorsRequest{Title: bookTitle}
		resp, err := server.GetAuthorsByBook(ctx, req)

		require.NoError(t, err)
		require.Equal(t, expectedAuthors, resp.Authors)
	})

	// Define a subtest for when no authors are found
	t.Run("no authors found", func(t *testing.T) {
		mock.ExpectQuery(query).
			WithArgs(bookTitle).
			WillReturnError(sql.ErrNoRows)

		req := &books.GetAuthorsRequest{Title: bookTitle}
		_, err := server.GetAuthorsByBook(ctx, req)

		require.Equal(t, codes.NotFound, status.Code(err))
	})

	// Define a subtest for an invalid request
	t.Run("invalid request", func(t *testing.T) {
		req := &books.GetAuthorsRequest{Title: ""}
		_, err := server.GetAuthorsByBook(ctx, req)

		require.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	err := mock.ExpectationsWereMet() // Check if all mock expectations were met
	require.NoError(t, err)
}

// TestBooksServer_GetBooksByAuthor tests the GetBooksByAuthor method of the Books server
func TestBooksServer_GetBooksByAuthor(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) // Create a mock database
	defer db.Close()                                                                  // Ensure the database will be closed after the test finishes

	repo := books.NewRepository(db)        // Create a new repository with the mock database
	server := books.NewBookServer(repo.DB) // Create a new Books server with the repository
	ctx := context.Background()
	authorName := "Author1" // Define an author name for testing

	query := `SELECT b.title 
		FROM books b 
		JOIN author_book ab ON ab.book_id = b.id 
		JOIN authors a ON ab.author_id = a.id 
		WHERE a.name = ? 
		ORDER BY b.title ASC` // Define the SQL query to be mocked

	// Define a subtest for a valid request
	t.Run("valid request", func(t *testing.T) {
		expectedBooks := []*books.Book{
			{Title: "Book1"},
			{Title: "Book2"},
			{Title: "Book3"},
		}

		rows := sqlmock.
			NewRows([]string{"title"}).
			AddRow(expectedBooks[0].Title).
			AddRow(expectedBooks[1].Title).
			AddRow(expectedBooks[2].Title)

		mock.ExpectQuery(query).
			WithArgs(authorName).
			WillReturnRows(rows)

		req := &books.GetBooksRequest{Name: authorName}
		resp, err := server.GetBooksByAuthor(ctx, req)

		require.NoError(t, err)
		require.Equal(t, expectedBooks, resp.Books)
	})

	// Define a subtest for when no books are found
	t.Run("no books found", func(t *testing.T) {
		mock.ExpectQuery(query).
			WithArgs(authorName).
			WillReturnError(sql.ErrNoRows)

		req := &books.GetBooksRequest{Name: authorName}
		_, err := server.GetBooksByAuthor(ctx, req)

		require.Equal(t, codes.NotFound, status.Code(err))
	})

	// Define a subtest for an invalid request
	t.Run("invalid request", func(t *testing.T) {
		req := &books.GetBooksRequest{Name: ""}
		_, err := server.GetBooksByAuthor(ctx, req)

		require.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	err := mock.ExpectationsWereMet() // Check if all mock expectations were met
	require.NoError(t, err)
}
