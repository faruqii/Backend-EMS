package app

import (
	"os"

	"github.com/Magetan-Boyz/Backend/internal/config/database"
	"github.com/Magetan-Boyz/Backend/internal/config/seeder"
	"github.com/gofiber/fiber/v2"
)

func Start() {

	app := fiber.New()
	db, err := database.Connect()

	if err != nil {
		panic(err)
	}

	seed := seeder.Seed{DB: db}
	seed.SuperAdminSeeder()
	seed.AdminSeeder()
	seed.StudentSeeder()

	err = app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		panic(err)
	}

}
