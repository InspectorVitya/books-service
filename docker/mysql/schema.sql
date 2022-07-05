CREATE TABLE IF NOT EXISTS `author` (
    `author_id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL,
    `surname` varchar(100) NOT NULL,
    PRIMARY KEY (`author_id`),
    UNIQUE KEY `author_UN` (`author_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE IF NOT EXISTS `books` (
    `book_id` int NOT NULL AUTO_INCREMENT,
    `title` varchar(100) NOT NULL,
    `description` varchar(100) NOT NULL,
    PRIMARY KEY (`book_id`),
    UNIQUE KEY `books_UN` (`book_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `author_book` (
   `author_id` int NOT NULL,
   `book_id` int NOT NULL,
    KEY `author_book_FK` (`author_id`),
    KEY `author_book_FK_1` (`book_id`),
    CONSTRAINT `author_book_FK` FOREIGN KEY (`author_id`) REFERENCES `author` (`author_id`) ON DELETE CASCADE,
    CONSTRAINT `author_book_FK_1` FOREIGN KEY (`book_id`) REFERENCES `books` (`book_id`) ON DELETE CASCADE
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


INSERT INTO author (name, surname)
VALUES ('Lev', 'Tolstoy'),
       ('J', 'Rowling');

INSERT INTO books (title, description)
VALUES ('Anna karenina', 'some good book'),
       ('детство', 'еще 1 книга Толстого'),
       ('Harry Potter', 'classic book'),
       ('1986', 'у нее нет автора в нашей табличке(((');

INSERT INTO author_book (author_id, book_id)
VALUES (1, 1),
       (1, 2),
       (2, 3);
