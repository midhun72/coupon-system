package main

import (
	"log"
	"net/http"

	"coupon-system/internal/api"
	"coupon-system/internal/repository"
	"coupon-system/internal/service"
)

func main() {
	repo := repository.NewCouponRepo()
	svc := service.NewCouponService(repo)
	api.InitHandlers(svc)

	http.HandleFunc("/admin/coupons", api.CreateCouponHandler)
	http.HandleFunc("/validate-coupon", api.ValidateCouponHandler)
	http.HandleFunc("/admin/coupons/list", api.GetAllCouponsHandler)

	log.Println("🚀 Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
