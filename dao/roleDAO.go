package dao

import (
	"context"
	"database/sql"
	"pis/domain"
)

type RolesDAO interface {
	CreateRole(ctx context.Context, role *domain.Role) error
	UpdateRole(ctx context.Context, role *domain.Role) error
	DeleteRole(ctx context.Context, roleID int) error
	GetRoleByID(ctx context.Context, roleID int) (*domain.Role, error)
}

type MySQLRolesDAO struct {
	db *sql.DB
}

func NewMySQLRolesDAO(db *sql.DB) *MySQLRolesDAO {
	return &MySQLRolesDAO{
		db: db,
	}
}

func (dao *MySQLRolesDAO) CreateRole(ctx context.Context, role *domain.Role) error {
	query := "INSERT INTO Roles (id, name) VALUES (?, ?)"
	_, err := dao.db.ExecContext(ctx, query, role.Id, role.Name)
	return err
}

func (dao *MySQLRolesDAO) UpdateRole(ctx context.Context, role *domain.Role) error {
	query := "UPDATE Roles SET name = ? WHERE id = ?"
	_, err := dao.db.ExecContext(ctx, query, role.Name, role.Id)
	return err
}

func (dao *MySQLRolesDAO) DeleteRole(ctx context.Context, roleID int) error {
	query := "DELETE FROM Roles WHERE id = ?"
	_, err := dao.db.ExecContext(ctx, query, roleID)
	return err
}

func (dao *MySQLRolesDAO) GetRoleByID(ctx context.Context, roleID int) (*domain.Role, error) {
	query := "SELECT id, name FROM Roles WHERE id = ?"
	row := dao.db.QueryRowContext(ctx, query, roleID)

	role := &domain.Role{}
	err := row.Scan(&role.Id, &role.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Role not found
		}
		return nil, err
	}

	return role, nil
}

func (dao *MySQLRolesDAO) GetRoleByName(ctx context.Context, roleName string) (*domain.Role, error) {
	query := "SELECT id, name FROM Roles WHERE id = ?"
	row := dao.db.QueryRowContext(ctx, query, roleName)

	role := &domain.Role{}
	err := row.Scan(&role.Id, &role.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Role not found
		}
		return nil, err
	}

	return role, nil
}
