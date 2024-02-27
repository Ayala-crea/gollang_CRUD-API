package belajarRoutes

import (
	belajarHandler "github.com/Ayala-Crea/toko-bersama/internal/handlers/belajar"

	"github.com/gofiber/fiber/v2"
)

func SetupBelajarRoutes(router fiber.Router) {
	user := router.Group("/belajar")
	user.Get("/", belajarHandler.GetUsers)
	user.Post("/tambahdata", belajarHandler.CreateUser)
}
