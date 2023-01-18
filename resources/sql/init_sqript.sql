DROP TABLE game, card, unit, fields;

INSERT INTO unit (id, card_id)
VALUES ('12345678123411781154567810', null);

SELECT *
FROM game;

CREATE TABLE game
(
    `id`   CHAR(36) PRIMARY KEY
);

CREATE TABLE card
(
    `id`     CHAR(36) PRIMARY KEY,
    `name`   VARCHAR(26),
    `damage` INT,
    `speed`  INT
);

CREATE TABLE unit
(
    `id`      CHAR(36) PRIMARY KEY,
    `card_id` CHAR(36),


    `health`  INT NOT NULL,
    CONSTRAINT `FK_card_id` FOREIGN KEY (`card_id`)
        REFERENCES `card` (`id`)
);

CREATE TABLE fields
(
    `number`  INT      NOT NULL,
    `unit_id` CHAR(36),
    `game_id` CHAR(36) NOT NULL,
    CHECK ( `number` <= 40),
    CONSTRAINT PRIMARY KEY (`number`, `game_id`),
    CONSTRAINT `FK_game_id` FOREIGN KEY (`game_id`)
        REFERENCES `game` (`id`),
    CONSTRAINT `FK_unit_id` FOREIGN KEY (`unit_id`)
        REFERENCES `unit` (`id`)
);
