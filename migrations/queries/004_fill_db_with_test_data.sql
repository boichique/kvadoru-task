-- +goose Up
INSERT INTO authors(name) VALUES ('J.K. Rowling');
INSERT INTO authors(name) VALUES ('Stephen King');
INSERT INTO authors(name) VALUES ('George R.R. Martin');
INSERT INTO authors(name) VALUES ('J.R.R. Tolkien');
INSERT INTO authors(name) VALUES ('Harper Lee');
INSERT INTO authors(name) VALUES ('F. Scott Fitzgerald');
INSERT INTO authors(name) VALUES ('John Grisham');
INSERT INTO authors(name) VALUES ('Dan Brown');
INSERT INTO authors(name) VALUES ('Agatha Christie');
INSERT INTO authors(name) VALUES ('Isaac Asimov');
INSERT INTO authors(name) VALUES ('Aldous Huxley');
INSERT INTO authors(name) VALUES ('Ray Bradbury');


INSERT INTO books(title) VALUES ('Harry Potter and the Sorcerers Stone');
INSERT INTO books(title) VALUES ('Harry Potter and the Chamber of Secrets');
INSERT INTO books(title) VALUES ('A Game of Thrones');
INSERT INTO books(title) VALUES ('A Clash of Kings');
INSERT INTO books(title) VALUES ('Misery');
INSERT INTO books(title) VALUES ('The Shining');
INSERT INTO books(title) VALUES ('The Lord of the Rings');
INSERT INTO books(title) VALUES ('To Kill a Mockingbird');
INSERT INTO books(title) VALUES ('The Da Vinci Code');
INSERT INTO books(title) VALUES ('And Then There Were None');
INSERT INTO books(title) VALUES ('Foundation');
INSERT INTO books(title) VALUES ('Brave New World');


INSERT INTO author_book(author_id, book_id) VALUES (1, 1);
INSERT INTO author_book(author_id, book_id) VALUES (1, 2);
INSERT INTO author_book(author_id, book_id) VALUES (2, 5);
INSERT INTO author_book(author_id, book_id) VALUES (2, 6);
INSERT INTO author_book(author_id, book_id) VALUES (3, 3);
INSERT INTO author_book(author_id, book_id) VALUES (3, 4);
INSERT INTO author_book(author_id, book_id) VALUES (4, 7);
INSERT INTO author_book(author_id, book_id) VALUES (5, 8);
INSERT INTO author_book(author_id, book_id) VALUES (6, 12);
INSERT INTO author_book(author_id, book_id) VALUES (7, 2);
INSERT INTO author_book(author_id, book_id) VALUES (1, 7);
INSERT INTO author_book(author_id, book_id) VALUES (8, 9);
INSERT INTO author_book(author_id, book_id) VALUES (9, 10);
INSERT INTO author_book(author_id, book_id) VALUES (10, 11);
INSERT INTO author_book(author_id, book_id) VALUES (11, 12);
INSERT INTO author_book(author_id, book_id) VALUES (12, 1);

-- +goose Down
DELETE FROM author_book WHERE author_id IN (1,2,3,4,5,6,7,8,9,10,11,12);
DELETE FROM books WHERE title IN ('Harry Potter', 'Harry Potter and the Chamber of Secrets', 'A Game of Thrones', 'A Clash of Kings', 'Misery', 'The Shining', 'The Lord of the Rings', 'To Kill a Mockingbird', 'The Da Vinci Code', 'And Then There Were None');
DELETE FROM authors WHERE name IN ('J.K. Rowling', 'Stephen King', 'George R.R. Martin', 'J. R. R. Tolkien', 'Harper Lee', 'F. Scott Fitzgerald', 'John Grisham', 'Dan Brown', 'Agatha Christie', 'Isaac Asimov', 'Aldous Huxley', 'Ray Bradbury');
