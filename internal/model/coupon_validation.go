package model

import "time"

type ValidateRequest struct {
	CouponCode   string    `json:"coupon_code"`
	UserID       string    `json:"user_id"`
	OrderValue   float64   `json:"order_value"`
	MedicineIDs  []string  `json:"medicine_ids"`
	CategoryIDs  []string  `json:"category_ids"`
	Timestamp    time.Time `json:"timestamp"`
}

type ValidateResponse struct {
	Valid          bool    `json:"valid"`
	DiscountAmount float64 `json:"discount_amount"`
	FinalPrice     float64 `json:"final_price"`
	Message        string  `json:"message"`
}
