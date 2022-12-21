# peanut

Golang template for API server

## Setup

- Install packages:\
  `$ go install`
- Update `.env`
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
$ ginkgo generate %FILENAME%
```

- Run test:

```
$ ginkgo ./usecase
// or
$ go test ./usecase
// or run all test in project
$ ginkgo ./...
```

## API docs:

- Install package:
  `$ go install github.com/swaggo/swag/cmd/swag@latest`
- Generate docs folder:\
  Write annotations before main() function and run:\
  `$ swag init`
- Generate API docs:\
  Write annotations before gin controller and run:\
  `$ swag init`

## Rules

- `domain`: entities defination
- `controller`: binding data, validation
- `usecase`: write business logic
- `resository`: get data from storage (DB, firebase, bigquery,..)
- `config`: setup or read configuration variables
- `pkg`: internal package
