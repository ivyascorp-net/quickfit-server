package main

import (
	"log"
	"quickfit-server/initializers"
	"quickfit-server/models"
)

func main() {
	initializers.ConnectDatabase()
	initializers.MigrateDatabase() // Always migrate to ensure tables exist

	// Truncate tables safely (including new normalized tables)
	tables := []string{"workout_exercises", "exercises", "workouts", "equipment", "exercise_categories", "exercise_equipments"}
	initializers.DB.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	for _, table := range tables {
		if err := initializers.DB.Exec("TRUNCATE TABLE " + table + ";").Error; err != nil {
			log.Printf("[WARN] Could not truncate table %s: %v", table, err)
		}
	}
	initializers.DB.Exec("SET FOREIGN_KEY_CHECKS = 1;")

	// Seed categories
	category := models.ExerciseCategory{Name: "Strength"}
	if err := initializers.DB.Create(&category).Error; err != nil {
		log.Fatalf("[ERROR] Failed to seed category: %v", err)
	}

	// Seed equipment
	equipment := []models.Equipment{
		{Name: "Bodyweight"},
		{Name: "Mat"},
	}
	for i := range equipment {
		if err := initializers.DB.Create(&equipment[i]).Error; err != nil {
			log.Fatalf("[ERROR] Failed to seed equipment %s: %v", equipment[i].Name, err)
		}
	}

	// Seed exercises with relations
	exercises := []models.Exercise{
		{
			Name:        "Push Ups",
			Description: "Standard push ups",
			Repetitions: 15,
			Sets:        3,
			CategoryID:  &category.ID,
			Equipment:   []models.Equipment{equipment[0]},
		},
		{
			Name:        "Plank",
			Description: "Core strength hold",
			Repetitions: 1,
			Sets:        3,
			CategoryID:  &category.ID,
			Equipment:   []models.Equipment{equipment[1]},
		},
	}
	for i := range exercises {
		if err := initializers.DB.Create(&exercises[i]).Error; err != nil {
			log.Fatalf("[ERROR] Failed to seed exercise %s: %v", exercises[i].Name, err)
		}
	}

	// Seed workouts
	workout := models.Workout{
		Name:     "Morning Routine",
		Duration: 45,
		Notes:    "Full body warmup",
	}
	if err := initializers.DB.Create(&workout).Error; err != nil {
		log.Fatalf("[ERROR] Failed to seed workout: %v", err)
	}

	// Link exercises to workout via WorkoutExercise
	links := []models.WorkoutExercise{
		{
			WorkoutID:   workout.ID,
			ExerciseID:  exercises[0].ID,
			Repetitions: 15,
			Sets:        3,
		},
		{
			WorkoutID:   workout.ID,
			ExerciseID:  exercises[1].ID,
			Repetitions: 1,
			Sets:        3,
		},
	}
	for i := range links {
		if err := initializers.DB.Create(&links[i]).Error; err != nil {
			log.Fatalf("[ERROR] Failed to link exercise %d to workout: %v", links[i].ExerciseID, err)
		}
	}

	log.Println("Database reseeded!")
}
