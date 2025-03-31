package s3client

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

var (
	client     *s3.Client
	initClient sync.Once
)

func GetClient() *s3.Client {
	initClient.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Println("Файл .env не найден или не удалось загрузить")
		}	

		cfg, err := config.LoadDefaultConfig(context.TODO(), 
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			os.Getenv("AWS_ID"), 
			os.Getenv("AWS_SECRET_KEY"), 
			"")),
		config.WithBaseEndpoint(os.Getenv("AWS_ENDPOINT")),
		config.WithRegion(os.Getenv("AWS_REGION")),
		)

		if err != nil {
			log.Fatalf("Ошибка конфигурации AWS: %v", err)
		}

		client = s3.NewFromConfig(cfg)
	})

	return client
}