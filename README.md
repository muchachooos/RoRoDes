# API Documentation

###  Создание игры 
*Метод*: POST   
*URL*: /init_game

Создание игры с пустыми полями

Получение Game_ID

###  Получение игры
*Метод*: GET    
*URL*: /get_game?game_id=some_game_id

Получение всех полей игры по её ID

###  Создание юнита
*Метод*: POST   
*URL*: /create_unit?card_id=some_card_id&game_id=some_game_id&y=some_coordinates&x=some_coordinates

Создание в игре юнита по карточке персонажа (card_id) на указанных координатах

###  Получение юнита
*Метод*: GET    
*URL*: /get_unit?unit_id=some_unit_id

Получение X, Y, Game_ID юнита по его Unit_ID

###  Передвижение юнита
*Метод*: POST   
*URL*: /move_unit?unit_id=some_unit_id&direction=some_direction

Передвижение юнита на одну клетку в указанном направлении (direction)

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
    `y`       INT NOT NULL,
    `x`       INT NOT NULL,
    `unit_id` CHAR(36),
    `game_id` CHAR(36) NOT NULL,
    CONSTRAINT PRIMARY KEY (`x`, `y`, `game_id`),
    CONSTRAINT `FK_game_id` FOREIGN KEY (`game_id`)
        REFERENCES `game` (`game_id`),
    CONSTRAINT `FK_unit_id` FOREIGN KEY (`unit_id`)
        REFERENCES `unit` (`unit_id`)
);
```