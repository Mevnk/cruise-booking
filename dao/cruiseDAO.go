package dao

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"pis/domain"
)

type CruisesDAO interface {
	CreateCruise(ctx context.Context, cruise *domain.Cruise) error
	UpdateCruise(ctx context.Context, cruise *domain.Cruise) error
	DeleteCruise(ctx context.Context, cruiseID uuid.UUID) error
	GetCruiseByID(ctx context.Context, cruiseID uuid.UUID) (*domain.Cruise, error)
}

type MySQLCruisesDAO struct {
	db *sql.DB
}

func NewMySQLCruisesDAO(db *sql.DB) *MySQLCruisesDAO {
	return &MySQLCruisesDAO{
		db: db,
	}
}

func (dao *MySQLCruisesDAO) CreateCruise(ctx context.Context, cruise *domain.Cruise) error {
	query := "INSERT INTO Cruises (id, ship_id, departure_date, price, excursions) VALUES (?, ?, ?, ?, ?)"
	_, err := dao.db.ExecContext(ctx, query, cruise.Id, cruise.ShipID, cruise.DepartureDate, cruise.Price)
	return err
}

func (dao *MySQLCruisesDAO) UpdateCruise(ctx context.Context, cruise *domain.Cruise) error {
	query := "UPDATE Cruises SET ship_id = ?, departure_date = ?, price = ?, excursions = ? WHERE id = ?"
	_, err := dao.db.ExecContext(ctx, query, cruise.ShipID, cruise.DepartureDate, cruise.Price, cruise.Id)
	return err
}

func (dao *MySQLCruisesDAO) DeleteCruise(ctx context.Context, cruiseID uuid.UUID) error {
	query := "DELETE FROM Cruises WHERE id = ?"
	_, err := dao.db.ExecContext(ctx, query, cruiseID)
	return err
}

func (dao *MySQLCruisesDAO) GetCruiseByID(ctx context.Context, cruiseID uuid.UUID) (*domain.Cruise, error) {
	query := "SELECT id, ship_id, departure_date, price, excursions FROM Cruises WHERE id = ?"
	row := dao.db.QueryRowContext(ctx, query, cruiseID)

	cruise := &domain.Cruise{}
	err := row.Scan(&cruise.Id, &cruise.ShipID, &cruise.DepartureDate, &cruise.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Cruise not found
		}
		return nil, err
	}

	return cruise, nil
}
