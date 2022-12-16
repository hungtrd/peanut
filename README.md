# peanut
Golang template for API server

## Setup
- Install packages:\
`go install`
- Update `.env`
- Install `air` for live-reload source code.\
`go install github.com/cosmtrek/air@latest`
- Run project:\
`air`

## Unit test
- Setup gomock:\
`go install github.com/golang/mock/mockgen@v1.6.0`
- Setup ginkgo:\
```
go install github.com/onsi/ginkgo/v2/ginkgo
go get github.com/onsi/gomega/...
```
- Create test suite:\
```
cd to/pkg/is/tested
ginkgo bootstrap
ginkgo generate %FILENAME%
```
- Create mock repository:\
`mockgen -source repository/user.go -package mock -destination repository/mock/user.go`

## Rules
- Domain: entities defination
- Controller: binding data, validation
- Usecase: write business logic
- Repository: get data from storage (DB, firebase, bigquery,..)

