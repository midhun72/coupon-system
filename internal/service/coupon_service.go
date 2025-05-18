package service

import (
	"fmt"
	"coupon-system/internal/model"
	"coupon-system/internal/repository"
	"coupon-system/internal/validation"
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
	fmt.Println(req)
	fmt.Println(req.CouponCode)
	fmt.Println(s.repo)
	fmt.Println(err)
	if err != nil {
		return model.ValidateResponse{Valid: false, Message: "Coupon not found"}
	}

	resp := validation.ValidateLogic(coupon, req, s.repo)
	if resp.Valid {
		s.repo.IncrementUsage(req.UserID, coupon.CouponCode)
	}
	return resp
}

func (s *CouponService) GetAllCoupons() []model.Coupon {
	return s.repo.GetAllCoupons()
}


