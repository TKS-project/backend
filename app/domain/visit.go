package domain

import (
	"time"
)

type Visit struct {
	ID        int       `json:"id"`
	BudgetId  int       `json:"budget_id"`
	VisitDay  time.Time `json:"visit_day"`
	Address   string    `json:"address"`
	Charge    int       `json:"charge"`
	Info      string    `json:"info"`
	Name      string    `json:"name"`
	VisitType string    `json:"visit_type"`
}
