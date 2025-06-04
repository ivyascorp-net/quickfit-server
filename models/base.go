package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.DB) error {

	if base.ID.String() == "" {
		base.ID = uuid.New()
	}

	if err := base.IsValid(); err != nil {
		return err
	}

	return nil

}

func (Base *Base) IsValid() error {
	return nil
}
