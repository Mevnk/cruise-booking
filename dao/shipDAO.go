package dao

import (
	"context"
	"database/sql"
	"pis/domain"
)

type ShipsDAO interface {
	CreateShip(ctx context.Context, ship *domain.Ship) error
	UpdateShip(ctx context.Context, ship *domain.Ship) error
	DeleteShip(ctx context.Context, shipID int) error
	GetShipByID(ctx context.Context, shipID int) (*domain.Ship, error)
}

type MySQLShipsDAO struct {
	db *sql.DB
}

func NewMySQLShipsDAO(db *sql.DB) *MySQLShipsDAO {
	return &MySQLShipsDAO{
		db: db,
	}
}

func (dao *MySQLShipsDAO) CreateShip(ctx context.Context, ship *domain.Ship) error {
	query := "INSERT INTO Ships (id, name, passenger_capacity, route, number_of_ports, duration, staff) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := dao.db.ExecContext(ctx, query, ship.Id, ship.Name, ship.PassCap, ship.Route, ship.NofPorts, ship.Duration, ship.Staff)
	return err
}

func (dao *MySQLShipsDAO) UpdateShip(ctx context.Context, ship *domain.Ship) error {
	query := "UPDATE Ships SET name = ?, passenger_capacity = ?, route = ?, number_of_ports = ?, duration = ?, staff = ? WHERE id = ?"
	_, err := dao.db.ExecContext(ctx, query, ship.Name, ship.PassCap, ship.Route, ship.NofPorts, ship.Duration, ship.Staff, ship.Id)
	return err
}

func (dao *MySQLShipsDAO) DeleteShip(ctx context.Context, shipID int) error {
	query := "DELETE FROM Ships WHERE id = ?"
	_, err := dao.db.ExecContext(ctx, query, shipID)
	return err
}

func (dao *MySQLShipsDAO) GetShipByID(ctx context.Context, shipID int) (*domain.Ship, error) {
	query := "SELECT id, name, passenger_capacity, route, number_of_ports, duration, staff FROM Ships WHERE id = ?"
	row := dao.db.QueryRowContext(ctx, query, shipID)

	ship := &domain.Ship{}
	err := row.Scan(&ship.Id, &ship.Name, &ship.PassCap, &ship.Route, &ship.NofPorts, &ship.Duration, &ship.Staff)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Ship not found
		}
		return nil, err
	}

	return ship, nil
}
