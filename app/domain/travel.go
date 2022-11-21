package domain

type Travel struct {
	ID            int `json:"id"`
	UserId        int `json:"user_id"`
	DestinationId int `json:"destination_id"`
	DepartureId   int `json:"departure_id"`
	MaxPerson     int `json:"max_person"`
	MaxDay        int `json:"max_day"`
}
