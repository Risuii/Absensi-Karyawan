CREATE TABLE `github.com/Risuii`.`employee` (
  `ID` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NULL,
  `password` VARCHAR(255) NULL,
  `email` VARCHAR(255) NULL,
  `checkin` DATETIME NULL DEFAULT(now()),
  `checkout` DATETIME NULL DEFAULT(now()),
  `created_at` DATETIME NULL DEFAULT (now()),
  `update_at` DATETIME NULL DEFAULT (now()),
  PRIMARY KEY (`ID`)
);

CREATE TABLE `github.com/Risuii`.`activity` (
    `ID` INT NOT NULL AUTO_INCREMENT,
    `userID` INT NOT NULL,
    `deskripsi` VARCHAR(255) NULL,
    `created_at` DATE NULL DEFAULT (now()),
    `update_at` DATE NULL DEFAULT (now()),
    PRIMARY KEY(`ID`),
    FOREIGN KEY (`userID`) REFERENCES employee(`ID`)
);

CREATE TABLE `github.com/Risuii`.`absen` (
  `ID` INT NOT NULL AUTO_INCREMENT,
  `userID` INT NOT NULL,
  `name` VARCHAR(255) NULL,
  `checkin` DATETIME NULL,
  `checkout` DATETIME NULL,
  PRIMARY KEY (`ID`),
  FOREIGN KEY (`userID`) REFERENCES employee(`ID`)
)
