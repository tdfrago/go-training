CREATE TABLE `users` (
  `Id` INT NOT NULL AUTO_INCREMENT,
  `LastName` VARCHAR(45) NOT NULL,
  `FirstName` VARCHAR(45) NOT NULL,
  `UserName` VARCHAR(45) NOT NULL,
  `Password` VARCHAR(100) NOT NULL,
  PRIMARY KEY (`Id`));

CREATE TABLE `movies` (
  `Id` INT NOT NULL AUTO_INCREMENT,
  `Title` VARCHAR(45) NOT NULL,
  `Genre` VARCHAR(45) NOT NULL,
  `Year` INT NOT NULL,
  `Director` VARCHAR(45) NOT NULL,
  `Language` VARCHAR(45) NOT NULL,
  `Country` VARCHAR(45) NOT NULL,
  `Status` VARCHAR(100) NOT NULL,
  `UserName` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`Id`));
