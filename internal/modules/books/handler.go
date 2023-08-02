package books

import (
	"context"
	"database/sql"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Ensure that BooksServer implements the BooksServiceV1Server interface.
var _ BooksServiceV1Server = (*BooksServer)(nil)

// BooksServer is a struct that implements the BooksServiceV1Server interface and contains a repository for database operations
type BooksServer struct {
	UnimplementedBooksServiceV1Server

	Repository *Repository
}

// NewBookServer creates a new BooksServer with a repository that uses the provided MySQL database
func NewBookServer(mysql *sql.DB) *BooksServer {
	return &BooksServer{Repository: NewRepository(mysql)}
}

// GetAuthorsByBook is a handler that receives a book title and returns a list of authors
func (b *BooksServer) GetAuthorsByBook(ctx context.Context, req *GetAuthorsRequest) (*AuthorsListResponse, error) {
	if len(req.Title) == 0 {
		return nil, status.Error(codes.InvalidArgument, "title is required")
	}

	// Get the authors from the repository
	authors, err := b.Repository.GetAuthorsByBook(ctx, req.Title)
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("get authors: %v", err))
	}

	return &AuthorsListResponse{Authors: authors}, nil
}

// GetBooksByAuthor is a handler that receives an author name and returns a list of books
func (b *BooksServer) GetBooksByAuthor(ctx context.Context, req *GetBooksRequest) (*BooksListResponse, error) {
	if len(req.Name) == 0 {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}

	// Get the books from the repository
	books, err := b.Repository.GetBooksByAuthor(ctx, req.Name)
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("get books: %v", err))
	}

	return &BooksListResponse{Books: books}, nil
}
