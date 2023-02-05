package entity

import (
	"time"

	"github.com/google/uuid"
)

type Price struct {
	Id        uint      `db:"id" json:"id"`
	Uuid      uuid.UUID `db:"uuid" json:"uuid"`
	Name      string    `db:"name" json:"name"`
	PriceType uint      `db:"price_type" json:"price_type"` // 0: guest(default), 1: normal, 2: admin
	Cuopon    uint      `db:"cuopon" json:"cuopon"`         // 0: no, 1: yes

	TotalScore uint `db:"total_score" json:"total_score"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`

	// relative to other tables
	ServiceId uint `db:"service_id" json:"service_id"`

	// relative to other tables
	// Prices []Price `db:"prices" json:"prices"`

	// conv
	View           uint `db:"view" json:"view"`
	Click          uint `db:"click" json:"click"`
	RequestDoc     uint `db:"request_doc" json:"request_doc"`
	RequestAppoint uint `db:"request_appoint" json:"request_appoint"`
	Order          uint `db:"order" json:"order"`
	OrderType      uint `db:"order_type" json:"order_type"` // 0: free, 1: paid, 2: keep
	Keep           uint `db:"keep" json:"keep"`
	KeepTerm       uint `db:"keep_term" json:"keep_term"` // 0: 1month, 1: 3month, 2: 6month, 3: 1year

	// not in db
	ClickPerView float64 `db:"click_per_view" json:"click_per_view"`

	RequestDocPerView      float64 `db:"request_doc_per_view" json:"request_doc_per_view"`
	RequestDocPerClick     float64 `db:"request_doc_per_click" json:"request_doc_per_click"`
	RequestAppointPerView  float64 `db:"request_appoint_per_view" json:"request_appoint_per_view"`
	RequestAppointPerClick float64 `db:"request_appoint_per_click" json:"request_appoint_per_click"`

	OrderPerView  float64 `db:"order_per_view" json:"order_per_view"`
	OrderPerClick float64 `db:"order_per_click" json:"order_per_click"`

	KeepPerView  float64 `db:"keep_per_view" json:"keep_per_view"`
	KeepPerClick float64 `db:"keep_per_click" json:"keep_per_click"`

	CostPerView           float64 `db:"cost_per_view" json:"cost_per_view"`
	CostPerClick          float64 `db:"cost_per_click" json:"cost_per_click"`
	CostPerRequestDoc     float64 `db:"cost_per_request_doc" json:"cost_per_request_doc"`
	CostPerRequestAppoint float64 `db:"cost_per_request_appoint" json:"cost_per_request_appoint"`
	CostPerOrder          float64 `db:"cost_per_order" json:"cost_per_order"`
	CostPerKeep           float64 `db:"cost_per_keep" json:"cost_per_keep"`
}

func NewPrice(
	name string,
	priceType uint,
) *Price {
	return &Price{
		Name:      name,
		PriceType: priceType,
	}
}
