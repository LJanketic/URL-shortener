package server

import (
	"Minifyr/model"
	"Minifyr/utils"
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

func createMinifyr(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var minifyr model.Minifyr
	err := c.BodyParser(&minifyr)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "error parsing JSON " + err.Error(),
		})
	}

	if minifyr.Random {
		minifyr.Minifyr = utils.RandomizeURL(8/*int(time.Now().Unix())*/)
	}

	err = model.CreateMinifyr(minifyr)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not create minifyr in the db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(minifyr)
}

func updateMinifyr(c *fiber.Ctx) error {
	c.Accepts("application.json")

	var minifyr model.Minifyr

	err := c.BodyParser(&minifyr)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "Parsing JSON failed " + err.Error(),
		})
	}

	err = model.UpdateMinifyr(minifyr)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not update minifyr in the db " + err.Error(),
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

	router.Get("/minifyr", getAllMinifyrs)
	router.Get("/minifyr/:id", getMinifyr)
	router.Post("/minifyr", createMinifyr)
	router.Patch("/minifyr", updateMinifyr)

	router.Listen(":3000")
}
