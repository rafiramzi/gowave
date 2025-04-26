package routes

import (
	"gowave/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	controller := controllers.Controller{DB: db}

	e.GET("", controller.GetHello)
	
}