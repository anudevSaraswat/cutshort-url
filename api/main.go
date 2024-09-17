package main

import (
	"os"

	"github.com/anudevSaraswat/cutshort-url/routes"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	engine := routes.SetupRoutes()

	err = engine.Run(os.Getenv("APP_PORT"))
	if err != nil {
		panic(err)
	}

}
