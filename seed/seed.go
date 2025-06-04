package main

import (
	"encoding/json"
	"log"
	"os"
	"quickfit-server/initializers"
	"quickfit-server/models"
)

func main() {
	initializers.ConnectDatabase()
	initializers.MigrateDatabase() // Always migrate to ensure tables exist

	// Clean tables before seeding to avoid duplicate primary key errors
	initializers.DB.Exec("DELETE FROM muscles")
	initializers.DB.Exec("DELETE FROM exercise_categories")
	initializers.DB.Exec("DELETE FROM equipment")
	initializers.DB.Exec("DELETE FROM exercises")

	// Seed categories
	muscle, err := fetchMusclesRawData("/home/ivyas/ivayscorp-net/quickfit/quickfit-server/rawdata/muscles.json")
	if err != nil {
		log.Println("error loading file", err)
	}

	if err := initializers.DB.CreateInBatches(muscle, 10).Error; err != nil {
		log.Println("[ERROR] Failed to seed muscle: ", err)
	}

	exerciseCategories, err := fetchExerciseCategoriesRawData("/home/ivyas/ivayscorp-net/quickfit/quickfit-server/rawdata/exercise-categories.json")
	if err != nil {
		log.Println("error loading file", err)
	}
	if err := initializers.DB.CreateInBatches(exerciseCategories, 10).Error; err != nil {
		log.Fatalf("[ERROR] Failed to seed muscle: %v", err)
	}
	// Seed equipment
	equipment, err := fetchEquipmentRawData("/home/ivyas/ivayscorp-net/quickfit/quickfit-server/rawdata/exercise-equipments.json")
	if err != nil {
		log.Println("error loading file", err)
	}
	if err := initializers.DB.CreateInBatches(equipment, 10).Error; err != nil {
		log.Fatalf("[ERROR] Failed to seed equipment: %v", err)
	}

	// Seed exercises
	exercises, err := fetchExercisesRawData("/home/ivyas/ivayscorp-net/quickfit/quickfit-server/rawdata/exercises.json")
	if err != nil {
		log.Println("error loading file", err)
	}
	if err := initializers.DB.CreateInBatches(exercises, 10).Error; err != nil {
		log.Fatalf("[ERROR] Failed to seed exercises: %v", err)
	}

	log.Println("Database reseeded!")
}

func fetchMusclesRawData(filename string) (data []models.Muscle, err error) {

	file, err := os.ReadFile(filename)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		return data, err
	}

	return data, err

}
func fetchExerciseCategoriesRawData(filename string) (data []models.ExerciseCategory, err error) {
	file, err := os.ReadFile(filename)

	if err != nil {
		return data, err
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		return data, err
	}

	return data, err

}
func fetchEquipmentRawData(filename string) (data []models.Equipment, err error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		return data, err
	}
	return data, err

}

func fetchExercisesRawData(filename string) (data []models.Exercise, err error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		return data, err
	}
	return data, err

}
