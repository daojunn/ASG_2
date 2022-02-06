-- DROP DATABASE IF EXISTS ratingdb;

CREATE DATABASE ratingdb;

USE ratingdb;

-- DROP TABLE IF EXISTS Rating;
CREATE TABLE Rating(
	`RatingID` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
	`CreatorID` int NULL,
	`CreatorType` varchar(255),
	`TargetID` int NOT NULL,
    `TargetType` varchar(255) NOT NULL,
	`RatingScore` int NOT NULL,
	`Anonymous` boolean NOT NULL DEFAULT False,
	`DateTimePublished` datetime DEFAULT CURRENT_TIMESTAMP());