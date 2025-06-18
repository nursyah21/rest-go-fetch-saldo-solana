package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var (
	RPC_HELIUS string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system env vars")
	}

	RPC_HELIUS = os.Getenv("RPC_HELIUS")
	if RPC_HELIUS == "" {
		log.Fatal("RPC_HELIUS is not set in environment or .env file")
	}
}

func main() {
	app := fiber.New()

	app.Post("/api/get-balances", GetBalance)

	app.Listen(":5000")
}
