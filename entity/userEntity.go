package entity

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"pis/dao"
	"pis/domain"
	"pis/pkg/customErrors"
	"pis/pkg/hash"
	"pis/pkg/jwt"
	"time"
)

type UserEntity interface {
	SignUp(params SignUpParams) (uuid.UUID, error)
	SignIn(params SignInParams) (string, error)
	GetUser(email string) (*domain.User, error)
	GetUsersList() ([]*domain.User, error)
	UpdateUserActiveStatus(params UpdateUserActiveParams) error
	UpdateUserRole(params UpdateUserRoleParams) error
}

type userEntity struct {
	userDAO        dao.MySQLUsersDAO
	roleDAO        dao.MySQLRolesDAO
	tokenGenerator *jwt.TokenGenerator
}

func NewUserEntity(userDAO dao.MySQLUsersDAO) *userEntity {
	return &userEntity{userDAO: userDAO}
}

func (uc userEntity) SignUp(ctx context.Context, params SignUpParams) (uuid.UUID, error) {
	passwordHash, err := hash.ToHashString(params.Password)
	if err != nil {
		return uuid.Nil, errors.New(fmt.Sprintf("password hashing error: %s", err.Error()))
	}
	role, err := uc.roleDAO.GetRoleByName(ctx, domain.USER)
	if err != nil {
		return uuid.Nil, err
	}
	user := domain.NewUser(params.Name, role.Id, params.Email, passwordHash)

	if err = uc.userDAO.CreateUser(ctx, user); err != nil {
		return uuid.Nil, err
	}

	return user.Id, nil
}

func (uc userEntity) SignIn(ctx context.Context, params SignInParams) (string, error) {
	user, err := uc.userDAO.GetUserByEmail(ctx, params.Email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", customErrors.NewNotFoundError("user not found")
	}
	if !hash.IsEqualWithHash(params.Password, user.Password) {
		return "", errors.New("incorrect password")
	}

	return uc.tokenGenerator.GenerateNewAccessToken(jwt.Params{
		Id:   user.Id.String(),
		Role: user.Role.Name,
		Ttl:  24 * time.Hour,
	})
}

func (uc userEntity) GetUser(ctx context.Context, email string) (*domain.User, error) {
	user, err := uc.userDAO.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, customErrors.NewNotFoundError("user not found")
	}

	return user, nil
}

func (uc userEntity) GetUsersList(ctx context.Context) ([]*domain.User, error) {
	return uc.userDAO.GetUsers(ctx)
}

func (uc userEntity) UpdateUserRole(ctx context.Context, params UpdateUserRoleParams) error {
	user, err := uc.userDAO.GetUserById(ctx, params.UserId)
	if err != nil {
		return err
	}
	if user.Role.Name == params.Role {
		return customErrors.NewUpdateError("The user already has this role")
	}
	switch params.Role {
	case domain.USER, domain.MANAGER, domain.ADMIN:
		role, err := uc.roleDAO.GetRoleByName(ctx, params.Role)
		if err != nil {
			return err
		}
		user.RoleId = role.Id
	default:
		return customErrors.NewUpdateError("Wrong role name")
	}
	return uc.userDAO.UpdateUser(ctx, user)
}

type SignUpParams struct {
	Name     string
	Email    string
	Password string
}

type SignInParams struct {
	Email    string
	Password string
}

type UpdateUserActiveParams struct {
	UserId   uuid.UUID
	IsActive bool
}

type UpdateUserRoleParams struct {
	UserId uuid.UUID
	Role   string
}
