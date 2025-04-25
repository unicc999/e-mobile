// @title Effective Mobile тестовое задание.
// @version 1.0
// @description Сервис, который получает по API ФИО, из открытых API обогащает ответ наиболее вероятными возрастом, полом и национальностью и сохраняет данные в БД. По запросу выдает инфу о найденных людях.
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8081
package main

import (
	"e-mobile/config"
	"e-mobile/controllers"
	_ "e-mobile/docs"
	"e-mobile/models"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	db := config.InitDB()
	if err := models.Migrate(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	r := gin.Default()

	personController := controllers.NewPersonController(db)

	api := r.Group("/api")
	{
		api.POST("/people", personController.CreatePerson)
		api.GET("/people", personController.GetPeople)
		api.GET("/people/:id", personController.GetPersonByID)
		api.PUT("/people/:id", personController.UpdatePerson)
		api.DELETE("/people/:id", personController.DeletePerson)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
