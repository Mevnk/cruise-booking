package business

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

type UserLogic interface {
	SignUp(ctx context.Context, params SignUpParams) (uuid.UUID, error)
	SignIn(ctx context.Context, params SignInParams) (string, error)
	GetUser(ctx context.Context, email string) (*domain.User, error)
	GetUsersList(ctx context.Context) ([]*domain.User, error)
	UpdateUserRole(ctx context.Context, params UpdateUserRoleParams) error
	DeleteUser(ctx context.Context, userId uuid.UUID) (err error)
}

type userLogic struct {
	userDAO        dao.MySQLUsersDAO
	roleDAO        dao.MySQLRolesDAO
	tokenGenerator *jwt.TokenGenerator
}

func NewUserLogic(userDAO dao.MySQLUsersDAO) *userLogic {
	return &userLogic{userDAO: userDAO}
}

func (uc userLogic) SignUp(ctx context.Context, params SignUpParams) (uuid.UUID, error) {
	userCheck, err := uc.userDAO.GetUserByEmail(ctx, params.Email)
	if userCheck != nil {
		return userCheck.Id, customErrors.NewCustomError(customErrors.User, customErrors.Creation)
	}

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

func (uc userLogic) SignIn(ctx context.Context, params SignInParams) (string, error) {
	user, err := uc.userDAO.GetUserByEmail(ctx, params.Email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", customErrors.NewCustomError(customErrors.User, customErrors.NotFound)
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

func (uc userLogic) GetUser(ctx context.Context, email string) (*domain.User, error) {
	user, err := uc.userDAO.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, customErrors.NewCustomError(customErrors.User, customErrors.NotFound)
	}

	return user, nil
}

func (uc userLogic) GetUsersList(ctx context.Context) ([]*domain.User, error) {
	return uc.userDAO.GetUsers(ctx)
}

func (uc userLogic) UpdateUserRole(ctx context.Context, params UpdateUserRoleParams) error {
	user, err := uc.userDAO.GetUserById(ctx, params.UserId)
	if err != nil {
		return err
	}
	switch params.Role {
	case domain.USER, domain.MANAGER, domain.ADMIN:
		role, err := uc.roleDAO.GetRoleByName(ctx, params.Role)
		if err != nil {
			return err
		}
		user.RoleId = role.Id
	default:
		return customErrors.NewCustomError(customErrors.User, customErrors.Update)
	}
	return uc.userDAO.UpdateUser(ctx, user)
}

func (uc userLogic) DeleteUser(ctx context.Context, id uuid.UUID) (err error) {
	err = uc.userDAO.DeleteUser(ctx, id)
	if err != nil {
		return customErrors.NewCustomError(customErrors.User, customErrors.Deletion)
	}
	return nil
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
