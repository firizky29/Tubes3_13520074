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
	DNA               string          `gorm:"type:text;NOT NULL"`
	DiseasePrediction string          `gorm:"type:varchar(255);NOT NULL"`
	SimilarityLevel   decimal.Decimal `gorm:"type:decimal(5, 2); NOT NULL"`
	Status            string          `gorm:"type:varchar(10);NOT NULL"`
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}
