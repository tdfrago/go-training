CREATE SCHEMA `testdb` ;

CREATE TABLE `testdb`.`students` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `firstname` VARCHAR(45) NULL,
  `lastname` VARCHAR(45) NULL,
  PRIMARY KEY (`id`));

INSERT INTO `testdb`.`students` (`id`, `firstname`, `lastname`) VALUES ('1', 'Jane', 'Doe');

UPDATE `testdb`.`students` SET `lastname` = 'Smith' WHERE (`id` = '2');

DELETE FROM `testdb`.`students` WHERE (`id` = '2');

DROP TABLE `testdb`.`students`;

CREATE TABLE `testdb`.`products` (
  `idproducts` INT NOT NULL,
  `name` VARCHAR(45) NOT NULL,
  `price` DECIMAL(10,2) NOT NULL,
  `description` VARCHAR(100) NOT NULL,
  PRIMARY KEY (`idproducts`),
  UNIQUE INDEX `name_UNIQUE` (`name` ASC) VISIBLE);

INSERT INTO `testdb`.`products` (`idproducts`, `name`, `price`, `description`) VALUES ('1', 'phone', '999.99', 'designed to shatter on impact if dropped');
