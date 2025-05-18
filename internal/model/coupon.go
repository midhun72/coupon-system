package model

import "time"

type TimeWindow struct {
	Start string `json:"start"` // "HH:MM"
	End   string `json:"end"`   // "HH:MM"
}

type Coupon struct {
	CouponCode            string     `json:"coupon_code"`
	ExpiryDate            time.Time  `json:"expiry_date"`
	UsageType             string     `json:"usage_type"` // "one_time", "multi_use", "time_based"
	ApplicableMedicineIDs []string   `json:"applicable_medicine_ids"`
	ApplicableCategories  []string   `json:"applicable_categories"`
	MinOrderValue         float64    `json:"min_order_value"`
	ValidTimeWindow       TimeWindow `json:"valid_time_window"`
	TermsAndConditions    string     `json:"terms_and_conditions"`
	DiscountType          string     `json:"discount_type"` // "percentage", "flat"
	DiscountValue         float64    `json:"discount_value"`
	MaxUsagePerUser       int        `json:"max_usage_per_user"`
}
