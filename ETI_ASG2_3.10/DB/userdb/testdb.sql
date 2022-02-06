-- DROP DATABASE IF EXISTS ETIAssignment2TestDB;

CREATE DATABASE testuserdb;

USE testuserdb;


CREATE TABLE Student(
	`StudentID` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
	`Name` varchar(255) NOT NULL);
    
INSERT INTO Student (`StudentID`, `Name`) VALUES (1, 'Ethan');
INSERT INTO Student (`StudentID`, `Name`) VALUES (2, 'Kester');
INSERT INTO Student (`StudentID`, `Name`) VALUES (3, 'John');
INSERT INTO Student (`StudentID`, `Name`) VALUES (4, 'Ben');
    

CREATE TABLE Tutor(
	`TutorID` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
	`Name` varchar(255) NOT NULL);
INSERT INTO Tutor(`TutorID`, `Name`) VALUES (1, 'Kenneth');
INSERT INTO Tutor (`TutorID`, `Name`) VALUES (2, 'Desmond');
INSERT INTO Tutor(`TutorID`, `Name`) VALUES (3, 'Susan');
INSERT INTO Tutor(`TutorID`, `Name`) VALUES (4, 'Karen');

