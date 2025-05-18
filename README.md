To run this go
 Command : go run cmd/server/main.go

 This will start the server

Then

Then to save coupons

url : http://localhost:8080/admin/coupons
Request data 
--------------
{
  "coupon_code": "MEDSAVE50",
  "expiry_date": "2025-05-18T23:59:59Z",
  "usage_type": "multi_use",
  "applicable_medicine_ids": ["MED1", "MED2"],
  "applicable_categories": ["painkiller","antibiotic"],
  "min_order_value": 500,
  "valid_time_window": {
    "start": "09:00",
    "end": "21:00"
  },
  "terms_and_conditions": "Use only once per day",
  "discount_type": "percentage",
  "discount_value": 20,
  "max_usage_per_user": 5
}

To validate coupen for a user
url : http://localhost:8080/validate-coupon
Request
------------
{
  "coupon_code": "MEDSAVE50",
  "user_id": "user123",
  "order_value": 600,
  "medicine_ids": ["MED2"],
  "category_ids": ["antibiotic"],
  "timestamp": "2025-05-18T14:30:00Z"
}

To Get all coupon details
--------------------------
Url (GET) : http://localhost:8080/admin/coupons/list
