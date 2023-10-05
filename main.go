package main

import (
	"eticketing/handler"
	"eticketing/repository"
	"eticketing/service"
	"eticketing/utils"
	"os"

	_ "github.com/joho/godotenv/autoload"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Pre(middleware.RemoveTrailingSlash())
	authenticated := e.Group("/v1/terminal")
	authenticated.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SIGNED_TOKEN")),
	}))

	db, err := utils.ConnectDB()
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	tRepo := repository.NewTerminalRepository(db)
	tService := service.NewTerminalService(tRepo)
	tHandler := handler.NewTerminalHandler(tService)

	// Routes
	e.POST("/v1/register", userHandler.Create)
	e.POST("/v1/login", userHandler.Login)
	authenticated.POST("", tHandler.Create)
	authenticated.GET("", tHandler.GetAll)

	// Start server
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
