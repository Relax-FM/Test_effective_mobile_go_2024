# Моя реализация тестового задания в effective mobile в 2024.
## Для начала

### Инициализация базы данных
Для начала вам потребуется установить PostgreSQL.

Далее, создать в ней базу "*__effective_mobile__*".

Для миграции таблицы пропишите команду, предварительно установив утилиту migrate:
```
migrate -path ./schema -database postgres://<your_db_user>:<your_db_password>@<your_db_host>:<your_db_port>/effective_mobile?sslmode=disable up 
```

### Добавим .env файл
Для корректного запуска сервера и для его настройки, добавим файл "*__.env__*" в директорию проекта (файл должен лежать на одном уровне с файлами server.go и music_entity.go и другими).

Сам файл "*__.env__*" должен выглядеть так (вместо ***** подставьте свои значения):
```
LOG_LEVEL=info

APP_PORT=*****

DB_PASSWORD=*****
DB_USERNAME=*****
DB_PORT=*****
DB_HOST=*****
DB_NAME=effective_mobile
SSL_MODE=disable

READ_TIMEOUT=*****
WRITE_TIMEOUT=***** 
```

### Запуск сервера
Для запуска сервера пропишите команду:
```
go run .\cmd\main.go 
```
Она должна загрузить все необходимые модули для корректного запуска сервера.

Подождите пока все модули будут установлены.

После этого введите команду повторно уже для запуска сервера.

Наслаждайтесь!






# My project for test my skills in go for "effective mobile" in 2024.
## Get start

### init db
For initialize db you need to download PostgreSQL then create "*__effective_mobile__*" db

Then write this command:
```
migrate -path ./schema -database postgres://<your_db_user>:<your_db_password>@<your_db_host>:<your_db_port>/effective_mobile?sslmode=disable up 
```

### add .env file
For start server you need to add ".env" in your project in project dir on root lvl

Your "*__.env__*" must be like:
```
LOG_LEVEL=info

APP_PORT=*****

DB_PASSWORD=*****
DB_USERNAME=*****
DB_PORT=*****
DB_HOST=*****
DB_NAME=effective_mobile
SSL_MODE=disable

READ_TIMEOUT=*****
WRITE_TIMEOUT=***** 
```

### Start server
For start server you need to write command:
```
go run .\cmd\main.go 
```
Then upload all needed modules.

After that try this command again for start server