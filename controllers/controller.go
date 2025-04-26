package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
}

func (pc *Controller) GetHello(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "GoWave",
	}
	return c.Render(http.StatusOK, "welcome.html", data)
}
