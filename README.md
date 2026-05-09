# go-fiber-product-api

REST API for product management built with Go Fiber and PostgreSQL.

## Tech Stack

- Go
- Fiber v2
- GORM
- PostgreSQL
- JWT

## Features

- User registration and login with JWT authentication
- HTTP-Only Cookie session management
- Full product CRUD scoped to authenticated user

## API Endpoints

### Authentication

```
POST /register     Create a new account
POST /login        Login and receive token
POST /logout       Logout and clear token
```

### Products

```
GET    /products      Get all products
GET    /products/:id  Get product by ID
POST   /products      Create a new product
PUT    /products/:id  Update a product
DELETE /products/:id  Delete a product
```

> All product endpoints require authentication.

## Getting Started

```bash
git clone https://github.com/LazyDevCode-888/go-fiber-product-api.git
cd go-fiber-product-api
go mod tidy
```

Copy the example environment file and fill in your values.

```bash
cp .env.example .env
```

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=your_database
DB_PORT=5432
JWT_SECRET=your_jwt_secret
```

```bash
go run main.go
```

Server will start at `http://localhost:3000`

## License

MIT