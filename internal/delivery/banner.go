package delivery

import (
	"net/http"

	"github.com/TheVovchenskiy/banners/pkg/queryManager"
	"github.com/TheVovchenskiy/banners/pkg/response"
	"github.com/TheVovchenskiy/banners/internal/usecase"
)

type BannerHandler struct {
	bannerUsecase *usecase.BannerUsecase
}

func NewBannerHandler(bannerUsecase *usecase.BannerUsecase) *BannerHandler {
	return &BannerHandler{
		bannerUsecase: bannerUsecase,
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
//	@Success		200			{array}	domain.Banner	"An array of banners"
//	@Failure		401			"User is unauthorized"
//	@Failure		403			"Forbidden"
//	@Failure		500			{object} serverErrors.APIError "Internal server error"
//	@Router			/banners [get]
func (handler *BannerHandler) HandleGetBanners(w http.ResponseWriter, r *http.Request) {
	queryParams, err := queryManager.ParseUrlQuery(*r.URL, []queryManager.QueryParam{
		{
			Name:         "feature_id",
			Type:         queryManager.IntType,
			Required:     false,
			DefaultValue: -1,
		},
		{
			Name:         "tag_id",
			Type:         queryManager.IntType,
			Required:     false,
			DefaultValue: -1,
		},
		{
			Name:         "limit",
			Type:         queryManager.IntType,
			Required:     false,
			DefaultValue: 10,
		},
		{
			Name:         "offset",
			Type:         queryManager.IntType,
			Required:     false,
			DefaultValue: 0,
		},
	})

	if err != nil {
		response.ServeJsonError(r.Context(), w, err)
		return
	}

	banners, err := handler.bannerUsecase.GetBanners(r.Context(), queryParams)
	if err != nil {
		response.ServeJsonError(r.Context(), w, err)
		return
	}

	response.ServerJsonData(r.Context(), w, banners)
}
