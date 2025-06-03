package main

import (
	"quickfit-server/controllers"

	"github.com/labstack/echo/v4"
)

// registerExerciseRoutes sets up all exercise-related routes.
func registerExerciseRoutes(e *echo.Echo) {
	e.POST("/exercises", controllers.CreateExercise)
	e.GET("/exercises/:id", controllers.GetExercise)
	e.PUT("/exercises/:id", controllers.UpdateExercise)
	e.DELETE("/exercises/:id", controllers.DeleteExercise)
	e.GET("/exercises", controllers.ListExercises)
}
