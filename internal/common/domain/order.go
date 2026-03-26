package domain

import "time"

// OrderStatus — статус заказа
type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"   // создан, ожидает оплаты/подтверждения
	OrderStatusAvailable OrderStatus = "available" // оплачен, ждет курьера
	OrderStatusAssigned  OrderStatus = "assigned"  // курьер назначен
	OrderStatusCompleted OrderStatus = "completed" // доставлен
	OrderStatusCancelled OrderStatus = "cancelled" // отменен
)

// Order — основная бизнес-сущность заказа
type Order struct {
	ID                 string      `json:"id"`
	Status             OrderStatus `json:"status"`
	UserID             string      `json:"user_id"`
	RestaurantID       string      `json:"restaurant_id"`
	RestaurantName     string      `json:"restaurant_name"`
	RestaurantAddress  string      `json:"restaurant_address"`
	RestaurantLocation Location    `json:"restaurant_location"`
	DeliveryAddress    string      `json:"delivery_address"`
	DeliveryLocation   Location    `json:"delivery_location"`
	ItemsCount         int         `json:"items_count"`
	DeliveryFee        float64     `json:"delivery_fee"`
	TotalAmount        float64     `json:"total_amount"`
	AssignedCourierID  string      `json:"assigned_courier_id,omitempty"`
	CreatedAt          time.Time   `json:"created_at"`
	CompletedAt        *time.Time  `json:"completed_at,omitempty"`
}
