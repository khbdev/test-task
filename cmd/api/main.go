package main

import (
	"log"
	api_server "test-task/api-server"
	"test-task/internal/config"
	"test-task/internal/handler"
	"test-task/internal/repostory/database"
	"test-task/internal/repostory/elastic"
	"test-task/internal/service"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	postgresDatabase := config.NewPostgresConnection()

	_ = postgresDatabase

	elasticsearchClient := config.NewElasticsearchClient()

	elastic.BootstrapIndex(elasticsearchClient)

	productRepo := database.NewProductRepository(postgresDatabase)
	productElastic := elastic.NewProductRepository(elasticsearchClient)

	productService := service.NewProductService(productRepo, productElastic)

	productHandler := handler.NewProductHandler(productService)

	router := api_server.NewRouter(productHandler)

	router.Run(":8083")

}
