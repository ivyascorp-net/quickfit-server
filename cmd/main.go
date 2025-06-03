package main

import (
	"quickfit-server/initializers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	initializers.ConnectDatabase()
	initializers.MigrateDatabase()
}

func main() {
	e := echo.New()

	// Add Echo's built-in logger middleware
	e.Use(middleware.Logger())

	registerWorkoutRoutes(e)
	registerExerciseRoutes(e)
	registerWorkoutExerciseRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
