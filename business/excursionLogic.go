package business

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"pis/dao"
	"pis/domain"
	"pis/pkg/customErrors"
	"pis/pkg/transactionManager"
)

type ExcursionLogic interface {
	CreateExcursion(ctx context.Context, params ExcursionParams) (id uuid.UUID, err error)
	UpdateExcursion(ctx context.Context, params ExcursionParams) (err error)
	DeleteExcursion(ctx context.Context, excursionId uuid.UUID) (err error)
}

type excursionLogic struct {
	excursionDAO dao.ExcursionsDAO
	bindDAO      dao.ExcursionCruiseDAO
	trxManager   transactionManager.TransactionManager
}

func (el excursionLogic) CreateExcursion(ctx context.Context, params ExcursionParams) (id uuid.UUID, err error) {
	excursion := domain.NewExcursion(params.Name, params.Description, params.Price)

	err = el.excursionDAO.CreateExcursion(ctx, excursion)
	if err != nil {
		return uuid.Nil, customErrors.NewCustomError(customErrors.Excursion, customErrors.Creation)
	}

	return excursion.Id, err
}

func (el excursionLogic) UpdateExcursion(ctx context.Context, params ExcursionParams) (err error) {
	excursion := domain.NewExcursion(params.Name, params.Description, params.Price)

	err = el.excursionDAO.UpdateExcursion(ctx, excursion)
	if err != nil {
		return customErrors.NewCustomError(customErrors.Excursion, customErrors.Creation)
	}

	return err
}

func (el excursionLogic) DeleteExcursion(ctx context.Context, excursionId uuid.UUID) (err error) {
	err = el.trxManager.RunTransaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		err := el.excursionDAO.DeleteExcursion(ctx, excursionId)
		if err != nil {
			return customErrors.NewCustomError(customErrors.Excursion, customErrors.Deletion)
		}
		err = el.bindDAO.RemoveExcursion(ctx, excursionId)
		if err != nil {
			return customErrors.NewCustomError(customErrors.CruiseExcursion, customErrors.Deletion)
		}
		return nil
	})
	if err != nil {
		return customErrors.NewCustomError(customErrors.Excursion, customErrors.Deletion)
	}
	return nil
}

type ExcursionParams struct {
	Name        string
	Description string
	Price       int
}
