# RSS Aggregator

A robust RSS feed aggregator API built with Go, allowing users to subscribe to and manage RSS feeds, and receive updates in a centralized location.

## Features

- 🔐 User authentication with API keys
- 📰 RSS feed management (create, read, delete)
- 👥 Feed following system
- 📝 Post aggregation from multiple feeds
- 🔄 Automatic feed updates
- 🛡️ CORS support
- 📊 PostgreSQL database with type-safe queries

## Tech Stack

- **Language:** Go
- **Framework:** Chi (HTTP router)
- **Database:** PostgreSQL
- **ORM:** SQLC (type-safe database operations)
- **Authentication:** API Key based
- **Documentation:** Postman Collection

## Prerequisites

- Go 1.24 or higher
- PostgreSQL
- Docker (optional, for SQLC)

## Environment Variables

Create a `.env` file in the root directory with the following variables:

```env
PORT=8000
DB_URL=postgres://username:password@localhost:5432/rssagg?sslmode=disable
```

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/rssagg.git
cd rssagg
```

2. Install dependencies:

```bash
go mod download
```

3. Generate database code:

```bash
docker run --rm -v "${PWD}:/src" -w /src kjconroy/sqlc generate
```

4. Run the application:

```bash
go run .
```

## API Endpoints

### Users

- `POST /v1/users` - Create a new user

  ```json
  {
    "name": "John Doe"
  }
  ```

- `GET /v1/users` - Get user details (requires API key)

### Feeds

- `POST /v1/feeds` - Create a new feed (requires API key)

  ```json
  {
    "name": "Tech Blog",
    "url": "https://example.com/feed.xml"
  }
  ```

- `GET /v1/feeds` - Get all feeds for authenticated user (requires API key)

### Feed Follows

- `POST /v1/feed_follows` - Follow a feed (requires API key)

  ```json
  {
    "feed_id": "uuid-of-feed"
  }
  ```

- `GET /v1/feed_follows` - Get all followed feeds (requires API key)
- `DELETE /v1/feed_follows/{feedFollowID}` - Unfollow a feed (requires API key)

### Posts

- `GET /v1/posts` - Get posts from followed feeds (requires API key)

## Authentication

The API uses API key authentication. Include the API key in the request header:

```
Authorization: ApiKey your-api-key-here
```

## Database Schema

### Users

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    api_key TEXT NOT NULL UNIQUE
);
```

### Feeds

```sql
CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);
```

### Feed Follows

```sql
CREATE TABLE feed_follows (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    UNIQUE(user_id, feed_id)
);
```

### Posts

```sql
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    published_at TIMESTAMP NOT NULL,
    url TEXT NOT NULL,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);
```

## API Documentation

For detailed API documentation and testing, use the provided [Postman Collection](https://web.postman.co/workspace/My-Workspace~d1615e25-6998-49ac-8295-35457901b082/collection/36200474-4565d479-ecc8-4b31-a141-ce1c632f56d3?action=share&creator=36200474)

## Project Structure

```
.
├── main.go                 # Application entry point
├── handler_*.go            # HTTP handlers
├── models.go              # Data models
├── json.go                # JSON response helpers
├── internal/
│   └── database/          # Database operations
├── sql/
│   ├── schema/           # Database schema
│   └── queries/          # SQL queries for SQLC
└── vendor/               # Dependencies
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
