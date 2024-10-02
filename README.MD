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