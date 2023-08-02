package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/boichique/kvadoru_task/internal/modules/books"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	// Command line flags
	addr := flag.String("addr", "localhost:9000", "Address of the gRPC server")
	authorFlag := flag.String("author", "", "Name of the author")
	bookFlag := flag.String("book", "", "Title of the book")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.DialContext(context.Background(), *addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		failOnError(err, "connect to gRPC server") // If there is an error connecting to the server, log the error and exit
	}
	defer conn.Close() // Ensure the connection is closed when main function returns

	c := books.NewBooksServiceV1Client(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second) // Create a context with a timeout
	defer cancel()                                                        // Ensure the context is cancelled when main function returns

	// If both author and book flags are set, log an error and exit
	if *authorFlag != "" && *bookFlag != "" {
		failOnError(errors.New("wrong request"), "only one of author or book flag should be set")
	}

	switch {
	case *authorFlag != "":
		resp, err := c.GetBooksByAuthor(ctx, &books.GetBooksRequest{Name: *authorFlag})
		checkRespError(err)
		fmt.Printf("Books: %s\n", resp.GetBooks())

	case *bookFlag != "":
		resp, err := c.GetAuthorsByBook(ctx, &books.GetAuthorsRequest{Title: *bookFlag})
		checkRespError(err)
		fmt.Printf("Authors: %s\n", resp.GetAuthors())

	default:
		failOnError(errors.New("wrong request"), "author or book must be specified") // If neither author nor book flag is set, print a message
	}
}

// failOnError logs if there an error and exits the program
func failOnError(err error, msg string) {
	if err != nil {
		slog.Error(
			msg,
			"error", err,
		)
		os.Exit(1)
	}
}

// checkRespError checks if gRPC response error and logs it
func checkRespError(err error) {
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			log.Fatal("unexpected error")
		}

		log.Fatalf("received an error: %s, code: %s", st.Message(), st.Code())
	}
}
