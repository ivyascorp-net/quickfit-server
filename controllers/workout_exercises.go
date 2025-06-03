package controllers

import (
	"net/http"
	"quickfit-server/initializers"
	"quickfit-server/models"

	"github.com/labstack/echo/v4"
)

func CreateWorkoutExercise(c echo.Context) error {
	var workoutExercise models.WorkoutExercise
	if err := c.Bind(&workoutExercise); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	if err := initializers.DB.Create(&workoutExercise).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, workoutExercise)
}

// GetWorkoutExercise retrieves a single workout exercise by its ID.
func GetWorkoutExercise(c echo.Context) error {
	id := c.Param("id")
	var workoutExercise models.WorkoutExercise
	if err := initializers.DB.First(&workoutExercise, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Workout exercise not found"})
	}
	return c.JSON(http.StatusOK, workoutExercise)
}

// UpdateWorkoutExercise updates an existing workout exercise by its ID.
func UpdateWorkoutExercise(c echo.Context) error {
	id := c.Param("id")
	var workoutExercise models.WorkoutExercise
	if err := initializers.DB.First(&workoutExercise, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Workout exercise not found"})
	}

	if err := c.Bind(&workoutExercise); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	if err := initializers.DB.Save(&workoutExercise).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, workoutExercise)
}

// DeleteWorkoutExercise deletes a workout exercise by its ID.
func DeleteWorkoutExercise(c echo.Context) error {
	id := c.Param("id")
	var workoutExercise models.WorkoutExercise
	if err := initializers.DB.First(&workoutExercise, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Workout exercise not found"})
	}

	if err := initializers.DB.Delete(&workoutExercise).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Workout exercise deleted successfully"})
}

// ListWorkoutExercises retrieves all workout exercises for a specific workout.
func ListWorkoutExercises(c echo.Context) error {
	workoutID := c.Param("workout_id")
	var workoutExercises []models.WorkoutExercise
	if err := initializers.DB.Where("workout_id = ?", workoutID).Find(&workoutExercises).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	if len(workoutExercises) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "No workout exercises found for this workout"})
	}

	return c.JSON(http.StatusOK, workoutExercises)
}

// ListAllWorkoutExercises retrieves all workout exercises across all workouts.
func ListAllWorkoutExercises(c echo.Context) error {
	var workoutExercises []models.WorkoutExercise
	if err := initializers.DB.Find(&workoutExercises).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	if len(workoutExercises) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "No workout exercises found"})
	}

	return c.JSON(http.StatusOK, workoutExercises)
}
