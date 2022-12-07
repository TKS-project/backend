package domain

//DB 用
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

//json response 用
type TransportsJson struct {
	MaxCharge    string      `json:"max_charge"`
	LeftToArrive string      `json:"left_to_arrive"`
	Transports   []Transport `json:"transports"`
}

//json response 用
type Transport struct {
	Type     string `json:"type"`
	Time     string `json:"time"`
	Name     string `json:"name"`
	Url      string `json:"url"`
	Plathome string `json:"plathome"`
}

type TransportsRequest struct {
	OrvStation string `json:"orv"`
	DnvStation string `json:"dnv"`
	Year       string `json:"year"`
	Month      string `json:"month"`
	Day        string `json:"day"`
	Hour       string `json:"hour"`
	Minute     string `json:"minute"`
}
