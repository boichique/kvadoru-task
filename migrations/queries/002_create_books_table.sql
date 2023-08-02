-- +goose Up
CREATE TABLE books (
    id INT AUTO_INCREMENT,
    title VARCHAR(127),
    PRIMARY KEY(id)
);

CREATE INDEX idx_books_title ON books(title);

-- +goose Down
DROP TABLE books;


