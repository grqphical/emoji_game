package main

import (
	"log"
	"net/http"
	"os"

	"github.com/grqphical/emoji-game/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	router := routes.MakeRouter()

	log.Println("Server is running...")
	err = http.ListenAndServe(os.Getenv("HOST_ADDRESS"), router)
	if err != nil {
		log.Fatal(err)
	}
}