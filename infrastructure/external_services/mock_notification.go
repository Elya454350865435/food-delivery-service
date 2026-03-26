package external_services

import (
	"log"
)

// MockNotificationService — мок-реализация сервиса уведомлений
type MockNotificationService struct{}

// NewMockNotificationService — создает новый мок
func NewMockNotificationService() *MockNotificationService {
	return &MockNotificationService{}
}

// SendOrderAcceptedNotification — уведомление о принятии заказа
func (m *MockNotificationService) SendOrderAcceptedNotification(courierID, orderID string) error {
	log.Printf("[MOCK NOTIFICATION] Courier %s accepted order %s", courierID, orderID)
	return nil
}

// SendOrderRejectedNotification — уведомление об отказе от заказа
func (m *MockNotificationService) SendOrderRejectedNotification(courierID, orderID, reason string) error {
	log.Printf("[MOCK NOTIFICATION] Courier %s rejected order %s. Reason: %s", courierID, orderID, reason)
	return nil
}

// SendOrderCompletedNotification — уведомление о завершении заказа
func (m *MockNotificationService) SendOrderCompletedNotification(userID, orderID string) error {
	log.Printf("[MOCK NOTIFICATION] User %s: order %s completed", userID, orderID)
	return nil
}

// SendNewOrderNotification — уведомление о новом заказе
func (m *MockNotificationService) SendNewOrderNotification(orderID string) error {
	log.Printf("[MOCK NOTIFICATION] New order available: %s", orderID)
	return nil
}
