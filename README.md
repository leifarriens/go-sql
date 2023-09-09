# go-sql

Start db

```sh
docker compose up -d
```

Create database in container

```sh
docker exec -ti postgres-db createdb -U postgres gopostgres
```

Inspect db

```sh
docker exec -ti postgres-db psql -U postgres
\c gopostgres
\dt
SELECT * FROM product
```

Run app

```sh
go run main.go
```
