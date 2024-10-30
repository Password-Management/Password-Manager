package helpers

import (
	"log"

	"github.com/joho/godotenv"
)

func Getenv() error {
	err := godotenv.Load()
	if err != nil {
		log.Print("error loading .env file" + err.Error())
		return err
	}
	return nil
}
