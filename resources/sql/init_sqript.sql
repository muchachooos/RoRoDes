CREATE TABLE game
(
    `id` CHAR(26) PRIMARY KEY,
    `field_1` CHAR(26),
    `field_2` CHAR(26),
    `field_3` CHAR(26),
    `field_4` CHAR(26),
    `field_5` CHAR(26),
    `field_6` CHAR(26),
    `field_7` CHAR(26),
    `field_8` CHAR(26),
    `field_9` CHAR(26),
    `field_10` CHAR(26),
    `field_11` CHAR(26),
    `field_12` CHAR(26),
    `field_13` CHAR(26),
    `field_14` CHAR(26),
    `field_15` CHAR(26),
    `field_16` CHAR(26),
    `field_17` CHAR(26),
    `field_18` CHAR(26),
    `field_19` CHAR(26),
    `field_20` CHAR(26),
    `field_21` CHAR(26),
    `field_22` CHAR(26),
    `field_23` CHAR(26),
    `field_24` CHAR(26),
    `field_25` CHAR(26),
    `field_26` CHAR(26),
    `field_27` CHAR(26),
    `field_28` CHAR(26),
    `field_29` CHAR(26),
    `field_30` CHAR(26),
    `field_31` CHAR(26),
    `field_32` CHAR(26),
    `field_33` CHAR(26),
    `field_34` CHAR(26),
    `field_35` CHAR(26),
    `field_36` CHAR(26),
    `field_37` CHAR(26),
    `field_38` CHAR(26),
    `field_39` CHAR(26),
    `field_40` CHAR(26),

    CONSTRAINT `FK_unit_id_F1` FOREIGN KEY (`field_1`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F2` FOREIGN KEY (`field_2`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F3` FOREIGN KEY (`field_3`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F4` FOREIGN KEY (`field_4`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F5` FOREIGN KEY (`field_5`)
        REFERENCES `unit` (`id`)
);

CREATE TABLE card
(
    `id`     CHAR(26) PRIMARY KEY,
    `name`   VARCHAR(26),
    `damage` INT,
    `speed`  INT
);

CREATE TABLE unit
(
    `id`      CHAR(26) PRIMARY KEY,
    `card_id` CHAR(26),
    CONSTRAINT `FK_card_id` FOREIGN KEY (`card_id`)
        REFERENCES `card` (`id`)
);
