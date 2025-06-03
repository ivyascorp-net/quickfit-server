package main

import (
	"encoding/json"
	"log"
	"os"
)

type WgerTranslation struct {
	Model  string `json:"model"`
	Pk     int    `json:"pk"`
	Fields struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Exercise    int    `json:"exercise"`
		UUID        string `json:"uuid"`
		Language    int    `json:"language"`
	} `json:"fields"`
}

type WgerExercise struct {
	Model  string `json:"model"`
	Pk     int    `json:"pk"`
	Fields struct {
		LicenseAuthor string `json:"license_author"`
		Category      int    `json:"category"`
		Created       string `json:"created"`
		LastUpdate    string `json:"last_update"`
		Equipment     []int  `json:"equipment"`
		UUID          string `json:"uuid"`
	} `json:"fields"`
}

type MyExercise struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	CategoryID   int    `json:"category_id"`
	EquipmentIDs []int  `json:"equipment_ids"`
	Author       string `json:"author"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	Status       string `json:"status"`
	Tags         string `json:"tags"`
	Repetitions  *int   `json:"repetitions"`
	Sets         *int   `json:"sets"`
}

func main() {
	data, err := os.ReadFile("rawdata/exercise-base-data.json")
	if err != nil {
		log.Fatal(err)
	}
	var raw []map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		log.Fatal(err)
	}

	// Load translations for names and descriptions
	transData, err := os.ReadFile("rawdata/translation.json")
	if err != nil {
		log.Fatal(err)
	}
	var transRaw []WgerTranslation
	if err := json.Unmarshal(transData, &transRaw); err != nil {
		log.Fatal(err)
	}
	// Build a map: uuid -> English translation (language=2), fallback to any language if not found
	type transInfo struct {
		Name        string
		Description string
	}
	transMap := make(map[string]transInfo)
	fallbackMap := make(map[string]transInfo)
	for _, t := range transRaw {
		if t.Model == "exercises.translation" && t.Fields.Name != "" && t.Fields.UUID != "" {
			if t.Fields.Language == 2 {
				transMap[t.Fields.UUID] = transInfo{t.Fields.Name, t.Fields.Description}
			} else if _, exists := fallbackMap[t.Fields.UUID]; !exists {
				fallbackMap[t.Fields.UUID] = transInfo{t.Fields.Name, t.Fields.Description}
			}
		}
	}

	var out []MyExercise
	for _, entry := range raw {
		if entry["model"] == "exercises.exercise" {
			var w WgerExercise
			b, _ := json.Marshal(entry)
			if err := json.Unmarshal(b, &w); err != nil {
				continue
			}
			tr, ok := transMap[w.Fields.UUID]
			if !ok {
				tr, ok = fallbackMap[w.Fields.UUID]
			}
			if !ok || tr.Name == "" {
				continue // skip exercises with no name
			}
			out = append(out, MyExercise{
				ID:           w.Pk,
				Name:         tr.Name,
				Description:  tr.Description,
				CategoryID:   w.Fields.Category,
				EquipmentIDs: w.Fields.Equipment,
				Author:       w.Fields.LicenseAuthor,
				CreatedAt:    w.Fields.Created,
				UpdatedAt:    w.Fields.LastUpdate,
				Status:       "active",
				Tags:         "",
				Repetitions:  nil,
				Sets:         nil,
			})
		}
	}
	if len(out) == 0 {
		log.Println("No exercises matched. Check translation and exercise UUIDs.")
	}
	outData, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("rawdata/exercises-filtered.json", outData, 0644); err != nil {
		log.Fatal(err)
	}
	log.Printf("Filtered %d exercises written to rawdata/exercises-filtered.json", len(out))
}
