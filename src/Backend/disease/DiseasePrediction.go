package disease

import (
	"time"

	"github.com/shopspring/decimal"

	"gorm.io/gorm"
)

type DiseasePrediction struct {
	ID                int             `gorm:"primaryKey"`
	CreatedAt         time.Time       `gorm:"type:DATE;NOT NULL"`
	UserName          string          `gorm:"type:varchar(100);NOT NULL"`
	DNA               string          `gorm:"type:varchar(255);NOT NULL"`
	DiseasePrediction string          `gorm:"type:varchar(255);NOT NULL"`
	SimilarityLevel   decimal.Decimal `gorm:"type:decimal(3, 2); NOT NULL"`
	Status            bool            `gorm:"type:tinyint(1);NOT NULL"`
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}
