# Go-CryptoManager

Go-CryptoManager is a RESTful API service written in Go, designed to manage cryptocurrency assets. It provides endpoints for creating, retrieving, updating, and deleting cryptocurrency information.

## Features

- List all cryptocurrencies
- Retrieve a single cryptocurrency by ID
- Create new cryptocurrency entries
- Update existing cryptocurrency entries
- Delete cryptocurrency entries

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What things you need to install the software and how to install them:

- Go (version 1.21.4 or higher)
- MySQL (or MariaDB)
- A REST client like Postman or cURL (for testing endpoints)

### Installing

A step-by-step series of examples that tell you how to get a development environment running:

1. Clone the repository to your local machine:
   ```sh
   git clone https://github.com/KaNiuSii/Go-CryptoManager.git
   ```

2. Navigate to the project directory:
   ```sh
   cd Go-CryptoManager
   ```

3. Build the project:
   ```sh
   go build
   ```

4. Run the service:
   ```sh
   go run main.go
   ```

5. The server will start at `localhost:8081`.

## Usage

Here are some example requests you can make to interact with the API:

### List all Cryptocurrencies

```http
GET /crypto/
```

### Get a Single Cryptocurrency

```http
GET /crypto/{cryptoId}
```

### Create a New Cryptocurrency

```http
POST /crypto/
```

### Update a Cryptocurrency

```http
PUT /crypto/{cryptoId}
```

### Delete a Cryptocurrency

```http
DELETE /crypto/{cryptoId}
```

## Built With

- [Go](https://golang.org/) - The Go programming language
- [Gorilla Mux](https://github.com/gorilla/mux) - The HTTP routing and URL matcher for Go
- [GORM](https://gorm.io/) - The ORM library for Go
- [MYSQL](https://www.mysql.com/) - For Database
- [POSTMAN](https://www.postman.com/) - API Testing