package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Recipe struct {
	ID           string `gorm:"primaryKey;type:uuid" json:"id"`
	CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
	Title string `json:"title"`
	Ingredients string `json:"ingredients"`
	Instructions string `json:"instructions"`
}

func (r *Recipe) BeforeCreate(tx *gorm.DB) (err error) {
    r.ID = uuid.New().String()
    return
}