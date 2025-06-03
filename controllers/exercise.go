package controllers

import (
	"net/http"
	"quickfit-server/initializers"
	"quickfit-server/models"

	"log"

	"github.com/labstack/echo/v4"
)

// CreateExercise handles the creation of a new exercise. It checks for duplicate names and requires an ID.
func CreateExercise(c echo.Context) error {

	var exercise models.Exercise
	if err := c.Bind(&exercise); err != nil {
		log.Printf("[ERROR] Failed to bind exercise: %v\n", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if exercise.ID == 0 {
		log.Printf("[WARN] Exercise ID is missing")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Exercise ID is required"})
	}

	// Check if exercise with same name exists
	var existing models.Exercise
	err := initializers.DB.Where("name = ?", exercise.Name).First(&existing).Error
	if err == nil {
		log.Printf("[WARN] Exercise already exists: %s\n", exercise.Name)
		return c.JSON(http.StatusConflict, echo.Map{"error": "Exercise already exists"})
	}
	if err.Error() != "record not found" && err.Error() != "gorm: record not found" {
		log.Printf("[ERROR] DB error: %v\n", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	if err := initializers.DB.Create(&exercise).Error; err != nil {
		log.Printf("[ERROR] Failed to create exercise: %v\n", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	log.Printf("[INFO] Exercise created: %+v\n", exercise)
	return c.JSON(http.StatusCreated, exercise)
}

// GetExercise retrieves a single exercise by its ID.
func GetExercise(c echo.Context) error {
	id := c.Param("id")
	var exercise models.Exercise
	if err := initializers.DB.First(&exercise, id).Error; err != nil {
		log.Printf("[ERROR] Exercise not found: %s\n", id)
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Exercise not found"})
	}

	log.Printf("[INFO] Exercise retrieved: %+v\n", exercise)
	return c.JSON(http.StatusOK, exercise)
}

// UpdateExercise updates an existing exercise by its ID.
func UpdateExercise(c echo.Context) error {
	id := c.Param("id")
	var exercise models.Exercise
	if err := initializers.DB.First(&exercise, id).Error; err != nil {
		log.Printf("[ERROR] Exercise not found: %s\n", id)
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Exercise not found"})
	}

	if err := c.Bind(&exercise); err != nil {
		log.Printf("[ERROR] Failed to bind exercise: %v\n", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := initializers.DB.Save(&exercise).Error; err != nil {
		log.Printf("[ERROR] Failed to update exercise: %v\n", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	log.Printf("[INFO] Exercise updated: %+v\n", exercise)
	return c.JSON(http.StatusOK, exercise)
}

// DeleteExercise deletes an exercise by its ID.
func DeleteExercise(c echo.Context) error {
	id := c.Param("id")
	var exercise models.Exercise
	if err := initializers.DB.First(&exercise, id).Error; err != nil {
		log.Printf("[ERROR] Exercise not found: %s\n", id)
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Exercise not found"})
	}

	if err := initializers.DB.Delete(&exercise).Error; err != nil {
		log.Printf("[ERROR] Failed to delete exercise: %v\n", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	log.Printf("[INFO] Exercise deleted: %s\n", id)
	return c.NoContent(http.StatusNoContent)
}

// ListExercises returns a list of all exercises.
func ListExercises(c echo.Context) error {
	var exercises []models.Exercise
	if err := initializers.DB.Find(&exercises).Error; err != nil {
		log.Printf("[ERROR] Failed to retrieve exercises: %v\n", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	log.Printf("[INFO] Exercises retrieved: %d found\n", len(exercises))
	return c.JSON(http.StatusOK, exercises)
}

// ListExercises returns a list of all exercises
