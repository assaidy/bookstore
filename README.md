# Bookstore API

A RESTful API for managing a bookstore. This API allows you to manage books, categories, customers, and orders in a simple, organized way. It supports CRUD operations and provides a structured way to handle bookstore data.

---

## Features

- **Books Management**: Create, read, update, and delete book records.
- **Category Categorization**: Organize books by category.
- **Customer Management**: Manage customer data.
- **Order Processing**: Track customer orders, including order status and history.
- **Search**: Search books by title, or category.

---

## Technologies Used

- **Backend**: Go Fiber
- **Database**: Postgres
- **Authentication**: JWT-based authentication

---

## Getting Started

### Prerequisites

- goose migration: https://github.com/pressly/goose
- Create a postgres database for the store

### Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/assaidy/bookstore.git
   cd bookstore
   ```

2. **Install Dependencies**

   ```bash
   go mod tidy
   ```

3. **Environment Variables**

   Create a `.env` file in the root directory and add the following variables:

   ```env
    PORT=8080

    DB_HOST=
    DB_PORT=
    DB_DATABASE=
    DB_USERNAME=
    DB_PASSWORD=
    DB_SCHEMA=

    JWT_SECRET=
   ```

4. **Migrate**
    ```bash
    make up
    ```

5. **Run the Server**

   ```bash
    make run
   ```

   The server should now be running at `http://localhost:8080`.

---

## API Endpoints

### User

- **POST** `/user/register` - register a user
- request:
    ```json
    {
      "name": "John Doe",
      "email": "johndoe@example.com",
      "username": "johndoe",
      "password": "SecurePassword123",
      "address": "123 Main St, Springfield"
    }
    ```
- response:
    ```json
    {
      "message": "created successfully",
      "data": {
        "token": "your_jwt_token_here",
        "user": {
          "id": 1,
          "name": "John Doe",
          "username": "johndoe",
          "email": "johndoe@example.com",
          "address": "123 Main St, Springfield",
          "joinedAt": "2023-01-01T00:00:00Z"
        }
      }
    }
    ```

### Books (TODO)


### Category (TODO)


### Orders (TODO)


---

## Authentication

This API uses **JWT (JSON Web Tokens)** for secure endpoints. To access protected routes, include the JWT token in the `Authorization` header:

```
Authorization: Bearer <your_token>
```

---

## Error Handling

Errors are returned in the following format:

```json
{
  "code": status_code,
  "message": "Error message here",
  "errors":  [] // optional: only for validation errors
}
```

Common HTTP response codes used:
- `200 OK` - Success
- `201 Created` - New resource created
- `400 Bad Request` - Validation error
- `401 Unauthorized` - Authentication required
- `404 Not Found` - Resource not found
- `409 Conflict` - Conflict, typically when a request could cause duplicate records or violates unique constraints
- `422 Unprocessable Entity` - The request is syntactically correct, but it cannot be processed due to semantic or logical errors
- `500 Internal Server Error` - Server error

