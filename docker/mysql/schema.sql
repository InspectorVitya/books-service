CREATE TABLE IF NOT EXISTS `author` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL,
    `surname` varchar(100) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `author_UN` (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE IF NOT EXISTS `books` (
    `id` int NOT NULL AUTO_INCREMENT,
    `title` varchar(100) NOT NULL,
    `description` varchar(100) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `books_UN` (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `author_book` (
   `author_id` int NOT NULL,
   `book_id` int NOT NULL,
    KEY `author_book_FK` (`author_id`),
    KEY `author_book_FK_1` (`book_id`),
    CONSTRAINT `author_book_FK` FOREIGN KEY (`author_id`) REFERENCES `author` (`id`) ON DELETE CASCADE,
    CONSTRAINT `author_book_FK_1` FOREIGN KEY (`book_id`) REFERENCES `books` (`id`) ON DELETE CASCADE
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


INSERT INTO author (name, surname)
VALUES ('Lev', 'Tolstoy'),
       ('J', 'Rowling'),
       ('A', 'Belov'),
       ('V', 'Nikolashin');

INSERT INTO books (title, description)
VALUES ('Anna karenina', 'some good book'),
       ('Сhildhood', 'some good book2'),
       ('Harry Potter', 'classic book'),
       ('Economy', 'for students');

INSERT INTO author_book (author_id, book_id)
VALUES (1, 1),
       (1, 2),
       (2, 3),
       (3, 4),
       (4, 4);
