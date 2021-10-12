package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	a := os.Getenv("DBNAME")
	log.Println(a)
	//http.App()
}
