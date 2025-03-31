package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zollidan/aafbet/db"
	"github.com/zollidan/aafbet/models"
)

// Группировка маршрутов
func APIDatabase(api fiber.Router) {
	dbGroup := api.Group("/database")

	dbGroup.Post("/", createRecord)
	dbGroup.Get("/", getAllRecords)
	dbGroup.Get("/:id", getRecordByID)
	dbGroup.Put("/:id", updateRecord)
	dbGroup.Delete("/:id", deleteRecord)
}

// Создание записи
func createRecord(c *fiber.Ctx) error {
	file := new(models.File)

	if err := c.BodyParser(file); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Неверный формат запроса"})
	}

	if result := db.DB.Create(&file); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Не удалось создать запись"})
	}

	return c.Status(fiber.StatusCreated).JSON(file)
}

// Получение всех записей
func getAllRecords(c *fiber.Ctx) error {
	var files []models.File

	if result := db.DB.Find(&files); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка при получении записей"})
	}

	return c.JSON(files)
}

// Получение записи по ID
func getRecordByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var file models.File

	if result := db.DB.First(&file, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Запись не найдена"})
	}

	return c.JSON(file)
}

// Обновление записи
func updateRecord(c *fiber.Ctx) error {
	id := c.Params("id")
	var file models.File

	if result := db.DB.First(&file, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Запись не найдена"})
	}

	var updateData models.File
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Неверный формат запроса"})
	}

	file.Name = updateData.Name
	file.FileURL = updateData.FileURL

	if result := db.DB.Save(&file); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Не удалось обновить запись"})
	}

	return c.JSON(file)
}

// Удаление записи
func deleteRecord(c *fiber.Ctx) error {
	id := c.Params("id")
	var file models.File

	if result := db.DB.First(&file, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Запись не найдена"})
	}

	if result := db.DB.Delete(&file); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Не удалось удалить запись"})
	}

	return c.JSON(fiber.Map{"message": "Запись удалена"})
}
