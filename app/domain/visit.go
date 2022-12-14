package domain

import "time"

type Visit struct {
	ID        int       `json:"id" gorm:"primary_key auto_increment"`
	BudgetId  int       `json:"budget_id"`
	VisitDay  time.Time `json:"visit_day"`
	Address   string    `json:"address"`
	Charge    int       `json:"charge"`
	Info      string    `json:"info"`
	Name      string    `json:"name"`
	VisitType string    `json:"visit_type"`
}

type VisitJson struct {
	ID        int    `json:"id"`
	BudgetId  int    `json:"budget_id"`
	VisitDay  string `json:"visit_day"`
	Address   string `json:"address"`
	Charge    int    `json:"charge"`
	Info      string `json:"info"`
	Name      string `json:"name"`
	VisitType string `json:"visit_type"`
}
