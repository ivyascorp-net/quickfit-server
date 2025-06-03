package controllers

import (
	"net/http"
	"quickfit-server/initializers"
	"quickfit-server/models"

	"github.com/labstack/echo/v4"
)

// ListWorkouts returns a list of all workouts.
func ListWorkouts(c echo.Context) error {
	var workouts []models.Workout
	if err := initializers.DB.Find(&workouts).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, workouts)
}

// GetWorkout retrieves a single workout by its ID.
func GetWorkout(c echo.Context) error {
	id := c.Param("id")
	var workout models.Workout
	if err := initializers.DB.First(&workout, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Workout not found"})
	}
	return c.JSON(http.StatusOK, workout)
}

// CreateWorkout handles the creation of a new workout.
func CreateWorkout(c echo.Context) error {
	var workout models.Workout
	if err := c.Bind(&workout); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	if err := initializers.DB.Create(&workout).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, workout)
}

// UpdateWorkout updates an existing workout by its ID.
func UpdateWorkout(c echo.Context) error {
	id := c.Param("id")
	var workout models.Workout
	if err := initializers.DB.First(&workout, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Workout not found"})
	}

	if err := c.Bind(&workout); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	if err := initializers.DB.Save(&workout).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, workout)
}

// DeleteWorkout deletes a workout by its ID., will also delete associated workout exercises
func DeleteWorkout(c echo.Context) error {

	id := c.Param("id")
	var workout models.Workout
	if err := initializers.DB.First(&workout, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Workout not found"})
	}
	// Delete associated workout exercises
	var workoutExercises []models.WorkoutExercise
	if err := initializers.DB.Where("workout_id = ?", id).Find(&workoutExercises).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve workout exercises"})
	}
	for _, we := range workoutExercises {
		if err := initializers.DB.Delete(&we).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete workout exercise"})
		}
	}

		if err := initializers.DB.Delete(&workout).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusNoContent, nil)
}
