# Go Gin Simple Blog with Full-Text Search

## Introduction

A blog API built with Go using Gin framework, featuring full-text search capabilities and JWT authentication.

## Features

- JWT Authentication
- Full-Text Search
- CRUD Operations for Articles
- CSV Import Support with go corroutine
- Swagger Documentation
- PostgreSQL Integration

## Project Structure

## API Endpoints

### Auth Routes

- POST /api/v1/register - Register new user
- POST /api/v1/login - User login
- POST /api/v1/refresh-token - Refresh JWT token
- POST /api/v1/change-password - Change password (Protected)
- POST /api/v1/check-username - Check username availability (Protected)

### Article Routes

- GET /api/v1/articles - Get all articles
- GET /api/v1/articles/:id - Get article by ID
- GET /api/v1/users/:id/articles - Get articles by user ID
- GET /api/v1/articles/search - Search articles
- POST /api/v1/articles - Create article (Protected)
- POST /api/v1/articles/csv - Import articles via CSV (Protected)
- PUT /api/v1/articles/:id - Update article (Protected)
- DELETE /api/v1/articles/:id - Delete article (Protected)

### Swagger Documentation

- GET /swagger/*any - Swagger UI endpoint

## Setup Instructions

1. Clone the repository
2. Create .env file with required configurations
3. Install dependencies:

   ```sh
   go mod tidy
   ```

4. Run the application:

   ```sh
   go run cmd/main.go
   ```

## Swagger Spesification

Access the Swagger UI at <http://localhost:5555/swagger/index.html>
Swagger Security Definitions

- API Key Auth (Bearer Token)
