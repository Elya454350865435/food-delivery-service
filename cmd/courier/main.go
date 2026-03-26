package main

import (
	"fmt"
	"log"

	"github.com/Elya454350865435/food-delivery-service/infrastructure/external_services"
	"github.com/Elya454350865435/food-delivery-service/infrastructure/repositories/memory"
	"github.com/Elya454350865435/food-delivery-service/internal/common/domain"
)

func main() {
	fmt.Println("=== Проверка инфраструктурных компонентов ===\n")

	// 1. Проверяем репозиторий
	fmt.Println("1. Тестируем InMemoryOrderRepo...")
	repo := memory.NewInMemoryOrderRepo()

	// Поиск доступных заказов
	orders, err := repo.FindAvailableOrders(10, 0)
	if err != nil {
		log.Printf("Ошибка: %v", err)
	} else {
		fmt.Printf("   Найдено доступных заказов: %d\n", len(orders))
		for _, o := range orders {
			fmt.Printf("   - Заказ %s: %s, доставка %.0f руб.\n", o.ID, o.RestaurantName, o.DeliveryFee)
		}
	}

	// Поиск заказа по ID
	order, err := repo.FindByID("order_001")
	if err != nil {
		log.Printf("Ошибка: %v", err)
	} else {
		fmt.Printf("   Заказ order_001 найден: статус %s\n", order.Status)
	}

	// Обновление статуса
	err = repo.UpdateOrderStatus("order_001", domain.OrderStatusAssigned, "courier_test")
	if err != nil {
		fmt.Printf("   Ошибка обновления: %v\n", err)
	} else {
		fmt.Println("   Статус заказа order_001 обновлен на 'assigned'")
	}

	fmt.Println("\n2. Тестируем MockNotificationService...")
	notifSvc := external_services.NewMockNotificationService()
	notifSvc.SendOrderAcceptedNotification("courier_001", "order_001")
	notifSvc.SendOrderRejectedNotification("courier_001", "order_002", "слишком далеко")

	fmt.Println("\n3. Тестируем MockMapService...")
	mapSvc := external_services.NewMockMapService()

	courierLoc := domain.Location{Lat: 59.940, Lon: 30.350}
	restaurantLoc := domain.Location{Lat: 59.934, Lon: 30.335}
	deliveryLoc := domain.Location{Lat: 59.931, Lon: 30.360}

	route, err := mapSvc.GetDeliveryRoute(courierLoc, restaurantLoc, deliveryLoc)
	if err != nil {
		log.Printf("Ошибка: %v", err)
	} else {
		fmt.Printf("   Маршрут: %.2f км, %d мин\n", route.TotalDistanceKm, route.TotalEstimatedMin)
		for i, leg := range route.Legs {
			fmt.Printf("   Участок %d: %s, %.2f км, %d мин\n", i+1, leg.Type, leg.DistanceKm, leg.EstimatedMinutes)
		}
	}

	fmt.Println("\n=== Проверка завершена ===")
}
