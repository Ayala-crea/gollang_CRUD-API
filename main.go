package main

import (
	// "fmt"

	"github.com/Ayala-Crea/toko-bersama/database"
	"github.com/Ayala-Crea/toko-bersama/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Membuat koneksi ke database
	// database.ConnectDB()

	app := fiber.New()

	database.ConnectDB()

	router.SetupRoutes(app)

	app.Listen(":3000")

}
