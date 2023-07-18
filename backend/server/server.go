package server

import (
	"Minifyr/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func getAllMinifyrs(c *fiber.Ctx) error {

	minifyrs, err := model.GetAllMinfyrs()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "Error getting all minifyr links " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(minifyrs)
}

func getMinifyr(c *fiber.Ctx) error {

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "Error could not parse id: " + err.Error(),
		})
	}

	minifyr, err := model.GetMinifyr(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "Error could not retrieve minfyr from db: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(minifyr)
}

func SetupServerListener() {

	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/getMinifyr", getAllMinifyrs)
	router.Get("/getMinifyr/:id", getMinifyr)
	router.Listen(":3000")
}
