package external_services

import (
	"math"

	"github.com/Elya454350865435/food-delivery-service/internal/common/domain"
)

// MockMapService — мок-реализация картографического сервиса
type MockMapService struct{}

// NewMockMapService — создает новый мок
func NewMockMapService() *MockMapService {
	return &MockMapService{}
}

// GetDeliveryRoute — построить маршрут (мок-версия)
func (m *MockMapService) GetDeliveryRoute(courierLoc, restaurantLoc, deliveryLoc domain.Location) (domain.Route, error) {
	// Рассчитываем расстояния по прямой (для мока)
	distToRestaurant := m.CalculateDistance(courierLoc, restaurantLoc)
	distToClient := m.CalculateDistance(restaurantLoc, deliveryLoc)

	// Примерное время: 2 минуты на километр
	timeToRestaurant := int(distToRestaurant * 2)
	timeToClient := int(distToClient * 2)

	route := domain.Route{
		OrderID:           "test_order",
		TotalDistanceKm:   distToRestaurant + distToClient,
		TotalEstimatedMin: timeToRestaurant + timeToClient,
		Legs: []domain.RouteLeg{
			{
				Type:             "to_restaurant",
				From:             courierLoc,
				To:               restaurantLoc,
				DistanceKm:       distToRestaurant,
				EstimatedMinutes: timeToRestaurant,
				Polyline:         "mock_polyline_to_restaurant",
			},
			{
				Type:             "to_client",
				From:             restaurantLoc,
				To:               deliveryLoc,
				DistanceKm:       distToClient,
				EstimatedMinutes: timeToClient,
				Polyline:         "mock_polyline_to_client",
			},
		},
	}

	return route, nil
}

// CalculateDistance — рассчитать расстояние между двумя точками (формула гаверсинусов)
func (m *MockMapService) CalculateDistance(from, to domain.Location) float64 {
	const earthRadiusKm = 6371

	lat1 := from.Lat * math.Pi / 180
	lat2 := to.Lat * math.Pi / 180
	dlat := (to.Lat - from.Lat) * math.Pi / 180
	dlon := (to.Lon - from.Lon) * math.Pi / 180

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dlon/2)*math.Sin(dlon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadiusKm * c
}
