package models

type Muscle struct {
	Base
	Name string `json:"name" gorm:"type:varchar(100);uniqueIndex"` // Unique name for the muscle
}
