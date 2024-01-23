// env.go
package services

import (
	"github.com/joho/godotenv"
	"log"
	//"os"
)

func LoadEnv() error {
	err := godotenv.Load("services.env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}
	return nil
}
