package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func APICelery(api fiber.Router) {

	api.Post("/parser/fake_parser", func (c *fiber.Ctx) error{
		
		fileName := "marafon-" + uuid.New().String() + ".xlsx"

		return c.JSON(fiber.Map{
			"filename": fileName,
		})
	})
}