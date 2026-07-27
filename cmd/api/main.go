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
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file, using environment variables")
	}
	postgresDatabase := config.NewPostgresConnection()

	_ = postgresDatabase

	elasticsearchClient := config.NewElasticsearchClient()

	elastic.BootstrapIndex(elasticsearchClient)

	productRepo := database.NewProductRepository(postgresDatabase)
	slotRepo := database.NewShelfSlotRepository(postgresDatabase)
	productElastic := elastic.NewProductRepository(elasticsearchClient)
	slotElastic := elastic.NewSlotRepository(elasticsearchClient)

	productService := service.NewProductService(productRepo, productElastic)
	slotService := service.NewSlotService(slotRepo, slotElastic)

	productHandler := handler.NewProductHandler(productService)
	slotHandler := handler.NewSlotHandler(slotService)
	router := api_server.NewRouter(productHandler, slotHandler)

	router.Run(":8083")

}
