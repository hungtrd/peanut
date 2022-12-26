migrateup:
	migrate -path db/migration -database "postgres://postgres:12345678@localhost:5431/mydb?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgres://postgres:12345678@localhost:5431/mydb?sslmode=disable" -verbose down
