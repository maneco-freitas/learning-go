package main

import (
	"api/internal/config"
	"api/internal/database"
	"api/internal/handler"
	"api/internal/repository"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	gin.SetMode(cfg.Server.GinMode)

	db := database.NewDBConnection(cfg)
	productRepository := repository.NewProductRepository(db)
	productHandler := handler.NewProductHandler(productRepository)

	r := gin.Default()
	r.SetTrustedProxies(nil)
	products := r.Group("/products")
	{
		products.GET("/", productHandler.GetAll)
		products.GET("/:id", productHandler.GetByID)
		products.POST("/", productHandler.Create)
		products.PUT("/:id", productHandler.Update)
		products.DELETE("/:id", productHandler.Delete)
	}

	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
