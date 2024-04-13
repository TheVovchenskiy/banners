package usecase

import (
	"context"
	"errors"

	"github.com/TheVovchenskiy/banners/configs"
	"github.com/TheVovchenskiy/banners/internal/domain"
	"github.com/TheVovchenskiy/banners/internal/repository"
	"github.com/TheVovchenskiy/banners/pkg/contextManager"
	"github.com/TheVovchenskiy/banners/pkg/hash"
	"github.com/TheVovchenskiy/banners/pkg/logging"
	"github.com/TheVovchenskiy/banners/pkg/token"
	"github.com/TheVovchenskiy/banners/pkg/validator"
	"github.com/google/uuid"
)

type RoleStorage interface {
	GetAllRoles(ctx context.Context) ([]string, error)
	GetRoleNameById(ctx context.Context, id uint) (string, error)
}

type UserStorage interface {
	StoreUser(ctx context.Context, user *domain.User) (uint, uint, error)
	GetUserByUsername(ctx context.Context, username string) (domain.User, error)
}

type AuthUsecase struct {
	roleStorage RoleStorage
	userStorage UserStorage
}

func NewAuthUsecase(userStorage UserStorage, roleStorage RoleStorage) *AuthUsecase {
	return &AuthUsecase{
		roleStorage: roleStorage,
		userStorage: userStorage,
	}
}

func (u *AuthUsecase) RegisterUser(ctx context.Context, registerInput domain.RegisterInput) (*domain.User, error) {
	registerInput.Trim()
	allowedRoles, err := u.roleStorage.GetAllRoles(ctx)
	if err != nil {
		return nil, err
	}

	err = validator.ValidateRegisterInput(registerInput, allowedRoles, configs.AdminKey)
	if err != nil {
		return nil, err
	}

	salt := uuid.NewString()
	user := registerInput.ToUser(salt)

	user.Id, user.Role.Id, err = u.userStorage.StoreUser(ctx, user)
	if err != nil {
		return nil, err
	}

	user.AccessToken, err = token.GenerateAccesToken(user.Id, registerInput.Username, user.Role.Name)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *AuthUsecase) LoginUser(ctx context.Context, loginInput domain.LoginInput) (*domain.User, error) {
	contextLogger := contextManager.GetContextLogger(ctx)
	loginInput.Trim()
	user, err := u.userStorage.GetUserByUsername(context.Background(), loginInput.Username)
	if err != nil {
		if errors.Is(err, repository.ErrNoUserFound) {
			logging.LogError(contextLogger, err, "while getting user by username")
			err = ErrInvalidLoginData
		}
		return nil, err
	}

	user.Role.Name, err = u.roleStorage.GetRoleNameById(ctx, user.Role.Id)
	if err != nil {
		return nil, err
	}

	if !hash.MatchPasswords(user.PasswordHash, loginInput.Password, user.Salt) {
		return nil, ErrInvalidLoginData
	}

	user.AccessToken, err = token.GenerateAccesToken(user.Id, loginInput.Username, user.Role.Name)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
