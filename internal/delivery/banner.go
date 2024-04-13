package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TheVovchenskiy/banners/internal/domain"
	"github.com/TheVovchenskiy/banners/internal/repository/psql"
	"github.com/TheVovchenskiy/banners/internal/usecase"
	"github.com/TheVovchenskiy/banners/pkg/queryManager"
	"github.com/TheVovchenskiy/banners/pkg/response"
	"github.com/TheVovchenskiy/banners/pkg/serverErrors"
)

type BannerHandler struct {
	bannerUsecase *usecase.BannerUsecase
}

func NewBannerHandler(bannerStorage psql.BannerPsqlRepo) *BannerHandler {
	return &BannerHandler{
		bannerUsecase: usecase.NewBannerUsecase(&bannerStorage),
	}
}

// HandleGetBanners godoc
//
//	@Summary		Handle getting list of banners
//	@Description	Get all banners with filtering by feature and/or tag
//	@Tags			banners
//	@Produce		json
//
//	@Param			Authorization	header	string			true	"Admin's token"	example(Bearer your_access_token)
//	@Param			feature_id		query	int				false	"Feature's id"
//	@Param			tag_id			query	int				false	"Tags's id"
//	@Param			limit			query	int				false	"Limit"
//	@Param			offset			query	int				false	"Offset"
//
//	@Success		200				{array}	domain.Banner	"An array of banners"
//	@Failure		401				"User is unauthorized"
//	@Failure		403				"Forbidden"
//	@Failure		500				{object}	serverErrors.APIError	"Internal server error"
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

// HandlePostBanners godoc
//
//	@Summary		Handle creating new banner
//	@Description	Create new banner
//	@Tags			banners
//	@Accept			json
//	@Produce		json
//
//	@Param			Authorization	header		string					true	"Admin's token"	example(Bearer your_access_token)
//	@Param			tagIds			body		[]uint					true	"Ids of tags" example([1, 2])
//	@Param			featureId		body		uint					true	"Feature's id" example(1)
//	@Param			content			body		string					true	"Banner's content" example("{\"title\": \"some_title\", \"text\": \"some_text\", \"url\": \"some_url\"}")
//
//	@Success		200				{object}	domain.Banner			"Created"
//	@Failure		400				{object}	serverErrors.APIError	"Bad request"
//	@Failure		401				"User is unauthorized"
//	@Failure		403				"Forbidden"
//	@Failure		500				{object}	serverErrors.APIError	"Internal server error"
//	@Router			/banners [post]
func (handler *BannerHandler) HandlePostBanners(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	createBanner := new(domain.CreateBanner)
	err := decoder.Decode(createBanner)
	if err != nil {
		err = fmt.Errorf("%w: unable to decode body", serverErrors.ErrInvalidBody)
		response.ServeJsonError(r.Context(), w, err)
		return
	}
	
	banner, err := handler.bannerUsecase.AddBanner(r.Context(), *createBanner)
	if err != nil {
		response.ServeJsonError(r.Context(), w, err)
		return
	}

	response.ServerJsonData(r.Context(), w, banner)
}
