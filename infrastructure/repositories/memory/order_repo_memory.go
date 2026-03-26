package memory

import (
	"errors"
	"sync"
	"time"

	"github.com/Elya454350865435/food-delivery-service/internal/common/domain"
)

// InMemoryOrderRepo — реализация репозитория в памяти
type InMemoryOrderRepo struct {
	orders map[string]domain.Order
	mu     sync.RWMutex
}

// NewInMemoryOrderRepo — создает новый репозиторий
func NewInMemoryOrderRepo() *InMemoryOrderRepo {
	repo := &InMemoryOrderRepo{
		orders: make(map[string]domain.Order),
	}
	// Добавляем тестовые данные
	repo.addTestOrders()
	return repo
}

// FindAvailableOrders — найти доступные заказы
func (r *InMemoryOrderRepo) FindAvailableOrders(limit, offset int) ([]domain.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []domain.Order
	for _, order := range r.orders {
		if order.Status == domain.OrderStatusAvailable {
			result = append(result, order)
		}
	}

	// Простая пагинация
	start := offset
	end := offset + limit
	if start > len(result) {
		return []domain.Order{}, nil
	}
	if end > len(result) {
		end = len(result)
	}

	return result[start:end], nil
}

// FindByID — найти заказ по ID
func (r *InMemoryOrderRepo) FindByID(orderID string) (domain.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	order, exists := r.orders[orderID]
	if !exists {
		return domain.Order{}, errors.New("order not found")
	}
	return order, nil
}

// UpdateOrderStatus — обновить статус заказа
func (r *InMemoryOrderRepo) UpdateOrderStatus(orderID string, newStatus domain.OrderStatus, courierID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	order, exists := r.orders[orderID]
	if !exists {
		return errors.New("order not found")
	}
	if order.Status != domain.OrderStatusAvailable {
		return errors.New("order is not available")
	}

	order.Status = newStatus
	order.AssignedCourierID = courierID
	r.orders[orderID] = order
	return nil
}

// CompleteOrder — завершить заказ
func (r *InMemoryOrderRepo) CompleteOrder(orderID string, signature, photoProof string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	order, exists := r.orders[orderID]
	if !exists {
		return errors.New("order not found")
	}
	if order.Status != domain.OrderStatusAssigned {
		return errors.New("order is not assigned")
	}

	order.Status = domain.OrderStatusCompleted
	now := time.Now()
	order.CompletedAt = &now
	r.orders[orderID] = order
	return nil
}

// CreateOrder — создать новый заказ
func (r *InMemoryOrderRepo) CreateOrder(order domain.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.orders[order.ID]; exists {
		return errors.New("order already exists")
	}
	r.orders[order.ID] = order
	return nil
}

// UpdateOrderCourier — назначить курьера на заказ
func (r *InMemoryOrderRepo) UpdateOrderCourier(orderID, courierID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	order, exists := r.orders[orderID]
	if !exists {
		return errors.New("order not found")
	}

	order.AssignedCourierID = courierID
	r.orders[orderID] = order
	return nil
}

// addTestOrders — добавляет тестовые заказы для проверки
func (r *InMemoryOrderRepo) addTestOrders() {
	now := time.Now()

	// Тестовый заказ 1
	r.orders["order_001"] = domain.Order{
		ID:                 "order_001",
		Status:             domain.OrderStatusAvailable,
		UserID:             "user_001",
		RestaurantID:       "rest_001",
		RestaurantName:     "Суши Wok",
		RestaurantAddress:  "Арсенальная набережная, 7",
		RestaurantLocation: domain.Location{Lat: 59.934, Lon: 30.335},
		DeliveryAddress:    "Торжковская, д. 15, кв. 240",
		DeliveryLocation:   domain.Location{Lat: 59.931, Lon: 30.360},
		ItemsCount:         3,
		DeliveryFee:        150.0,
		TotalAmount:        1250.0,
		CreatedAt:          now,
	}

	// Тестовый заказ 2
	r.orders["order_002"] = domain.Order{
		ID:                 "order_002",
		Status:             domain.OrderStatusAvailable,
		UserID:             "user_002",
		RestaurantID:       "rest_002",
		RestaurantName:     "Бургерная №1",
		RestaurantAddress:  "Невский проспект, 100",
		RestaurantLocation: domain.Location{Lat: 59.931, Lon: 30.360},
		DeliveryAddress:    "Лиговский проспект, 50",
		DeliveryLocation:   domain.Location{Lat: 59.920, Lon: 30.355},
		ItemsCount:         2,
		DeliveryFee:        100.0,
		TotalAmount:        850.0,
		CreatedAt:          now,
	}

	// Тестовый заказ 3 (уже назначен, не доступен)
	r.orders["order_003"] = domain.Order{
		ID:                 "order_003",
		Status:             domain.OrderStatusAssigned,
		UserID:             "user_003",
		RestaurantID:       "rest_001",
		RestaurantName:     "Суши Wok",
		RestaurantAddress:  "Арсенальная набережная, 7",
		RestaurantLocation: domain.Location{Lat: 59.934, Lon: 30.335},
		DeliveryAddress:    "Пискаревский проспект, 25",
		DeliveryLocation:   domain.Location{Lat: 59.950, Lon: 30.380},
		ItemsCount:         1,
		DeliveryFee:        150.0,
		TotalAmount:        450.0,
		AssignedCourierID:  "courier_001",
		CreatedAt:          now,
	}
}
