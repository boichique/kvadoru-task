-- +goose Up
CREATE TABLE author_book (
    author_id INT,
    book_id INT,
    PRIMARY KEY(author_id, book_id),
    FOREIGN KEY(author_id) REFERENCES authors(id),
    FOREIGN KEY(book_id) REFERENCES books(id)
);

CREATE INDEX idx_author_book_author_id ON author_book(author_id);
CREATE INDEX idx_author_book_book_id ON author_book(book_id);

-- +goose Down
DROP TABLE author_book;
