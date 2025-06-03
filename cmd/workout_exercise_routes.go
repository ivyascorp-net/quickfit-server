package main

import (
	"quickfit-server/controllers"

	"github.com/labstack/echo/v4"
)

// registerWorkoutExerciseRoutes sets up all workout-exercise bridge routes.
func registerWorkoutExerciseRoutes(e *echo.Echo) {
	e.POST("/workout_exercises", controllers.CreateWorkoutExercise)
	e.GET("/workout_exercises/:id", controllers.GetWorkoutExercise)
	e.PUT("/workout_exercises/:id", controllers.UpdateWorkoutExercise)
	e.DELETE("/workout_exercises/:id", controllers.DeleteWorkoutExercise)
	e.GET("/workout_exercises/:workout_id", controllers.ListWorkoutExercises)
}
