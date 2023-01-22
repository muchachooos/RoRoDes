DROP TABLE game, card, unit, fields;

INSERT INTO card (`card_id`, `name`, `damage`, `speed`, `health`, `picture`)
VALUES ('80cbfcb0-1111-4e44-a484-a222d665f05b', 'eblanNumberOne', 33, 250, 980, null);

CREATE TABLE game
(
    `game_id` CHAR(36) PRIMARY KEY
);

CREATE TABLE card
(
    `card_id` CHAR(36) PRIMARY KEY,
    `name`    VARCHAR(26),
    `damage`  INT,
    `speed`   INT,
    `health`  INT NOT NULL,
    `picture` MEDIUMBLOB
);

CREATE TABLE unit
(
    `unit_id`   CHAR(36) PRIMARY KEY,
    `card_id`   CHAR(36),
    `game_id`   CHAR(36),
    `field_num` INT,
    `name`      VARCHAR(26),
    `damage`    INT,
    `speed`     INT,
    `health`    INT NOT NULL,
    CONSTRAINT `FK_card_id` FOREIGN KEY (`card_id`)
        REFERENCES `card` (`card_id`)
);

CREATE TABLE fields
(
    `number`  INT      NOT NULL,
    `unit_id` CHAR(36),
    `game_id` CHAR(36) NOT NULL,
    CHECK ( `number` <= 40),
    CONSTRAINT PRIMARY KEY (`number`, `game_id`),
    CONSTRAINT `FK_game_id` FOREIGN KEY (`game_id`)
        REFERENCES `game` (`game_id`),
    CONSTRAINT `FK_unit_id` FOREIGN KEY (`unit_id`)
        REFERENCES `unit` (`unit_id`)
);
