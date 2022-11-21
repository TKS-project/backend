package domain

type Transportation struct {
	ID          int    `json:"id" gorm:"primary_key auto_increment"`
	BudgetId    int    `json:"budget_id"`
	Destination string `json:"destination"`
	Departure   string `json:"departure"`
	Type        string `json:"type"`
	Charge      int    `json:"charge"`
	Order       int    `json:"order"`
	IsWayThere  bool   `json:"is_way_there"`
}