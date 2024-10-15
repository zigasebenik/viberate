
## 1. Steps to run (type in console):

```
psql -U postgres

CREATE DATABASE viberate;

\q

migrate -database postgres://postgres:root@localhost:5432/viberate?sslmode=disable -path ./migrations up

psql -U postgres -d viberate -f migrations/seed.sql

go run main.go
```