package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"quickfit-server/models"
	"strings"

	"github.com/google/uuid"
)

type FreeExercise struct {
	Name             string   `json:"name"`
	Force            string   `json:"force"`
	Level            string   `json:"level"`
	Mechanic         string   `json:"mechanic"`
	Equipment        string   `json:"equipment"`
	PrimaryMuscles   []string `json:"primaryMuscles"`
	SecondaryMuscles []string `json:"secondaryMuscles"`
	Instructions     []string `json:"instructions"`
	Category         string   `json:"category"`
	Images           []string `json:"images"`
	ID               string   `json:"id"`
}

func main() {
	rawFile, err := os.ReadFile("/home/ivyas/ivayscorp-net/free-exercise-db/dist/exercises.json")
	if err != nil {
		log.Printf("Error reading raw exercises.json file, %v", err)
	}
	fex := []FreeExercise{}
	err = json.Unmarshal(rawFile, &fex)
	if err != nil {
		log.Fatal("error loading file in go struct", err)
	}
	filteredEuipment := buildEquipmentMap(fex)
	writeFilteredExercises("exercise-equipments", filteredEuipment)
	filteredCategories := buildExerciseCategories(fex)
	writeFilteredExercises("exercise-categories", filteredCategories)
	filteredMuscles := buildMuscles(fex)
	writeFilteredExercises("muscles", filteredMuscles)
	filteredExercises := buildExercise(fex, filteredMuscles, filteredCategories, filteredEuipment)
	writeFilteredExercises("exercises", filteredExercises)
}

func buildEquipmentMap(rawEx []FreeExercise) (filteredEquipments []models.Equipment) {

	ex := map[string]uuid.UUID{}

	for _, v := range rawEx {

		if _, ok := ex[v.Equipment]; !ok {

			ex[v.Equipment] = uuid.New()

		}

	}

	for k, v := range ex {
		eq := models.Equipment{
			Base: models.Base{ID: v},
			Name: k,
		}
		filteredEquipments = append(filteredEquipments, eq)

	}

	return filteredEquipments

}

func buildExerciseCategories(rawEX []FreeExercise) (filteredCategories []models.ExerciseCategory) {
	categoryMap := map[string]uuid.UUID{}

	for _, v := range rawEX {

		if _, ok := categoryMap[v.Category]; !ok {
			categoryMap[v.Category] = uuid.New()
		}

	}

	for k, v := range categoryMap {
		fc := models.ExerciseCategory{
			Base: models.Base{ID: v},
			Name: k,
		}
		filteredCategories = append(filteredCategories, fc)

	}

	return filteredCategories
}

func buildMuscles(rawEX []FreeExercise) (filteredMuscles []models.Muscle) {

	muscleMap := map[string]uuid.UUID{}

	for _, v := range rawEX {
		for _, pm := range v.PrimaryMuscles {

			if _, ok := muscleMap[pm]; !ok {

				muscleMap[pm] = uuid.New()
			}

		}
		for _, sm := range v.SecondaryMuscles {

			if _, ok := muscleMap[sm]; !ok {

				muscleMap[sm] = uuid.New()
			}

		}
	}
	for k, v := range muscleMap {
		m := models.Muscle{
			Base: models.Base{ID: v},
			Name: k,
		}
		filteredMuscles = append(filteredMuscles, m)
	}

	return filteredMuscles

}

func getMuscleID(pm []string, filteredMuscles []models.Muscle) (muscleID []uuid.UUID) {
	for _, muscle := range filteredMuscles {
		for _, m := range pm {
			if muscle.Name == m {
				muscleID = append(muscleID, muscle.ID)
			}
		}
	}
	return muscleID
}
func getCategoryID(category string, filteredCategories []models.ExerciseCategory) (categoryID uuid.UUID) {
	for _, cat := range filteredCategories {
		if cat.Name == category {
			categoryID = cat.ID
			return categoryID
		}
	}
	return uuid.Nil // Return zero UUID if not found
}
func getEquipmentID(equipment string, filteredEquipments []models.Equipment) (equipmentID uuid.UUID) {
	for _, eq := range filteredEquipments {
		if eq.Name == equipment {
			equipmentID = eq.ID
			return equipmentID
		}
	}
	return uuid.Nil // Return zero UUID if not found
}

func buildExercise(rawEx []FreeExercise, filteredMuscles []models.Muscle, filteredCategories []models.ExerciseCategory, filteredEquipments []models.Equipment) (filteredExercises []models.Exercise) {

	for _, v := range rawEx {
		primaryMuscleIDs := getMuscleID(v.PrimaryMuscles, filteredMuscles)
		secondaryMuscleIDs := getMuscleID(v.SecondaryMuscles, filteredMuscles)
		pm := []string{}
		sm := []string{}
		log.Println("Primary Muscles:", len(primaryMuscleIDs))
		log.Println("Secondary Muscles:", len(secondaryMuscleIDs))
		for _, v := range primaryMuscleIDs {
			if v == uuid.Nil {
				continue
			}
			pm = append(pm, v.String())

		}
		for _, v := range secondaryMuscleIDs {
			if v == uuid.Nil {
				continue
			}
			sm = append(sm, v.String())
		}
		ex := models.Exercise{
			Base:             models.Base{ID: uuid.New()},
			Name:             v.Name,
			Force:            v.Force,
			Difficulty:       v.Level,
			Mechanic:         v.Mechanic,
			EquipmentID:      getEquipmentID(v.Equipment, filteredEquipments).String(),
			PrimaryMuscles:   strings.Join(pm, ","),
			SecondaryMuscles: strings.Join(sm, ","),
			Instructions:     strings.Join(v.Instructions, ","),
			Images:           strings.Join(v.Images, ","),
			CategoryID:       getCategoryID(v.Category, filteredCategories).String(),
		}
		filteredExercises = append(filteredExercises, ex)
	}

	return filteredExercises
}

func writeFilteredExercises(fileName string, filtered any) {
	outData, err := json.MarshalIndent(filtered, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(fmt.Sprintf("/home/ivyas/ivayscorp-net/quickfit/quickfit-server/rawdata/%s.json", fileName), outData, 0644); err != nil {
		log.Fatal(err)
	}
	log.Printf("Filtered strength exercises written to %s.json", fileName)
}
