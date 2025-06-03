package main

import (
	"quickfit-server/controllers"

	"github.com/labstack/echo/v4"
)

// registerWorkoutRoutes sets up all workout-related routes.
func registerWorkoutRoutes(e *echo.Echo) {
	e.POST("/workouts", controllers.CreateWorkout)
	e.GET("/workouts/:id", controllers.GetWorkout)
	e.PUT("/workouts/:id", controllers.UpdateWorkout)
	e.DELETE("/workouts/:id", controllers.DeleteWorkout)
	e.GET("/workouts", controllers.ListWorkouts)
}
