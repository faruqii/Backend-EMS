package app

import (
	"os"

	"github.com/Magetan-Boyz/Backend/internal/config/database"
	"github.com/gofiber/fiber/v2"
)

func Start() {

	app := fiber.New()
	database.Connect()

	err := app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		panic(err)
	}

}
