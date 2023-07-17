package server

import (
	"Minifyr/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func getAllRedirects(ctx *fiber.Ctx) error {

	minifyrs, err := model.GetAllMinfyrs()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "Error getting all minifyr links " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(minifyrs)
}

func SetupServerListener() {

	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/getAllMinifyrs", getAllRedirects)
	router.Listen(":3000")
}
