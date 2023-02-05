package entity

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	Id          uint      `db:"id" json:"id"`
	Uuid        uuid.UUID `db:"uuid" json:"uuid"`
	Name        string    `db:"name" json:"name"`
	ServiceType uint      `db:"service_type" json:"service_type"` // 0: guest(default), 1: normal, 2: admin
	MaxPrice    uint      `db:"max_price" json:"max_price"`
	MinPrice    uint      `db:"min_price" json:"min_price"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`

	// relative to other tables
	User_Id uint `db:"user_id" json:"user_id"`

	// relative to other tables
	PriceCandles []PriceCandle `db:"price_candles" json:"price_candles"`
	Conpetitors  []Competitor  `db:"competitors" json:"competitors"`

	// not exsist in db
	PriceCandleCount uint `db:"price_candle_count" json:"price_candle_count"`
	CompetitorCount  uint `db:"competitor_count" json:"competitor_count"`
}

func NewService(
	name string,
	serviceType uint,
) *Service {
	return &Service{
		Name:        name,
		ServiceType: serviceType,
	}
}
