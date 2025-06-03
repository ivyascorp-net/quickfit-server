package models

type Exercise struct {
	ID           int               `json:"id"`
	Name         string            `json:"name" gorm:"type:varchar(191);uniqueIndex"`
	Description  string            `json:"description,omitempty"`
	Difficulty   string            `json:"difficulty,omitempty"` // e.g., Beginner, Intermediate, Advanced
	Language     string            `json:"language,omitempty"`   // e.g., en, fr
	Author       string            `json:"author,omitempty"`     // Creator or last editor
	Status       string            `json:"status,omitempty"`     // e.g., active, pending
	Tags         string            `json:"tags,omitempty"`       // Comma-separated tags
	Repetitions  int               `json:"repetitions,omitempty"`
	Sets         int               `json:"sets,omitempty"`
	CreatedAt    string            `json:"created_at,omitempty"`
	UpdatedAt    string            `json:"updated_at,omitempty"`
	CategoryID   *int              `json:"category_id,omitempty" gorm:"index"` // Foreign key to ExerciseCategory
	Category     *ExerciseCategory `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	EquipmentIDs []int             `json:"equipment_ids,omitempty" gorm:"-"` // Used for binding
	Equipment    []Equipment       `json:"equipment,omitempty" gorm:"many2many:exercise_equipments;"`
}

type Workout struct {
	ID        int    `json:"id"`
	Name      string `json:"name" gorm:"type:varchar(191);uniqueIndex"` // Unique index with length for MySQL
	Duration  int    `json:"duration"`                                  // in minutes
	Notes     string `json:"notes,omitempty"`                           // optional notes for the workout
	CreatedAt string `json:"created_at,omitempty"`                      // timestamp of creation
	// Remove direct exercises slice, use WorkoutExercise as bridge
}

type WorkoutExercise struct {
	ID          int      `json:"id"`
	WorkoutID   int      `json:"workout_id" gorm:"index"`                                       // Foreign key to Workout
	ExerciseID  int      `json:"exercise_id" gorm:"index"`                                      // Foreign key to Exercise
	Repetitions int      `json:"repetitions,omitempty"`                                         // number of repetitions for this exercise in the workout
	Sets        int      `json:"sets,omitempty"`                                                // number of sets for this exercise in the workout
	CreatedAt   string   `json:"created_at,omitempty"`                                          // timestamp of creation
	UpdatedAt   string   `json:"updated_at,omitempty"`                                          // timestamp of last update
	Exercise    Exercise `json:"exercise,omitempty" gorm:"foreignKey:ExerciseID;references:ID"` // Include Exercise details
	Workout     Workout  `json:"workout,omitempty" gorm:"foreignKey:WorkoutID;references:ID"`   // Include Workout details
}
