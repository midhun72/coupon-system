package validation

import (
	"time"

	"github.com/midhun72/coupon-system/internal/model"
	"github.com/midhun72/coupon-system/internal/repository"
)

func ValidateLogic(coupon model.Coupon, req model.ValidateRequest, repo *repository.CouponRepo) model.ValidateResponse {
	now := req.Timestamp

	if now.After(coupon.ExpiryDate) {
		return model.ValidateResponse{Valid: false, Message: "Coupon expired"}
	}

	if coupon.MinOrderValue > 0 && req.OrderValue < coupon.MinOrderValue {
		return model.ValidateResponse{Valid: false, Message: "Minimum order value not met"}
	}

	if len(coupon.ApplicableMedicineIDs) > 0 && !anyOverlap(req.MedicineIDs, coupon.ApplicableMedicineIDs) {
		return model.ValidateResponse{Valid: false, Message: "Coupon not applicable to selected medicines"}
	}

	if len(coupon.ApplicableCategories) > 0 && !anyOverlap(req.CategoryIDs, coupon.ApplicableCategories) {
		return model.ValidateResponse{Valid: false, Message: "Coupon not applicable to selected categories"}
	}

	if coupon.ValidTimeWindow.Start != "" && coupon.ValidTimeWindow.End != "" {
		start, _ := time.Parse("15:04", coupon.ValidTimeWindow.Start)
		end, _ := time.Parse("15:04", coupon.ValidTimeWindow.End)
		hourMinute := now.Format("15:04")
		current, _ := time.Parse("15:04", hourMinute)

		if current.Before(start) || current.After(end) {
			return model.ValidateResponse{Valid: false, Message: "Coupon not valid at this time"}
		}
	}

	if coupon.UsageType == "one_time" {
		if repo.HasUserUsed(req.UserID, coupon.CouponCode) {
			return model.ValidateResponse{Valid: false, Message: "Coupon already used"}
		}
	}

	if coupon.UsageType == "multi_use" {
		if repo.GetUsageCount(req.UserID, coupon.CouponCode) >= coupon.MaxUsagePerUser {
			return model.ValidateResponse{Valid: false, Message: "Coupon usage limit reached"}
		}
	}

	// Calculate discount
	var discount float64
	if coupon.DiscountType == "percentage" {
		discount = req.OrderValue * coupon.DiscountValue / 100
	} else if coupon.DiscountType == "flat" {
		discount = coupon.DiscountValue
	}
	finalPrice := req.OrderValue - discount
	if finalPrice < 0 {
		finalPrice = 0
	}

	return model.ValidateResponse{
		Valid:          true,
		DiscountAmount: discount,
		FinalPrice:     finalPrice,
		Message:        "Coupon applied successfully",
	}
}

func anyOverlap(a, b []string) bool {
	set := make(map[string]struct{})
	for _, item := range b {
		set[item] = struct{}{}
	}
	for _, item := range a {
		if _, exists := set[item]; exists {
			return true
		}
	}
	return false
}
