# Simple Bank

## Requirements

- [PostgreSQL](https://www.postgresql.org/download/)
- [Go](https://go.dev/doc/install)

### Libraries Used for Go
- Gin (https://github.com/gin-gonic/gin)
- JWT-Go (https://github.com/dgrijalva/jwt-go)
- Godotenv (https://github.com/joho/godotenv)
- Gorm (gorm.io/gorm)
- Postgre Driver for Gorm (gorm.io/driver/postgres)

## Installation
1. Install PostgreSQL on your machine.
2. Create a database by running the SQL script in the `db/all-table.sql` file.
3. Install the necessary Go packages by running `go get`.
4. Clone this repository and navigate to the `simple-bank` directory.

## Pre-Usage

Before using the API, make sure to set the environment variables in the `.env` file. The following variables are required:

- **DB_HOST** (PostgreSQL host)
- **DB_PORT** (PostgreSQL port)
- **DB_USER** (PostgreSQL username)
- **DB_PASSWORD** (PostgreSQL password)
- **DB_NAME** (PostgreSQL database name)
- **SSL_MODE** (PostgreSQL SSL mode)
- **SERVER_PORT** (Port on which the server will run)
- **API_SECRET** (Secret key for JWT authentication)
- **TOKEN_HOUR_LIFESPAN** (Lifetime of JWT tokens in hours)

## API Reference

### User Endpoint

#### User registration

```http
POST /v1/register
```

| Parameter | Type     | Description    |
| :-------- | :------- | :------------- |
| `username` | `string` | **Required**.  |
| `password` | `string` | **Required**.  |


#### User login

```http
POST /v1/login
```

| Parameter | Type     | Description    |
| :-------- | :------- | :------------- |
| `username` | `string` | **Required**.  |
| `password` | `string` | **Required**.  |

Return value:

`message` `string` that contains Bearer Token for Authorization

#### User Update

```http
PUT /v1/user/
```

| Parameter | Type     | Description    |
| :-------- | :------- | :------------- |
| `username` | `string` | **Required**.  |
| `password` | `string` | **Required**.  |

#### User Profile

```http
GET /v1/user/profile
```

#### User logout

```http
POST /v1/logout
```

### Deposit Endpoint

#### New deposit

```http
POST /v1/deposit
```

| Parameter | Type     | Description    |
| :-------- | :------- | :------------- |
| `deposit_amount` | `float` | **Required**. |
| `deposit_description` | `string` | **Required**. |

### Payment Endpoint

#### New payment

```http
POST /v1/payment
```

| Parameter | Type     | Description    |
| :-------- | :------- | :------------- |
| `receipt_username` | `string` | **Required**. |
| `payment_amount` | `float` | **Required**. |
| `payment_description` | `string` | **Required**. |

### Log Endpoint

#### Find all transaction history

```http
GET /v1/history/transaction
```

#### Find all login/logout history

```http
GET /v1/history/auth
```
