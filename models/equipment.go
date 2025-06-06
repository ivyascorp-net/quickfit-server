package models

// Equipment represents a piece of equipment used for exercises (e.g., Dumbbell, Barbell, Bodyweight)
type Equipment struct {
	Base
	Name string `json:"name" gorm:"type:varchar(100);uniqueIndex"`
}

// ExerciseCategory represents a category for exercises (e.g., Strength, Cardio)
type ExerciseCategory struct {
	Base
	Name string `json:"name" gorm:"type:varchar(100);uniqueIndex"`
}
