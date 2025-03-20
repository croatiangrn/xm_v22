# XM_V22 Project

Welcome to the **XM_V22** project! This is a Go-based application that provides RESTful APIs for managing companies. Below, you'll find instructions on how to set up, run, and use the project.

---

## Table of Contents
1. [Prerequisites](#prerequisites)
2. [Setup](#setup)
3. [Environment Variables](#environment-variables)
4. [Running the Project](#running-the-project)
5. [Authentication](#authentication)
6. [API Endpoints](#api-endpoints)
7. [Examples](#examples)

---

## Prerequisites

Before you begin, ensure you have the following installed:
- **Go** (version 1.20 or higher)
- **Docker** (optional, for running the database)
- **Make** (for using the provided Makefile)

---

## Setup

1. **Clone the repository:**
   ```bash
   git clone https://github.com/croatiangrn/xm_v22.git
   cd xm_v22
   ```

2. **Create the `.env` file:**
   Copy the `.env.sample` file to `.env` and update the values as needed:
   ```bash
   cp .env.sample .env
   ```

   Example `.env` file:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=xm_v22
   JWT_SECRET=your_jwt_secret
   SERVER_PORT=8080
   ```

3. **Install dependencies:**
   ```bash
   go mod tidy
   ```

---

## Environment Variables

The following environment variables are required:

| Variable       | Description                          | Example Value       |
|----------------|--------------------------------------|---------------------|
| `DB_HOST`      | Database host address                | `localhost`         |
| `DB_PORT`      | Database port                        | `5432`              |
| `DB_USER`      | Database username                    | `your_db_user`      |
| `DB_PASSWORD`  | Database password                    | `your_db_password`  |
| `DB_NAME`      | Database name                        | `xm_v22`            |
| `JWT_SECRET`   | Secret key for JWT token generation  | `your_jwt_secret`   |
| `SERVER_PORT`  | Port on which the server will run    | `8080`              |

---

## Running the Project

1. **Build the project:**
   ```bash
   make build
   ```

2. **Start the server:**
   ```bash
   make start
   ```

   The server will start on the port specified in the `.env` file (default: `8080`).

---

## Authentication

To access protected endpoints, you need to authenticate and obtain a JWT token.

### Login to Obtain JWT Token

Send a `POST` request to `/v1/login` with a username and password.

**Example cURL Request:**
```bash
curl -X POST http://localhost:8080/v1/login \
-H "Content-Type: application/json" \
-d '{
    "username": "admin",
    "password": "password"
}'
```

**Response:**
```json
{
    "token": "your_jwt_token_here"
}
```

### Using the JWT Token
Include the JWT token in the `Authorization` header for all protected endpoints:
```
Authorization: Bearer <your_jwt_token>
```

---

## API Endpoints

### 1. **Create a Company**
- **Method:** `POST`
- **Endpoint:** `/v1/companies`
- **Payload:**
  ```json
  {
      "name": "My Company 24",
      "amount_of_employees": 1,
      "registered": false,
      "type": "NonProfit"
  }
  ```

### 2. **Update a Company**
- **Method:** `PATCH`
- **Endpoint:** `/v1/companies/{UUID}`
- **Payload:**
  ```json
  {
      "name": "Updated Company Name",
      "type": "Corporations"
  }
  ```

### 3. **Delete a Company (Soft Delete)**
- **Method:** `DELETE`
- **Endpoint:** `/v1/companies/{UUID}`

### 4. **Get a Company**
- **Method:** `GET`
- **Endpoint:** `/v1/companies/{UUID}`

---

## Examples

### Create a Company
```bash
curl -X POST http://localhost:8080/v1/companies \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_jwt_token>" \
-d '{
    "name": "My Company 24",
    "amount_of_employees": 1,
    "registered": false,
    "type": "NonProfit"
}'
```

### Update a Company
```bash
curl -X PATCH http://localhost:8080/v1/companies/00000000-0000-0000-0000-000000000000 \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_jwt_token>" \
-d '{
    "name": "Updated Company Name",
    "type": "Corporations"
}'
```

### Delete a Company
```bash
curl -X DELETE http://localhost:8080/v1/companies/00000000-0000-0000-0000-000000000000 \
-H "Authorization: Bearer <your_jwt_token>"
```

### Get a Company
```bash
curl -X GET http://localhost:8080/v1/companies/00000000-0000-0000-0000-000000000000 \
-H "Authorization: Bearer <your_jwt_token>"
```

## Support

For any issues or questions, please open an issue on the [GitHub repository](https://github.com/croatiangrn/xm_v22/issues).