mysql -u root

CREATE DATABASE be04;
USE be04;
CREATE USER `user04`@`localhost` IDENTIFIED BY '1234567';
GRANT ALL PRIVILEGES ON be04.* TO `user04`@`localhost`;
FLUSH PRIVILEGES; 
exit

mysql -u user04 -p be04

CREATE TABLE students (
	`id` INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
	`nim` VARCHAR(12) NOT NULL,
	`name` VARCHAR(64) NOT NULL,
	`semester` SMALLINT NOT NULL,
	`created_at` DATETIME NOT NULL,
	`updated_at` DATETIME NOT NULL
);

CREATE TABLE antrian (
	`id` INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
	`nomor` int NOT NULL,
	`status` VARCHAR(64) NOT NULL,
	`pelanggan` VARCHAR(64) NOT NULL,
	`teller` VARCHAR(64) NOT NULL,
	`created_at` DATETIME NOT NULL,
	`updated_at` DATETIME NOT NULL,
	`expired_at` DATETIME NOT NULL
);
