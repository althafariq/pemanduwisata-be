## Installation

To install Gin package, you need to install Go and set your Go workspace first.

You first need [Go](https://go.dev/) installed (**version 1.16+ is required**), then you can use the below Go command to install Gin.

```sh
go get -u github.com/gin-gonic/gin
```

```sh
go get github.com/go-playground/validator/v10
```

```sh
go get github.com/mattn/go-sqlite3 database/sql
```

```sh
go get golang.org/x/crypto/bcrypt
```

## Run Migration and Seeder

In root directory, run this command to run migration and seeder

```sh
go run db/main.go
```

## Run the app

Running in `http://localhost:8080`

```sh
go run main.go
```
