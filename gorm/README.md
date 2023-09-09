# go-sql

Start db

```sh
docker compose up -d
```

Create database in container

```sh
docker exec -ti postgres-db createdb -U postgres gogorm
```

Inspect db

```sh
docker exec -ti postgres-db psql -U postgres
\c gogorm
\dt
SELECT * FROM products;
```

Run app

```sh
go run main.go
```
