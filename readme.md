
## 1. Steps to run (type in console):

```
psql -U postgres

CREATE DATABASE viberate;

\q

migrate -database postgres://USERNAME:PASSWROD@localhost:5432/viberate?sslmode=disable -path ./migrations up

psql -U USERNAME -d viberate -f migrations/seed.sql

go mod tidy

go run main.go
```