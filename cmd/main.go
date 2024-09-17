package main

import (
	"ramengo/infrastructure/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // load .env to os.Getenv() access
	http.Init()
}
