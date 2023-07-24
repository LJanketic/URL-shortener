package server

import (
	"Minifyr/model"
	"Minifyr/utils"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func getAllMinifyrs(c *fiber.Ctx) error {

	minifyrs, err := model.GetAllMinfyrs()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": utils.ERR_GET_ALL + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(minifyrs)
}

func getMinifyr(c *fiber.Ctx) error {

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": utils.ERR_PARSE_JSON + err.Error(),
		})
	}

	minifyr, err := model.GetMinifyr(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": utils.ERR_GET_ONE + err.Error(),
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
			"message": utils.ERR_PARSE_JSON + err.Error(),
		})
	}

	if minifyr.Random {
		minifyr.Minifyr = utils.RandomizeURL(8/*int(time.Now().Unix())*/)
	}

	err = model.CreateMinifyr(minifyr)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": utils.ERR_CREATE_ONE + err.Error(),
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
			"message": utils.ERR_PARSE_JSON + err.Error(),
		})
	}

	err = model.UpdateMinifyr(minifyr)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": utils.ERR_UPDATE_ONE + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(minifyr)
}

func deleteMinifyr(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": utils.ERR_PARSE_JSON + err.Error(),
		})
	}

	err = model.DeleteMinifyr(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": utils.ERR_DELETE_ONE + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map {
		"message": utils.SCS_DELETE_ONE,
	})
}

func redirectViaMinifyr(c *fiber.Ctx) error {
	minifyrUrl := c.Params("redirect")

	minifyr, err := model.FindByMinifyrURL(minifyrUrl)
	if err != nil {
		fmt.Printf(utils.ERR_GET_ONE)
	}

	minifyr.Clicked += 1
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": utils.CLKD_ERR + err.Error(),
		})
	}

	return c.Redirect(minifyr.Redirect, fiber.StatusTemporaryRedirect)
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
	router.Delete("/minifyr/:id", deleteMinifyr)

	router.Get("/r/:redirect", redirectViaMinifyr)

	router.Listen(":3000")
}
