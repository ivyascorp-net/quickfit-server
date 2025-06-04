package models

import (
	"github.com/google/uuid"
)

type Exercise struct {
	Base
	Name             string `json:"name" gorm:"type:varchar(191);uniqueIndex"`
	Description      string `json:"description,omitempty"`
	Difficulty       string `json:"difficulty,omitempty"` // e.g., Beginner, Intermediate, Advanced
	Force            string `json:"force,omitempty"`      // e.g., Push, Pull, Legs
	Language         string `json:"language,omitempty"`   // e.g., en, fr
	Author           string `json:"author,omitempty"`     // Creator or last editor
	Status           string `json:"status,omitempty"`     // e.g., active, pending
	Tags             string `json:"tags,omitempty"`       // Comma-separated tags
	Repetitions      int    `json:"repetitions,omitempty"`
	Sets             int    `json:"sets,omitempty"`
	CreatedAt        string `json:"created_at,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
	Mechanic         string `json:"mechanic,omitempty"`                                  // e.g., Isolated, Compound
	Images           string `json:"images,omitempty" gorm:"type:text"`                   // Store as JSON string or comma-separated
	PrimaryMuscles   string `json:"primary_muscles,omitempty" gorm:"type:text"`          // Store as JSON string or comma-separated
	SecondaryMuscles string `json:"secondary_muscles,omitempty" gorm:"type:text"`        // Store as JSON string or comma-separated
	EquipmentID      string `json:"equipment_id,omitempty" gorm:"type:varchar(36)"`      // Foreign key to Equipment
	CategoryID       string `json:"category_id,omitempty" gorm:"type:varchar(36);index"` // Foreign key to ExerciseCategory
	Instructions     string `json:"instructions,omitempty" gorm:"type:text"`             // Store as JSON string or comma-separated
}

type Workout struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(191);uniqueIndex"` // Unique index with length for MySQL
	Duration  int       `json:"duration"`                                  // in minutes
	Notes     string    `json:"notes,omitempty"`                           // optional notes for the workout
	CreatedAt string    `json:"created_at,omitempty"`                      // timestamp of creation
	// Remove direct exercises slice, use WorkoutExercise as bridge
}

type WorkoutExercise struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	WorkoutID   uuid.UUID `json:"workout_id" gorm:"type:uuid;index"`                             // Foreign key to Workout
	ExerciseID  uuid.UUID `json:"exercise_id" gorm:"type:uuid;index"`                            // Foreign key to Exercise
	Repetitions int       `json:"repetitions,omitempty"`                                         // number of repetitions for this exercise in the workout
	Sets        int       `json:"sets,omitempty"`                                                // number of sets for this exercise in the workout
	CreatedAt   string    `json:"created_at,omitempty"`                                          // timestamp of creation
	UpdatedAt   string    `json:"updated_at,omitempty"`                                          // timestamp of last update
	Exercise    Exercise  `json:"exercise,omitempty" gorm:"foreignKey:ExerciseID;references:ID"` // Include Exercise details
	Workout     Workout   `json:"workout,omitempty" gorm:"foreignKey:WorkoutID;references:ID"`   // Include Workout details
}
