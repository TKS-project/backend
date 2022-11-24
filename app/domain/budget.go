package domain

type Budget struct {
	ID              int    `json:"id" gorm:"primary_key auto_increment"`
	TravelId        int    `json:"travel_id"`
	Transportations uint32 `json:"transports"`
	Accommodation   uint32 `json:"accommodation"`
	Sightseeing     uint32 `json:"sightseeing"`
	Meal            uint32 `json:"meal"`
	Sum             uint32 `json:"sum"`
}

/*
{
"travel_id": 3,"transports": 0, "accommodation": 0, "sightseeing": 0, "meal": 0, "sum": 50000
}
*/
