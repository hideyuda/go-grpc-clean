package entity

import (
	"time"

	"github.com/google/uuid"
)

type PriceCandle struct {
	Id              uint      `db:"id" json:"id"`
	Uuid            uuid.UUID `db:"uuid" json:"uuid"`
	Name            string    `db:"name" json:"name"`
	PriceCandleType uint      `db:"PriceCandle_type" json:"PriceCandle_type"` // 0: guest(default), 1: normal, 2: admin
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`

	// relative to other tables
	ServiceId uint `db:"service_id" json:"service_id"`

	// relative to other tables
	Prices []Price `db:"prices" json:"prices"`

	// not exist in db
	AveragePrice        float64 `db:"average_price" json:"average_price"`
	AverageCostPerClick float64 `db:"average_cost_per_click" json:"average_cost_per_click"`
	AverageClicks       float64 `db:"average_clicks" json:"average_clicks"`
	AverageViews        float64 `db:"average_views" json:"average_views"`
	AverageOrder        float64 `db:"average_order" json:"average_order"`
	AverageOrderRate    float64 `db:"average_order_rate" json:"average_order_rate"`
	AverageKeep         float64 `db:"average_keep" json:"average_keep"`

	AverageRevenue    float64 `db:"average_revenue" json:"average_revenue"`
	AverageProfit     float64 `db:"average_profit" json:"average_profit"`
	AverageProfitRate float64 `db:"average_profit_rate" json:"average_profit_rate"`

	AnnualRevenue    float64 `db:"annual_revenue" json:"annual_revenue"`
	AnnualProfit     float64 `db:"annual_profit" json:"annual_profit"`
	AnnualProfitRate float64 `db:"annual_profit_rate" json:"annual_profit_rate"`

	MonthlyRevenue    float64 `db:"monthly_revenue" json:"monthly_revenue"`
	MonthlyProfit     float64 `db:"monthly_profit" json:"monthly_profit"`
	MonthlyProfitRate float64 `db:"monthly_profit_rate" json:"monthly_profit_rate"`

	WeeklyRevenue    float64 `db:"weekly_revenue" json:"weekly_revenue"`
	WeeklyProfit     float64 `db:"weekly_profit" json:"weekly_profit"`
	WeeklyProfitRate float64 `db:"weekly_profit_rate" json:"weekly_profit_rate"`

	DailyRevenue    float64 `db:"daily_revenue" json:"daily_revenue"`
	DailyProfit     float64 `db:"daily_profit" json:"daily_profit"`
	DailyProfitRate float64 `db:"daily_profit_rate" json:"daily_profit_rate"`

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

func NewPriceCandle(
	name string,
	PriceCandleType uint,
) *PriceCandle {
	return &PriceCandle{
		Name:            name,
		PriceCandleType: PriceCandleType,
	}
}
