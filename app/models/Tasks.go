package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID           uint32    `gorm:"unique;not null;index;primary_key;auto_increment" json:"id" binding:"required"`
	Title        string    `gorm:"size:255;not null;" json:"title" binding:"required"`
	Content      string    `gorm:"size:255" json:"content" db:"text"`
	Priority     string    `json:"priority"`
	Category     string    `json:"category"`
	Referer      string    `json:"referer"`
	IsOverdue    bool      `json:"isoverdue,omitempty"`
	IsHidden     int       `json:"ishidden,omitempty`
	CompletedMsg string    `json:"ishidden,omitempty"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// func (u *User) Prepare() {
// 	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
// 	u.CreatedAt = time.Now()
// 	u.UpdatedAt = time.Now()
// }

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
