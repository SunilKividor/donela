package main

import (
	"fmt"
	"os"

	"github.com/SunilKividor/donela/internal/di"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../../.env"); err == nil {
		fmt.Println("[INFO] Loaded .env from project root")
	}

	server, err := di.Initialize()
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	server.Serve(port)
}
