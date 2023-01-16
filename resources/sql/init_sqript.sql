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

CREATE TABLE game
(
    `id`       CHAR(36) PRIMARY KEY,
    `unit_on_field_1`  CHAR(26),
    `field_2`  CHAR(26),
    `field_3`  CHAR(26),
    `field_4`  CHAR(26),
    `field_5`  CHAR(26),
    `field_6`  CHAR(26),
    `field_7`  CHAR(26),
    `field_8`  CHAR(26),
    `field_9`  CHAR(26),
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
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F6` FOREIGN KEY (`field_6`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F7` FOREIGN KEY (`field_7`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F8` FOREIGN KEY (`field_8`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F9` FOREIGN KEY (`field_9`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F10` FOREIGN KEY (`field_10`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F11` FOREIGN KEY (`field_11`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F12` FOREIGN KEY (`field_12`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F13` FOREIGN KEY (`field_13`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F14` FOREIGN KEY (`field_14`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F15` FOREIGN KEY (`field_15`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F16` FOREIGN KEY (`field_16`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F17` FOREIGN KEY (`field_17`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F18` FOREIGN KEY (`field_18`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F19` FOREIGN KEY (`field_19`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F20` FOREIGN KEY (`field_20`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F21` FOREIGN KEY (`field_21`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F22` FOREIGN KEY (`field_22`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F23` FOREIGN KEY (`field_23`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F24` FOREIGN KEY (`field_24`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F25` FOREIGN KEY (`field_25`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F26` FOREIGN KEY (`field_26`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F27` FOREIGN KEY (`field_27`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F28` FOREIGN KEY (`field_28`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F29` FOREIGN KEY (`field_29`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F30` FOREIGN KEY (`field_30`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F31` FOREIGN KEY (`field_31`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F32` FOREIGN KEY (`field_32`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F33` FOREIGN KEY (`field_33`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F34` FOREIGN KEY (`field_34`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F35` FOREIGN KEY (`field_35`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F36` FOREIGN KEY (`field_36`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F37` FOREIGN KEY (`field_37`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F38` FOREIGN KEY (`field_38`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F39` FOREIGN KEY (`field_39`)
        REFERENCES `unit` (`id`),
    CONSTRAINT `FK_unit_id_F40` FOREIGN KEY (`field_40`)
        REFERENCES `unit` (`id`)
);


DROP TABLE game, card, unit
