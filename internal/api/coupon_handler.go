// 📁 internal/api/coupon_handler.go
package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/midhun72/coupon-system/internal/model"
	"github.com/midhun72/coupon-system/internal/service"
)

var couponService *service.CouponService

func InitHandlers(s *service.CouponService) {
	couponService = s
}

func CreateCouponHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var coupon model.Coupon
	if err := json.NewDecoder(r.Body).Decode(&coupon); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	couponService.CreateCoupon(coupon)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Coupon created successfully"})
}

func ValidateCouponHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req model.ValidateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Ensure timestamp is provided or fallback to now
	if req.Timestamp.IsZero() {
		req.Timestamp = time.Now()
	}

	resp := couponService.ValidateCoupon(req)
	status := http.StatusOK
	if !resp.Valid {
		status = http.StatusBadRequest
	}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}
