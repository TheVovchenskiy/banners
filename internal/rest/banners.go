package rest

import "net/http"

type BannerHandler struct {
	// authUsecase *usecase.AuthUsecase
}

func NewBannerHandler() *BannerHandler {
	return &BannerHandler{
		// authUsecase: usecase.NewAuthUsecase(authStorage),
	}
}

// HandleGetBanners godoc
//
//	@Summary		Handle getting list of banners
//	@Description	Get all banners with filtering by feature and/or tag
//	@Tags			banners
//	@Produce		json
//
//	@Param			token		header	string			false	"Admin's token"
//	@Param			feature_id	query	int				false	"Feature's id"
//	@Param			tag_id		query	int				false	"Tags's id"
//	@Param			limit		query	int				false	"Limit"
//	@Param			offset		query	int				false	"Offset"
//
//	@Success		200			{array}	model.Banner	"An array of banners"
//	@Failure		401			"User is unauthorized"
//	@Failure		403			"Forbidden"
//	@Failure		500			{object} serverErrors.APIError "Internal server error"
//	@Router			/banners [get]
func (handler *BannerHandler) HandleGetBanners(w http.ResponseWriter, r *http.Request) {
}
