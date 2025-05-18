package service

import (
	"github.com/midhun72/coupon-system/internal/model"
	"github.com/midhun72/coupon-system/internal/repository"
	"github.com/midhun72/coupon-system/internal/validation"
)

type CouponService struct {
	repo *repository.CouponRepo
}

func NewCouponService(r *repository.CouponRepo) *CouponService {
	return &CouponService{repo: r}
}

func (s *CouponService) CreateCoupon(coupon model.Coupon) {
	s.repo.Save(coupon)
}

func (s *CouponService) ValidateCoupon(req model.ValidateRequest) model.ValidateResponse {
	coupon, err := s.repo.Get(req.CouponCode)
	if err != nil {
		return model.ValidateResponse{Valid: false, Message: "Coupon not found"}
	}

	resp := validation.ValidateLogic(coupon, req, s.repo)
	if resp.Valid {
		s.repo.IncrementUsage(req.UserID, coupon.CouponCode)
	}
	return resp
}

