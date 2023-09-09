package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-faker/faker/v4"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	connStr string
)

type Product struct {
	gorm.Model
	Name      string  `faker:"word"`
	Price     float64 `faker:"oneof: 4.95, 9.99, 1600"`
	Available bool
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	database, databaseSet := os.LookupEnv("DATABASE")

	if !databaseSet {
		log.Fatalln("DATABASE not set")
	}

	connStr = database
}

func main() {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Product{})

	product := Product{}
	err = faker.FakeData(&product)

	if err != nil {
		log.Fatalln(err)
	}

	db.Create(&product)

	singleProduct := getProduct(db, 1)

	db.Model(&singleProduct).Update("Price", 200)
	db.Model(&singleProduct).Updates(Product{Price: 300, Available: true})

	allProduct := getAllProduct(db)

	fmt.Println(allProduct)

	allProduct = getAllAvailableProduct(db)

	fmt.Println(allProduct)
	// db.Delete(&singleProduct, 1)
}

func getProduct(db *gorm.DB, pk int) Product {
	var product Product

	db.First(&product, 1)

	return product
}

func getAllProduct(db *gorm.DB) []Product {
	var products []Product

	result := db.Find(&products)

	amount := result.RowsAffected

	fmt.Printf("All: %d\n", amount)

	return products
}

func getAllAvailableProduct(db *gorm.DB) []Product {
	var products []Product

	result := db.Where("available = ?", true).Find(&products)

	amount := result.RowsAffected

	fmt.Printf("All available: %d\n", amount)

	return products
}
