package belajarHandler

import (
    "github.com/Ayala-Crea/toko-bersama/database"
    "github.com/Ayala-Crea/toko-bersama/internal/models"
    "github.com/gofiber/fiber/v2"
)

// GetUsers adalah handler untuk mendapatkan pengguna dari database
func GetUsers(c *fiber.Ctx) error {
    {
    var users []model.User

    // Ambil data pengguna dari database
    result := database.DB.Find(&users)

    if result.Error != nil {
        // Jika terjadi kesalahan, kembalikan respons dengan status Bad Request dan pesan kesalahan
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": result.Error.Error(),
        })
    }

    // Jika berhasil, kembalikan respons dengan status OK dan data pengguna
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Data User Berhasil Ditampilkan!",
        "data":    users,
    })
}
}

func CreateUser(c *fiber.Ctx) error{
    var user model.User
    if err := c.BodyParser(&user); err != nil{
        return err
    }

    result := database.DB.Create(&user)

    if result.Error != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": result.Error.Error(),
        })
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "User Berhasil Ditambahkan",
        "data":     user,
    })
}

// func Index(c *gin.Context){}
