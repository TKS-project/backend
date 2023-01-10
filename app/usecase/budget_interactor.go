package usecase

import "github.com/Kantaro0829/clean-architecture-in-go/domain"

type BudgetInteractor struct {
	BudgetRepository BudgetRepository
}

func (interactor *BudgetInteractor) Add(b domain.Budget) (int, error) {
	budgetId, err := interactor.BudgetRepository.Insert(b)
	if err != nil {
		return 0, err
	}
	return budgetId, nil
}

func (interactor *BudgetInteractor) UpdateBudget(b domain.BudgetUpdate, userId int) error {
	err := interactor.BudgetRepository.Update(b, userId)
	if err != nil {
		return err
	}
	return nil
}
