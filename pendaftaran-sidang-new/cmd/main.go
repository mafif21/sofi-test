package main

import (
	_ "embed"
	"github.com/joho/godotenv"
	"pendaftaran-sidang-new/internal/app"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	app.StartApp()
}
