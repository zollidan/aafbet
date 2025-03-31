package routes

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/zollidan/aafbet/s3client"
)


func APIS3(api fiber.Router) {

	api.Get("/s3/files", func(c *fiber.Ctx) error {

		client := s3client.GetClient()

		output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
			Bucket: &[]string{os.Getenv("AWS_BUCKET")}[0],
		})
		if err != nil {
			return c.Status(500).SendString("Ошибка при получении списка файлов: " + err.Error())
		}
	
		files := make([]fiber.Map, 0)
		for _, item := range output.Contents {
			files = append(files, fiber.Map{
				"key":          *item.Key,
				"size":         item.Size,
				"lastModified": item.LastModified,
			})
		}
	
		return c.JSON(files)
	}) 

	api.Get("/s3/files/:file_key", func(c *fiber.Ctx) error {

		// fileID := c.Params("file_id")
		// client := s3client.GetClient()
	
		// output, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		// 	Bucket: &[]string{os.Getenv("AWS_BUCKET")}[0],
		// 	Key:    &fileID,
		// })
		// if err != nil {
		// 	return c.Status(500).SendString("Ошибка при получении файла: " + err.Error())
		// }
		// defer output.Body.Close()
	
		// buf := make([]byte, output.ContentLength)
		// _, err = output.Body.Read(buf)
		// if err != nil {
		// 	return c.Status(500).SendString("Ошибка чтения файла: " + err.Error())
		// }
	
		// c.Type(*output.ContentType)
		// return c.Send(buf)

		return c.JSON(fiber.Map{
			"input": "hello",
		})
	})
}
