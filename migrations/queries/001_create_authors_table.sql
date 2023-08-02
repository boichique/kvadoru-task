-- +goose Up
CREATE TABLE authors (
    id INT AUTO_INCREMENT,
    name VARCHAR(255),
    PRIMARY KEY(id)
);

CREATE INDEX idx_authors_name ON authors(name);

-- +goose Down
DROP TABLE authors;