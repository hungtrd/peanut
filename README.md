# peanut

A boilerplate for Golang backend

[github.com/hungtrd](https://github.com/hungtrd)
## Setup

- Install packages:\
  `$ go install`
- Copy `.env.example` to `.env` and update your enviroment
- Install `air` for live-reload source code.\
  `$ go install github.com/cosmtrek/air@latest`
- Run project:\
  `$ air`

## Unit test

- Setup gomock:\
  `$ go install github.com/golang/mock/mockgen@v1.6.0`
- Generate mock repository:\
  `$ mockgen -source repository/user.go -package mock -destination repository/mock/user.go`

- Setup ginkgo:

```
$ go install github.com/onsi/ginkgo/v2/ginkgo
$ go get github.com/onsi/gomega/...
```

- Create test suite:

```
$ cd to/pkg/is/tested
$ ginkgo bootstrap
$ ginkgo generate <file_name>
```

- Run test:

```
$ ginkgo ./usecase
// or
$ go test ./usecase
// or run all test in project
$ ginkgo ./...
```

## DB Migration

- Install cmd package:
  `$ go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
- Create migration files:
  `$ migrate create -ext sql -dir db/migration -seq <file_name>`
- Run migrate:
  - UP:
    `$ migrate -path db/migration -database "postgres://<user>:<pwd>@localhost:5432/<db_name>?sslmode=disable" -verbose up`
  - DOWN:
    `$ migrate -path db/migration -database "postgres://<user>:<pwd>@localhost:5432/<db_name>?sslmode=disable" -verbose down`
  - Or use makefile
    ```
    $ make migrateup
    $ make migratedown
    ```

## API docs:

> URL: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

- Install package:
  `$ go install github.com/swaggo/swag/cmd/swag@latest`
- Generate docs folder:  
  Write annotations before main() function and run:\
  `$ swag init`
- Generate API docs:  
  Write annotations before gin controller and run:
  ```
  $ swag init
  // If error, try
  $ swag init --parseDependency --parseInternal
  ```
- Format annotations code:  
  `$ swag fmt`

## Rules

- `domain`: entities defination
- `controller`: binding data, validation
- `usecase`: write business logic
- `resository`: get data from storage (DB, firebase, bigquery,..)
- `config`: setup or read configuration variables
- `pkg`: internal package
