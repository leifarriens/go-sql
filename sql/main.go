package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-faker/faker/v4"
	env "github.com/leifarriens/sql/internal"

	_ "github.com/lib/pq"
)

var (
	connStr string
)

type Product struct {
	Name      string  `faker:"word"`
	Price     float64 `faker:"oneof: 4.95, 9.99, 1600"`
	Available bool
}

func init() {
	database := env.GetDatabase("DATABASE")

	connStr = database
}

func main() {
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}

	createExampleTable(db)

	product := Product{}
	err = faker.FakeData(&product)

	if err != nil {
		log.Fatalln(err)
	}

	pk := insertProduct(db, product)

	singleProduct := getProduct(db, pk)

	fmt.Printf("Name: %s\n", singleProduct.Name)
	fmt.Printf("Name: %f\n", singleProduct.Price)
	fmt.Printf("Name: %t\n", singleProduct.Available)

	allProducts := getAllProduct(db)

	fmt.Println(allProducts)
}

func createExampleTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS product(
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		price NUMERIC(6,2) NOT NULL,
		available BOOLEAN,
		created timestamp DEFAULT NOW()
	)`

	_, err := db.Exec(query)

	if err != nil {
		log.Fatalln(err)
	}
}

func insertProduct(db *sql.DB, product Product) int {
	query := `INSERT INTO product (name, price, available) VALUES ($1, $2, $3) RETURNING id`

	var pk int

	err := db.QueryRow(query, product.Name, product.Price, product.Available).Scan(&pk)

	if err != nil {
		log.Fatalln(err)
	}

	return pk
}

func getProduct(db *sql.DB, pk int) Product {
	query := "SELECT name, price, available FROM product WHERE id = $1"

	var name string
	var price float64
	var available bool

	err := db.QueryRow(query, pk).Scan(&name, &price, &available)

	if err != nil {
		log.Fatalln(err)
	}

	return Product{name, price, available}
}

func getAllProduct(db *sql.DB) []Product {
	data := []Product{}

	query := "SELECT name, price, available FROM product"
	rows, err := db.Query(query)

	if err != nil {
		log.Fatalln(err)
	}

	defer rows.Close()

	var product Product

	for rows.Next() {
		err := rows.Scan(&product.Name, &product.Price, &product.Available)

		if err != nil {
			log.Fatalln(err)
		}

		data = append(data, product)
	}

	return data
}
