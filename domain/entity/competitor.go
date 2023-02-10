package entity

import (
	"time"

	"github.com/google/uuid"
)

type Competitor struct {
	Id             uint      `db:"id" json:"id"`
	Uuid           uuid.UUID `db:"uuid" json:"uuid"`
	Name           string    `db:"name" json:"name"`
	CompetitorType uint      `db:"competitor_type" json:"competitor_type"` // 0: guest(default), 1: normal, 2: admin
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`

	OverPrice uint `db:"over_price" json:"over_price"` // 0: under, 1: over

	// relative to other tables
	User_Id uint `db:"user_id" json:"user_id"`

	// relative to other tables
	PriceCandleCount uint `db:"price_candle_count" json:"price_candle_count"`
}

func NewCompetitor(
	name string,
	CompetitorType uint,
) *Competitor {
	return &Competitor{
		Name:           name,
		CompetitorType: CompetitorType,
	}
}
