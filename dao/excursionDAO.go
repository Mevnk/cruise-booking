package dao

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"pis/domain"
)

type ExcursionsDAO interface {
	CreateExcursion(ctx context.Context, excursion *domain.Excursion) error
	UpdateExcursion(ctx context.Context, excursion *domain.Excursion) error
	DeleteExcursion(ctx context.Context, excursionID uuid.UUID) error
	GetExcursionByID(ctx context.Context, excursionID uuid.UUID) (*domain.Excursion, error)
}

type MySQLExcursionsDAO struct {
	db *sql.DB
}

func NewMySQLExcursionsDAO(db *sql.DB) *MySQLExcursionsDAO {
	return &MySQLExcursionsDAO{
		db: db,
	}
}

func (dao *MySQLExcursionsDAO) CreateExcursion(ctx context.Context, excursion *domain.Excursion) error {
	query := "INSERT INTO Excursions (id, name, description, price) VALUES (?, ?, ?, ?)"
	_, err := dao.db.ExecContext(ctx, query, excursion.Id, excursion.Name, excursion.Description, excursion.Price)
	return err
}

func (dao *MySQLExcursionsDAO) UpdateExcursion(ctx context.Context, excursion *domain.Excursion) error {
	query := "UPDATE Excursions SET name = ?, description = ?, price = ? WHERE id = ?"
	_, err := dao.db.ExecContext(ctx, query, excursion.Name, excursion.Description, excursion.Price, excursion.Id)
	return err
}

func (dao *MySQLExcursionsDAO) DeleteExcursion(ctx context.Context, excursionID uuid.UUID) error {
	query := "DELETE FROM Excursions WHERE id = ?"
	_, err := dao.db.ExecContext(ctx, query, excursionID)
	return err
}

func (dao *MySQLExcursionsDAO) GetExcursionByID(ctx context.Context, excursionID uuid.UUID) (*domain.Excursion, error) {
	query := "SELECT id, name, description, price FROM Excursions WHERE id = ?"
	row := dao.db.QueryRowContext(ctx, query, excursionID)

	excursion := &domain.Excursion{}
	err := row.Scan(&excursion.Id, &excursion.Name, &excursion.Description, &excursion.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Excursion not found
		}
		return nil, err
	}

	return excursion, nil
}
