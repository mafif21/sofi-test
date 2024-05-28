package main

import (
	"github.com/joho/godotenv"
	"penjadwalan-sidang-new/internal/app"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	app.StartApp()
}
