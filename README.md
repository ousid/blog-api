# Blog API

A RESTful blog API built with Go, Chi router, GORM, and SQLite.

## Stack

- **[Chi](https://github.com/go-chi/chi)** — HTTP router
- **[GORM](https://gorm.io)** — ORM with SQLite driver
- **[go-playground/validator](https://github.com/go-playground/validator)** — request validation
- **[gosimple/slug](https://github.com/gosimple/slug)** — slug generation
- **[godotenv](https://github.com/joho/godotenv)** — `.env` loading

## Getting started

```bash
cp .env.example .env   # configure environment
make run               # start the server
```

## Environment variables

| Variable  | Default        | Description           |
|-----------|----------------|-----------------------|
| `APP_PORT` | `8080`        | Port the server listens on |
| `APP_ENV`  | —             | Set to `development` to seed the database |
| `DB_PATH`  | `./database.db` | Path to the SQLite file |

## API

Base path: `/api`

### Posts

| Method | Endpoint         | Description      |
|--------|------------------|------------------|
| GET    | `/posts`         | List all posts   |
| GET    | `/posts/{slug}`  | Get a post       |
| POST   | `/posts`         | Create a post    |
| PUT    | `/posts/{id}`    | Update a post    |
| DELETE | `/posts/{id}`    | Delete a post    |

### Health check

```
GET /health  →  {"status": "ok"}
```

## Development

```bash
make run    # run the server
make lint   # run golangci-lint
make fmt    # format code
make test   # run tests
```
