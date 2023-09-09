package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetDatabase(envname string) string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	database, databaseSet := os.LookupEnv(envname)

	if !databaseSet {
		log.Fatalf("%s not set\n", envname)
	}

	return database
}
