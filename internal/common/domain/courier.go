package domain

// CourierStatus — статус курьера
type CourierStatus string

const (
	CourierStatusFree CourierStatus = "free" // свободен
	CourierStatusBusy CourierStatus = "busy" // занят (выполняет заказ)
)

// Courier — бизнес-сущность курьера
type Courier struct {
	ID              string        `json:"id"`
	Name            string        `json:"name"`
	Phone           string        `json:"phone"`
	Status          CourierStatus `json:"status"`
	CurrentLocation Location      `json:"current_location"`
}
