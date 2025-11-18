package main

import (
	"context"
	"fmt"

	"github.com/SunilKividor/donela/internal/di"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../../.env"); err == nil {
		fmt.Println("[INFO] Loaded .env from project root")
	}

	worker, err := di.InitializeWorker()
	if err != nil {
		panic(err)
	}

	worker.Start(context.Background())

}
