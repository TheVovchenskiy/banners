package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TheVovchenskiy/banners/internal/domain"
	"github.com/TheVovchenskiy/banners/internal/usecase"
	"github.com/TheVovchenskiy/banners/pkg/response"
	"github.com/TheVovchenskiy/banners/pkg/serverErrors"
)

type AuthHandler struct {
	authUsecase *usecase.AuthUsecase
}

func NewAuthHandler(authStorage usecase.UserStorage, roleStorage usecase.RoleStorage) *AuthHandler {
	return &AuthHandler{
		authUsecase: usecase.NewAuthUsecase(authStorage, roleStorage),
	}
}

// HandleRegistration godoc
//
//	@Summary		Handle registrating new users
//	@Description	Register new user or admin
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//
//	@Param			username	body		string					true	"Username of new user"
//	@Param			password	body		string					true	"Password of new user"
//	@Param			role		body		string					true	"Desired user's role: either 'admin' or 'user'"
//	@Param			admin_key	body		string					false	"Admin's registration key"
//
//	@Success		200			{object}	domain.User				"An object wtih new user's info"
//	@Failure		400			{object}	serverErrors.APIError	"Bad request"
//	@Failure		409			{object}	serverErrors.APIError	"Conflict"
//	@Failure		500			{object}	serverErrors.APIError	"Internal server error"
//	@Router			/register [post]
func (handler *AuthHandler) HandleRegistration(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	registrationInput := new(domain.RegisterInput)
	err := decoder.Decode(registrationInput)
	if err != nil {
		err = fmt.Errorf("%w: unable to decode body", serverErrors.ErrInvalidBody)
		response.ServeJsonError(r.Context(), w, err)
		return
	}

	user, err := handler.authUsecase.RegisterUser(r.Context(), *registrationInput)
	if err != nil {
		response.ServeJsonError(r.Context(), w, err)
		return
	}

	response.ServerJsonData(r.Context(), w, user)
}

// HandleLogin godoc
//
//	@Summary		Handle login
//	@Description	Login user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//
//	@Param			username	body		string					true	"Username"
//	@Param			password	body		string					true	"User's password"
//
//	@Success		200			{object}	domain.User				"An object wtih new user's info"
//	@Failure		400			{object}	serverErrors.APIError	"Bad request"
//	@Failure		500			{object}	serverErrors.APIError	"Internal server error"
//	@Router			/register [post]
func (handler *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	loginInput := new(domain.LoginInput)
	err := decoder.Decode(loginInput)
	if err != nil {
		err = fmt.Errorf("%w: unable to decode body", serverErrors.ErrInvalidBody)
		response.ServeJsonError(r.Context(), w, err)
		return
	}
	user, err := handler.authUsecase.LoginUser(r.Context(), *loginInput)
	if err != nil {
		response.ServeJsonError(r.Context(), w, err)
		return
	}

	response.ServerJsonData(r.Context(), w, user)

}
