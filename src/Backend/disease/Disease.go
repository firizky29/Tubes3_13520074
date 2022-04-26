package disease

import (
	"time"

	"gorm.io/gorm"
)

type Disease struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100);unique;NOT NULL"`
	DNA       string `gorm:"type:varchar(255);unique;NOT NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
