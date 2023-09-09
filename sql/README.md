# sql

Create database in container

```sh
docker exec -ti postgres-db createdb -U postgres gosql
```

Inspect db

```sh
docker exec -ti postgres-db psql -U postgres
\c gosql
\dt
SELECT * FROM product;
```

Run app

```sh
go run main.go
```
