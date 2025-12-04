package main

import (
	"context"
	"fmt"

	"github.com/SunilKividor/donela/internal/di"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err == nil {
		fmt.Println("[INFO] Loaded .env from project root")
	}

	w, err := di.InitializeWorker()
	if err != nil {
		panic(err)
	}

	w.Start(context.Background())
}
