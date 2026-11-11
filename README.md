# EventDiscovery

A demo REST API built with [Gin](https://github.com/gin-gonic/gin) in Go, showcasing CRUD operations and JWT-based authentication.

This project is for learning/demo purposes — it covers the basics of structuring a Gin application with a database layer, models, routes, middleware, and auth.

## Tech Stack

- **Go** + **Gin** — HTTP router/framework
- **SQLite** (`mattn/go-sqlite3`) — database
- **golang-jwt/jwt** — JWT authentication
- **bcrypt** — password hashing

## Features

- User signup/login with hashed passwords
- JWT issued on login, required for protected routes
- Full CRUD on events
- Event registration/cancellation for authenticated users
- Ownership checks (only the event creator can update/delete it)

## Project Structure

```
db/           # database connection + table creation
middlewares/  # auth middleware (JWT verification)
models/       # Event and User models + queries
routes/       # route handlers
utils/        # JWT and password hashing helpers
api-test/     # .http request samples (REST Client)
```

## Getting Started

### Prerequisites

- Go 1.26+

### Run

```bash
go run main.go
```

The server starts on `http://localhost:8080` and creates `api.db` (SQLite) on first run.

## API Endpoints

### Auth

| Method | Endpoint  | Description          |
|--------|-----------|-----------------------|
| POST   | `/signup` | Create a new user     |
| POST   | `/login`  | Log in, returns a JWT |

### Events

| Method | Endpoint              | Auth required | Description            |
|--------|------------------------|---------------|-------------------------|
| GET    | `/events`              | No            | List all events         |
| GET    | `/events/:id`          | No            | Get a single event      |
| POST   | `/events`              | Yes           | Create an event         |
| PUT    | `/events/:id`          | Yes (owner)   | Update an event         |
| DELETE | `/events/:id`          | Yes (owner)   | Delete an event         |
| POST   | `/events/:id/register` | Yes          | Register for an event   |
| DELETE | `/events/:id/cancel`   | Yes           | Cancel a registration   |

### Authentication

Protected routes require an `Authorization` header containing the JWT returned from `/login`:

```
Authorization: <token>
```

## Example Requests

Signup:
```http
POST /signup
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "secret"
}
```

Login:
```http
POST /login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "secret"
}
```

Create event (authenticated):
```http
POST /events
Content-Type: application/json
Authorization: <token>

{
  "name": "Test Event",
  "description": "A test event",
  "location": "Test Location",
  "dateTime": "2025-01-01T15:30:00.000Z"
}
```

More sample requests are available in [`api-test/`](./api-test).

## Disclaimer

This is a demo/learning project, not production-ready. The JWT secret is hardcoded and there's no input sanitization beyond basic required-field validation.
