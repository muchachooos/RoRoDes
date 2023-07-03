# API Documentation

###  Создание игры 
*Метод*: POST   
*URL*: /init_game

Создание игры с пустыми полями и получение её ID    
Ожидаемый ответ:
```json
"d125e254-41d1-409b-b970-43afd1cf5582"
```

###  Получение игры
*Метод*: GET    
*URL*: /get_game?game_id=some_game_id

Получение всех полей игры по её ID  
Ожидаемый ответ:
```json
[
  {
    "y": 0,
    "x": 0,
    "unit_id": null
  },
  {
    "y": 1,
    "x": 0,
    "unit_id": null
  },
  .
  .
  .
  {
    "y": 4,
    "x": 7,
    "unit_id": null
  }
]
```

###  Создание юнита
*Метод*: POST   
*URL*: /create_unit?card_id=some_card_id&game_id=some_game_id&y=some_coordinates&x=some_coordinates

Создание в игре юнита по карточке персонажа (card_id) на указанных координатах  
Ожидаемый ответ:
```json
[
  {
    "y": 3,
    "x": 7,
    "unit_id": "0617c7f4-b79b-4381-a245-80a607d16562",
    "card_id": "6e7a3251-3333-4e15-7777-0a3739b1be0e",
    "game_id": "f7829347-370a-46db-b996-d0ed3cdc17d1",
    "name": "some_name",
    "damage": 33,
    "speed": 250,
    "health": 980
  }
]
```

###  Получение юнита
*Метод*: GET    
*URL*: /get_unit?unit_id=some_unit_id

Получение X, Y, Game_ID юнита по его Unit_ID    
Ожидаемый ответ:
```json
[
  {
    "y": 3,
    "x": 7,
    "unit_id": "0617c7f4-b79b-4381-a245-80a607d16562",
    "card_id": "6e7a3251-3333-4e15-7777-0a3739b1be0e",
    "game_id": "f7829347-370a-46db-b996-d0ed3cdc17d1",
    "name": "some_name",
    "damage": 33,
    "speed": 250,
    "health": 980
  }
]
```

###  Передвижение юнита
*Метод*: POST   
*URL*: /move_unit?unit_id=some_unit_id&direction=some_direction

Передвижение юнита на одну клетку в указанном направлении (direction)   
Ожидаемый ответ:
```json
true
```
# Service Configuration

Файл конфигурации называется configuration.json и должен находиться в одной директории с исполняемым файлом

```json
{
  "port": 8080,
  "DataBase": {
    "user": "admin",
    "password": "admin",
    "host": "127.0.0.1",
    "dataBaseName": "some_DB_name",
    "db_port": 3306
  }
}
```

# Database Schema

Требуемые таблицы

```sql
CREATE TABLE game
(
    `game_id`     CHAR(36) PRIMARY KEY,
    `first_user`  VARCHAR(33) NOT NULL,
    `second_user` VARCHAR(33)
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

CREATE TABLE card
(
    `card_id` CHAR(36) PRIMARY KEY,
    `name`    VARCHAR(26),
    `damage`  INT,
    `speed`   INT,
    `health`  INT NOT NULL,
    `picture` MEDIUMBLOB
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
    `card_id` CHAR(36),
    `count`   INT,
    CONSTRAINT PRIMARY KEY `PK_card_in_deck` (`deck_id`, `card_id`)
);
```