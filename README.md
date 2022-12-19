# peanut
Golang template for API server

## Setup
- Install packages:\
`go install`
- Update `.env`
- Install `air` for live-reload source code.\
```go install github.com/cosmtrek/air@latest```
- Run project:\
`air`
## Rules
- Domain: entities defination
- Controller: binding data, validation
- Usecase: write business logic
- Repository: get data from storage (DB, firebase, bigquery,..)
## Use
- hash password\
```go get golang.org/x/crypto/bcrypt```
- testcase\
```go install github.com/golang/mock/mockgen@v1.6.0```
- jwt\
```go get -u github.com/golang-jwt/jwt/v4```