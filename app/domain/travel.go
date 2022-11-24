package domain

type Travel struct {
	ID            int `json:"id" gorm:"primary_key auto_increment"`
	UserId        int `json:"user_id"`
	DestinationId int `json:"destination_id"`
	DepartureId   int `json:"departure_id"`
	MaxPerson     int `json:"max_person"`
	MaxDay        int `json:"max_day"`
}

/*
{
	"destination_id": 24, "departure_id": 25, "max_person": 1, "max_day": 1
}
*/
