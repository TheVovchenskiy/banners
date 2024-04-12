package usecase

import (
	"context"

	"github.com/TheVovchenskiy/banners/configs"
	"github.com/TheVovchenskiy/banners/internal/domain"
	"github.com/TheVovchenskiy/banners/pkg/token"
	"github.com/TheVovchenskiy/banners/pkg/validator"
	"github.com/google/uuid"
)

type RoleStorage interface {
	GetAllRoles(ctx context.Context) ([]string, error)
}

type UserStorage interface {
	StoreUser(ctx context.Context, user *domain.User) (uint, uint, error)
	// GetUserByUsername(ctx context.Context, username string) (domain.User, error)
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

// func (u *AuthUsecase) LoginUser(loginInput domain.LoginInput) (*domain.User, error) {
// 	loginInput.Trim()
// 	user, err := u.userStorage.GetUserByUsername(context.Background(), loginInput.Username)
// 	if err != nil {
// 		if errors.Is(err, repository.ErrNoUserFound) {
// 			err = ErrInvalidLoginData
// 		}
// 		return nil, err
// 	}

// 	if !hash.MatchPasswords(user.PasswordHash, loginInput.Password, user.Salt) {
// 		return nil, ErrInvalidLoginData
// 	}

// 	user.AccessToken, err = token.GenerateAccesToken(user.Id, loginInput.Username)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &user, nil

// }
