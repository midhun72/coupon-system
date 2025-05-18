package main

import (
	"log"
	"net/http"

	"github.com/midhun72/coupon-system/internal/api"
	"github.com/midhun72/coupon-system/internal/repository"
	"github.com/midhun72/coupon-system/internal/service"
)

func main() {
	repo := repository.NewCouponRepo()
	svc := service.NewCouponService(repo)
	api.InitHandlers(svc)

	http.HandleFunc("/admin/coupons", api.CreateCouponHandler)
	http.HandleFunc("/validate-coupon", api.ValidateCouponHandler)

	log.Println("🚀 Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
