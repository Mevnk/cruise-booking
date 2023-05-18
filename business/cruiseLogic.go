package business

import (
	"context"
	"github.com/google/uuid"
	"pis/dao"
	"pis/domain"
	"pis/pkg/customErrors"
	"time"
)

type CruiseLogic interface {
	CreateCruise(ctx context.Context, params CruiseParams) (cruiseId uuid.UUID, err error)
	GetCruiseData(ctx context.Context, id uuid.UUID) (cruise *domain.Cruise, err error)
	DeleteCruise(ctx context.Context, id uuid.UUID) (err error)
	AddExcursion(ctx context.Context, cruiseId uuid.UUID, excursionId uuid.UUID) (err error)
	DeleteExcursion(ctx context.Context, cruiseId uuid.UUID, excursionId uuid.UUID) (err error)
	GetExcursions(ctx context.Context, cruiseId uuid.UUID) (excursions []domain.Excursion, err error)
}

type cruiseLogic struct {
	cruiseDAO          dao.CruisesDAO
	excursionDAO       dao.ExcursionsDAO
	cruiseExcursionDAO dao.ExcursionCruiseDAO
}

func (cl cruiseLogic) CreateCruise(ctx context.Context, params CruiseParams) (cruiseId uuid.UUID, err error) {
	cruise := domain.NewCruise(params.ShipID, params.DepartureDate, params.Price)

	if err = cl.cruiseDAO.CreateCruise(ctx, cruise); err != nil {
		return uuid.Nil, customErrors.NewCustomError(customErrors.Cruise, customErrors.Creation)
	}

	return cruise.Id, err
}

func (cl cruiseLogic) GetCruiseData(ctx context.Context, uuid uuid.UUID) (cruise *domain.Cruise, err error) {
	cruise, err = cl.cruiseDAO.GetCruiseByID(ctx, uuid)
	if err != nil {
		return nil, err
	}
	if cruise == nil {
		return nil, customErrors.NewCustomError(customErrors.Cruise, customErrors.NotFound)
	}

	return cruise, nil
}

func (cl cruiseLogic) DeleteCruise(ctx context.Context, id uuid.UUID) (err error) {
	err = cl.cruiseDAO.DeleteCruise(ctx, id)
	if err != nil {
		return customErrors.NewCustomError(customErrors.Cruise, customErrors.Deletion)
	}
	return nil
}

func (cl cruiseLogic) AddExcursion(ctx context.Context, cruiseId uuid.UUID, excursionId uuid.UUID) (err error) {
	checkResult, err := cl.cruiseExcursionDAO.CheckBind(ctx, cruiseId, excursionId)
	if checkResult {
		return nil
	} else {
		err = cl.cruiseExcursionDAO.AddExcursionToCruise(ctx, excursionId, cruiseId)
		if err != nil {
			return customErrors.NewCustomError(customErrors.CruiseExcursion, customErrors.Creation)
		}
	}
	return nil
}

func (cl cruiseLogic) DeleteExcursion(ctx context.Context, cruiseId uuid.UUID, excursionId uuid.UUID) (err error) {
	checkResult, err := cl.cruiseExcursionDAO.CheckBind(ctx, cruiseId, excursionId)
	if !checkResult {
		return nil
	} else {
		err = cl.cruiseExcursionDAO.RemoveExcursionFromCruise(ctx, excursionId, cruiseId)
		if err != nil {
			return customErrors.NewCustomError(customErrors.CruiseExcursion, customErrors.Deletion)
		}
	}
	return nil
}

func (cl cruiseLogic) GetExcursions(ctx context.Context, cruiseId uuid.UUID) (excursions []domain.Excursion, err error) {
	var excursionIDs []uuid.UUID
	excursionIDs, err = cl.cruiseExcursionDAO.GetExcursionsByCruise(ctx, cruiseId)
	if err != nil {
		return nil, customErrors.NewScanError(customErrors.CruiseExcursion, customErrors.Iteration)
	}

	for id := range excursionIDs {
		excursion, err := cl.excursionDAO.GetExcursionByID(ctx, id)
		if err != nil {
			continue
		}
		excursions = append(excursions, *excursion)
	}
	return excursions, nil
}

type CruiseParams struct {
	ShipID        uuid.UUID
	DepartureDate time.Time
	Price         int
}
