package utils

import (
	"log"
	"net/http"
	"os"

	"gowave/database"
	"gowave/middlewares"
	"gowave/routes"
	templates "gowave/utils/gowave_templates"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)
func Boot(){
	if len(os.Args) >= 2 && os.Args[1] == "generate" {
		generate()
		return
	}
	
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	if dbUser == "" || dbPassword == "" || dbName == "" {
		log.Println("⚠️ Skipping database connection: missing DB_USER, DB_PASSWORD, or DB_NAME")
	} else {
		log.Println("DB_NAME:", dbName)

		db, err := database.Init()
		if err != nil {
			log.Fatalf("Could not connect to the database: %v", err)
		}
		e := echo.New()
		e.Use(middlewares.ResponseMiddleware)
		routes.RegisterRoutes(e, db)
		port := os.Getenv("APP_PORT")
		if port == "" {
			port = "8080"
		}
		e.Logger.Fatal(e.Start(":" + port))
		return
	}
	e := echo.New()
	e.Renderer = NewRenderer("views/*.html")
	e.Static("/static", "static")

	e.Use(middlewares.ResponseMiddleware)
	routes.RegisterRoutes(e, nil)
	apiMode := os.Getenv("API_MODE")

	if(apiMode != "true"){
		e.HTTPErrorHandler = NotFoundHandler
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}

func NotFoundHandler(err error, c echo.Context) {
	code := 404
	var message string
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message.(string)
	} else {
		message = err.Error()
	}

	if code == 404 {
		c.HTML(http.StatusNotFound, templates.NotFound())
	} else {
		c.String(code, message)
	}
}

func generate() {
	if len(os.Args) < 4 {
		log.Fatalf("GoWave : Invalid arguments.")
	}

	folderName := os.Args[2]
	fileName := os.Args[3]

	Generate(folderName, fileName)
}