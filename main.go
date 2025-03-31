package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/zollidan/aafbet/routes"
)

func main() {

	file_log, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(file_log)

    app := fiber.New(fiber.Config{
		AppName: "aafbet API 0.0.1",
	})

    api := app.Group("/api")

    routes.APIConnectionCheck(api)
	routes.APIS3(api)
	routes.APIDatabase(api)
	

    log.Fatalln(app.Listen("127.0.0.1:3000"))
}
