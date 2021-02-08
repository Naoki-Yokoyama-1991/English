package models

import (
	"time"
)

type Phrase struct {
	ID        uint32    `gorm:"unique;not null;index;primary_key;auto_increment" json:"id" binding:"required"`
	Phrase    string    `gorm:"size:255;not null;" json:"phrase" binding:"required"`
	Describe  string    `gorm:"size:255" json:"describe",omitempty"`
	Mastery   string    `json:"mastery,omitempty"`
	Archive   string    `json:"category,omitempty"`
	Complete  string    `json:"ishidden,omitempty"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
