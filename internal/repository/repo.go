package repository

import (
	"sync"
	"errors"
	"github.com/midhun72/coupon-system/internal/model"
)
var ErrCouponNotFound = errors.New("coupon not found")

type CouponRepo struct {
	mu              sync.Mutex
	coupons         map[string]model.Coupon
	userCouponUsage map[string]map[string]int
}

func NewCouponRepo() *CouponRepo {
	return &CouponRepo{
		coupons:         make(map[string]model.Coupon),
		userCouponUsage: make(map[string]map[string]int),
	}
}

func (r *CouponRepo) Save(coupon model.Coupon) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.coupons[coupon.CouponCode] = coupon
}

func (r *CouponRepo) Get(couponCode string) (model.Coupon, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	coupon, ok := r.coupons[couponCode]
	if !ok {
		return model.Coupon{}, ErrCouponNotFound
	}
	return coupon, nil
}


func (r *CouponRepo) HasUserUsed(userID, couponCode string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if userCoupons, exists := r.userCouponUsage[userID]; exists {
		return userCoupons[couponCode] > 0
	}
	return false
}

func (r *CouponRepo) GetUsageCount(userID, couponCode string) int {
	r.mu.Lock()
	defer r.mu.Unlock()
	if userCoupons, exists := r.userCouponUsage[userID]; exists {
		return userCoupons[couponCode]
	}
	return 0
}

func (r *CouponRepo) IncrementUsage(userID, couponCode string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.userCouponUsage[userID]; !exists {
		r.userCouponUsage[userID] = make(map[string]int)
	}
	r.userCouponUsage[userID][couponCode]++
}
