CREATE TABLE `game`
(
    `ID` CHAR(26) PRIMARY KEY,
    `GAME`       VARCHAR(50)
);

CREATE TABLE `card`
(
    `ID` CHAR(26) PRIMARY KEY,
    `Name` VARCHAR(26),
    `Damage` INT,
    `Speed` INT
);

CREATE TABLE `unit`
(
    `ID` CHAR(26) PRIMARY KEY,
    `CARD_ID` CHAR(26),
    CONSTRAINT `FK_card_id` FOREIGN KEY (`CARD_ID`)
        REFERENCES `card` (`ID`)
);