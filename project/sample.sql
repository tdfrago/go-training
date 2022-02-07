CREATE TABLE `testdb`.`users` (
  `Id` INT NOT NULL AUTO_INCREMENT,
  `LastName` VARCHAR(45) NOT NULL,
  `FirstName` VARCHAR(45) NOT NULL,
  `UserName` VARCHAR(45) NOT NULL,
  `Password` VARCHAR(100) NOT NULL,
  PRIMARY KEY (`Id`));

CREATE TABLE `testdb`.`movies` (
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

  SELECT Id FROM testdb.movies WHERE UserName = ? AND Title = ? AND Genre = ? AND Year = ? AND Director = ? AND Language = ? AND Country= ?;
