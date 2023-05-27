CREATE TABLE game
(
    `game_id` CHAR(36) PRIMARY KEY,
    `first_user` VARCHAR(33) NOT NULL,
    `second_user` VARCHAR(33)
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
    `unit_id` CHAR(36) PRIMARY KEY,
    `card_id` CHAR(36),
    `game_id` CHAR(36),
    `name`    VARCHAR(26),
    `damage`  INT,
    `speed`   INT,
    `health`  INT NOT NULL,
    CONSTRAINT `FK_card_id` FOREIGN KEY (`card_id`)
        REFERENCES `card` (`card_id`)
);

CREATE TABLE field
(
    `y`       INT      NOT NULL,
    `x`       INT      NOT NULL,
    `unit_id` CHAR(36),
    `game_id` CHAR(36) NOT NULL,
    CONSTRAINT PRIMARY KEY (`x`, `y`, `game_id`),
    CONSTRAINT `FK_game_id` FOREIGN KEY (`game_id`)
        REFERENCES `game` (`game_id`),
    CONSTRAINT `FK_unit_id` FOREIGN KEY (`unit_id`)
        REFERENCES `unit` (`unit_id`)
);

CREATE TABLE user
(
    `login` VARCHAR(25) PRIMARY KEY
);

CREATE TABLE deck
(
    `id`         CHAR(36),
    `user_login` CHAR(36),
    CONSTRAINT PRIMARY KEY `PK_deck_user` (`id`, `user_login`)
);

CREATE TABLE card_in_deck
(
    `deck_id` CHAR(36),
    `card_id` CHAR(36)
);

INSERT INTO deck
VALUES ('6e7a3251-3333-4e15-7777-0a3739b11111',
        'Boba');

INSERT INTO card_in_deck
VALUES ('6e7a3251-3333-4e15-7777-0a3739b11111',
        '6e7a3251-0404-4e15-7777-0a3739b11111');



