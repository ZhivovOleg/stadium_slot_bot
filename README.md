# Stadium Slot Bot

Бот для записи на тренировки, получения информации о стадионе, событиях и тд

## Used packages

### backend:

Golang
- [go-gin webserver](https://github.com/gin-gonic/gin)
- [gin-swagger](https://github.com/swaggo/gin-swagger)
- [pgx - PostgreSQL Driver](https://github.com/jackc/pgx)
- [pgxpool](https://pkg.go.dev/github.com/jackc/pgx/v4/pgxpool)
- [zap logger](https://github.com/uber-go/zap)
- [go-telegram bot](https://github.com/go-telegram/bot)

## Develop

Recommended IDE - VSCode.
<br />
Environment for project in `./.vscode/launch.json`.
<br />
For another IDE's don't forget set up env variables:
```bash
"StadiumSlotBotPort": "6000",
"StadiumSlotBotDbConnectionString": "postgres://pg:1@localhost:5432/stadiumSlotBot_db",
"StadiumSlotBotEnv": "dev",
"StadiumSlotBotAdminId": "[your telegram ID]"
```

## Debug & Test

### Debug backend

1. Install golang `swag` utility:
    ```bash
    go install github.com/swaggo/swag/cmd/swag@latest
    ```
0. Install dependencies:
    ```bash
    go get .
    ```
0. Run test database docker image:
    ```bash
    docker compose up -d ./.test/database/docker-compose.yml
    ```

0. After changes regenerate swagger files:
    ```bash
    swag init -g ./cmd/StadiumSlotBot/main.go -o ./docs
    ```
0. Set breakpoints and press `F5`