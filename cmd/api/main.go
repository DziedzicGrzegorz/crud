package main

import (
	"crud/internal/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	server.StartApi()

}
