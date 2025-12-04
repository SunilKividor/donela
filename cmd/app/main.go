package main

import (
	"fmt"

	"github.com/SunilKividor/donela/internal/di"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err == nil {
		fmt.Println("[INFO] Loaded .env from project root")
	}

	server, err := di.InitializeApp()
	if err != nil {
		panic(err)
	}

	err = server.Serve(server.Port)
	if err != nil {
		panic(err)
	}
}
