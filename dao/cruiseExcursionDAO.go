package dao

import (
	"context"
	"database/sql"
	"fmt"
	"pis/pkg/customErrors"

	"github.com/google/uuid"
)

type ExcursionCruiseDAO interface {
	AddExcursionToCruise(ctx context.Context, excursionID, cruiseID uuid.UUID) error
	RemoveExcursionFromCruise(ctx context.Context, excursionID, cruiseID uuid.UUID) error
	GetExcursionsByCruise(ctx context.Context, cruiseID uuid.UUID) ([]uuid.UUID, error)
	CheckBind(ctx context.Context, cruiseId uuid.UUID, excusrionId uuid.UUID) (result bool, err error)
	RemoveExcursion(ctx context.Context, excursionID uuid.UUID) error
	RemoveCruise(ctx context.Context, excursionID uuid.UUID) error
}

type excursionCruiseDAO struct {
	db *sql.DB
}

func NewExcursionCruiseDAO(db *sql.DB) ExcursionCruiseDAO {
	return &excursionCruiseDAO{
		db: db,
	}
}

func (dao *excursionCruiseDAO) AddExcursionToCruise(ctx context.Context, excursionID, cruiseID uuid.UUID) error {
	query := "INSERT INTO excursion_cruise (excursion_id, cruise_id) VALUES (?, ?)"

	_, err := dao.db.ExecContext(ctx, query, excursionID.String(), cruiseID.String())
	if err != nil {
		return err
	}

	return nil
}

func (dao *excursionCruiseDAO) RemoveExcursionFromCruise(ctx context.Context, excursionID, cruiseID uuid.UUID) error {
	query := "DELETE FROM excursion_cruise WHERE excursion_id = ? AND cruise_id = ?"

	_, err := dao.db.ExecContext(ctx, query, excursionID.String(), cruiseID.String())
	if err != nil {
		return err
	}

	return nil
}

func (dao *excursionCruiseDAO) RemoveExcursion(ctx context.Context, excursionID uuid.UUID) error {
	query := "DELETE FROM excursion_cruise WHERE excursion_id = ?"

	_, err := dao.db.ExecContext(ctx, query, excursionID.String())
	if err != nil {
		return err
	}

	return nil
}

func (dao *excursionCruiseDAO) RemoveCruise(ctx context.Context, cruiseId uuid.UUID) error {
	query := "DELETE FROM excursion_cruise WHERE cruise_id = ?"

	_, err := dao.db.ExecContext(ctx, query, cruiseId.String())
	if err != nil {
		return err
	}

	return nil
}

func (dao *excursionCruiseDAO) GetExcursionsByCruise(ctx context.Context, cruiseID uuid.UUID) ([]uuid.UUID, error) {
	query := "SELECT excursion_id FROM excursion_cruise WHERE cruise_id = ?"

	rows, err := dao.db.QueryContext(ctx, query, cruiseID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var excursionIDs []uuid.UUID

	for rows.Next() {
		var excursionID string
		err := rows.Scan(&excursionID)
		if err != nil {
			return nil, customErrors.NewScanError(customErrors.Excursion, customErrors.Scan)
		}

		uuidID, err := uuid.Parse(excursionID)
		if err != nil {
			return nil, customErrors.NewScanError(customErrors.Excursion, customErrors.Parse)
		}

		excursionIDs = append(excursionIDs, uuidID)
	}

	if err := rows.Err(); err != nil {
		return nil, customErrors.NewScanError(customErrors.Excursion, customErrors.Iteration)
	}

	return excursionIDs, nil
}

func (dao *excursionCruiseDAO) CheckBind(ctx context.Context, cruiseId uuid.UUID, excusrionId uuid.UUID) (result bool, err error) {
	query := "SELECT COUNT(*) FROM excursion_cruise WHERE excursion_id = ? AND cruise_id = ?"

	var count int
	err = dao.db.QueryRowContext(ctx, query, excusrionId, cruiseId).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("failed to check if excursion is bound to cruise: %v", err)
	}

	return count > 0, nil
}
