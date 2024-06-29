package main

import (
	"fmt"

	"github.com/Magetan-Boyz/Backend/internal/app"
	// "github.com/joho/godotenv"
)

func main() {
	// err := godotenv.Load()

	// if err != nil {
	// 	fmt.Println("Error loading .env file")
	// 	return
	// }

	app.Start()
}
