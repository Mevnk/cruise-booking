package dao

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"pis/domain"
)

type UsersDAO interface {
	CreateUser(ctx context.Context, user *domain.User) error
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, userID int) error
	GetUserByID(ctx context.Context, userID int) (*domain.User, error)
}

type MySQLUsersDAO struct {
	db *sql.DB
}

func NewMySQLUsersDAO(db *sql.DB) *MySQLUsersDAO {
	return &MySQLUsersDAO{
		db: db,
	}
}

func (dao *MySQLUsersDAO) CreateUser(ctx context.Context, user *domain.User) error {
	query := "INSERT INTO Users (id, name, password, email, role_id) VALUES (?, ?, ?, ?, ?)"
	_, err := dao.db.ExecContext(ctx, query, user.Id, user.Name, user.Password, user.RoleId)
	return err
}

func (dao *MySQLUsersDAO) UpdateUser(ctx context.Context, user *domain.User) error {
	query := "UPDATE Users SET name = ?, password = ?, email = ?, role_id = ? WHERE id = ?"
	_, err := dao.db.ExecContext(ctx, query, user.Name, user.Password, user.RoleId, user.Id)
	return err
}

func (dao *MySQLUsersDAO) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	query := "DELETE FROM Users WHERE id = ?"
	_, err := dao.db.ExecContext(ctx, query, userID)
	return err
}

func (dao *MySQLUsersDAO) GetUserByEmail(ctx context.Context, userEmail string) (*domain.User, error) {
	query := "SELECT id, name, email, password, role_id FROM Users WHERE email = ?"
	row := dao.db.QueryRowContext(ctx, query, userEmail)

	user := &domain.User{}
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.RoleId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}

	return user, nil
}

func (dao *MySQLUsersDAO) GetUserById(ctx context.Context, userId uuid.UUID) (*domain.User, error) {
	query := "SELECT id, name, email, password, role_id FROM Users WHERE id = ?"
	row := dao.db.QueryRowContext(ctx, query, userId)

	user := &domain.User{}
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.RoleId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}

	return user, nil
}

func (dao *MySQLUsersDAO) GetUsers(ctx context.Context) ([]*domain.User, error) {
	query := "SELECT id, name, email, role_id FROM Users"
	rows, _ := dao.db.QueryContext(ctx, query)

	var users []*domain.User
	for rows.Next() {
		user := new(domain.User)
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.RoleId)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil // User not found
			}
			return nil, err
		}
		users = append(users, user)

	}

	return users, nil
}
