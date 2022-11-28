package domain

type Transportation struct {
	ID          int    `json:"id" gorm:"primary_key auto_increment"`
	BudgetId    int    `json:"budget_id"`
	Destination string `json:"destination"`
	Departure   string `json:"departure"`
	Type        string `json:"type"`
	Charge      uint16 `json:"charge"`
	OrderNo     uint16 `json:"order_no"`
	WayThere    bool   `json:"way_there"`
}
