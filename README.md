# Golang API Project

This project is an API built with Golang. The API is designed to manage orders, users, and kitchens with functionalities like user authentication, order management, and more.x

## Table of Contents

- [Technologies](#technologies)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Authentication](#authentication)
- [Environment Variables](#environment-variables)
- [Contributing](#contributing)
- [License](#license)

## Technologies

This project uses the following technologies:

- [Golang](https://golang.org/) for backend development
- [Echo](https://echo.labstack.com/) as the web framework
- [MySQL](https://www.mysql.com/) as the relational database
- JWT (JSON Web Token) for user authentication and authorization
- [bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt) for password hashing

## Project Structure

```bash
├── application
│   ├── dtos
│   ├── services
├── domain
│   ├── models
│   ├── repositories
├── infrastructure
│   ├── controllers
│   ├── db
│   ├── middlewares
│   └── security
├── cmd
│   └── main.go
├── go.mod
├── go.sum
└── README.md
```

- **application/**: Contains application logic, DTOs, and services.
- **domain/**: Defines domain models and repository interfaces.
- **infrastructure/**: Contains controllers, database connections, and middleware.
- **cmd/**: Entry point for the application.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/murilosopi/ramengo-api.git
   ```

2. Navigate into the project directory:

   ```bash
   cd ramengo-api
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

4. Set up a MySQL database by `db/schema.sql` and `db/seed.sql` and configure the `.env` file (see [Environment Variables](#environment-variables)).

5. Run the project:
   ```bash
   go run cmd/main.go
   ```

## Usage


### API Endpoints

#### 1. **Authentication**

- **Endpoint:** `/auth`
- **Method:** POST
- **Description:** Authenticates a user and optionally assigns a kitchen.
- **Request Body:**
  ```json
  {
    "email": "string",
    "password": "string",
    "kitchen": optional integer
  }
  ```
- **Example Request:**
  ```bash
  curl --location 'http://localhost:8080/auth' \
  --header 'Content-Type: application/json' \
  --data-raw '{
      "email": "john.doe@example.com",
      "password": "password123"
  }'
  ```

  ```bash
  curl --location 'http://localhost:8080/auth' \
  --header 'Content-Type: application/json' \
  --data-raw '{
      "email": "john.doe@example.com",
      "password": "password123",
      "kitchen": 2
  }'
  ```

#### 2. **Create User**

- **Endpoint:** `/user`
- **Method:** POST
- **Description:** Creates a new user. The address can be provided either as a nested object or by specifying an address ID.
- **Request Body:**
  ```json
  {
    "name": "string",
    "password": "string",
    "email": "string",
    "address": {
      "street": "string",
      "zipCode": "string",
      "number": integer
    },
    "addressID": optional integer
  }
  ```
- **Example Request:**
  ```bash
  curl --location 'http://localhost:8080/user' \
  --data-raw '{
      "name": "Alice Smith",
      "password": "alicepass",
      "email": "alice.smith@example.com",
      "address": {
          "street": "123 Elm Street",
          "zipCode": "12345",
          "number": 678
      }
  }'
  ```

  ```bash
  curl --location 'http://localhost:8080/user' \
  --data-raw '{
      "name": "Bob Johnson",
      "password": "bobpass",
      "email": "bob.johnson@example.com",
      "addressID": 5
  }'
  ```

#### 3. **Get User Order History**

- **Endpoint:** `/user/orders`
- **Method:** GET
- **Description:** Retrieves the order history for the authenticated user.
- **Authentication:** Bearer Token (User)
- **Request Body:** None
- **Example Request:**
  ```bash
  curl --location 'http://localhost:8080/user/orders' \
  --header 'Authorization: Bearer <User Token>'
  ```

#### 4. **Create Order**

- **Endpoint:** `/order`
- **Method:** POST
- **Description:** Creates a new order for the authenticated user.
- **Authentication:** Bearer Token (User)
- **Request Body:**
  ```json
  {
    "kitchenId": integer,
    "status": integer,
    "userId": integer
  }
  ```
- **Example Request:**
  ```bash
  curl --location --request POST 'http://localhost:8080/order' \
  --header 'Authorization: Bearer <User Token>' \
  --data ''
  ```

#### 5. **Change Order Status**

- **Endpoint:** `/order/status`
- **Method:** PATCH
- **Description:** Updates the status of an order.
- **Authentication:** Bearer Token (Kitchen)
- **Request Body:**
  ```json
  {
    "id": integer,
    "status": integer
  }
  ```
- **Example Request:**
  ```bash
  curl --location --request PATCH 'http://localhost:8080/order/status' \
  --header 'Content-Type: application/json' \
  --header 'Authorization: Bearer <Kitchen Token>' \
  --data '{
      "id": 45,
      "status": 3
  }'
  ```

#### 6. **Add User to Kitchen**

- **Endpoint:** `/kitchen/user`
- **Method:** POST
- **Description:** Adds a user to a kitchen.
- **Authentication:** Bearer Token (Kitchen)
- **Request Body:**
  ```json
  {
    "userID": integer
  }
  ```
- **Example Request:**
  ```bash
  curl --location 'http://localhost:8080/kitchen/user' \
  --header 'Content-Type: application/json' \
  --data '{
      "userID": 12
  }' \
  --header 'Authorization: Bearer <Kitchen Token>'
  ```

#### 7. **Get Current Orders for Kitchen**

- **Endpoint:** `/kitchen/orders`
- **Method:** GET
- **Description:** Retrieves current orders for the kitchen.
- **Authentication:** Bearer Token (Kitchen)
- **Request Body:** None
- **Example Request:**
  ```bash
  curl --location 'http://localhost:8080/kitchen/orders' \
  --header 'Authorization: Bearer <Kitchen Token>'
  ```

#### 8. **Cancel Not Ready Orders**

- **Endpoint:** `/kitchen/orders`
- **Method:** DELETE
- **Description:** Cancels orders that are not ready.
- **Authentication:** Bearer Token (Kitchen)
- **Request Body:** None
- **Example Request:**
  ```bash
  curl --location --request DELETE 'http://localhost:8080/kitchen/orders' \
  --header 'Authorization: Bearer <Kitchen Token>'
  ```

## Authentication

This API uses JWT (JSON Web Tokens) for authentication. Users must pass the token in the `Authorization` header for protected routes.

Example:

```bash
Authorization: Bearer <your-jwt-token>
```

To generate a token, use the `/login` endpoint with valid user credentials.

## Environment Variables

The project uses environment variables to manage configurations. Create a `.env` file in the root directory with the following variables:

```bash
JWT_SECRET_KEY=your_secret_key
DB_HOST=localhost
DB_USER=root
DB_PASSWORD=yourpassword
DB_NAME=yourdbname
```

These variables can be loaded using the [godotenv](https://github.com/joho/godotenv) package.