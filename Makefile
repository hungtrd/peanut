migrateup:
	migrate -path db/migration -database "postgres://postgres:123456@localhost:5432/golang?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgres://postgres:123456@localhost:5432/golang?sslmode=disable" -verbose down