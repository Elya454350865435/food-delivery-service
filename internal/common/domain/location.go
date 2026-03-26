package domain

// Location — географическая координата
type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// Route — маршрут доставки
type Route struct {
	OrderID           string     `json:"order_id"`
	TotalDistanceKm   float64    `json:"total_distance_km"`
	TotalEstimatedMin int        `json:"total_estimated_minutes"`
	Legs              []RouteLeg `json:"legs"`
}

// RouteLeg — один участок маршрута
type RouteLeg struct {
	Type             string   `json:"type"` // "to_restaurant" или "to_client"
	From             Location `json:"from"`
	To               Location `json:"to"`
	DistanceKm       float64  `json:"distance_km"`
	EstimatedMinutes int      `json:"estimated_minutes"`
	Polyline         string   `json:"polyline"` // закодированная полилиния для отрисовки на карте
}
