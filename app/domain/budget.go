package domain

type Budget struct {
	ID            int `json:"id"`
	TravelId      int `json:"travel_id"`
	Transports    int `json:"transports"`
	Accommodation int `json:"accommodation"`
	Sightseeing   int `json:"sightseeing"`
	Meal          int `json:"meal"`
}
