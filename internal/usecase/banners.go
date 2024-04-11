package usecase

import (
	"context"

	"github.com/TheVovchenskiy/banners/internal/domain"
	"github.com/TheVovchenskiy/banners/pkg/queryManager"
)

type BannerStorage interface {
	GetBanners(ctx context.Context, queryParams queryManager.ParsedQueryParams) ([]domain.Banner, error)
}

type BannerUsecase struct {
}

func NewBannerUsecase() *BannerUsecase {
	return &BannerUsecase{}
}

func (u *BannerUsecase) GetBanners(ctx context.Context, queryParams queryManager.ParsedQueryParams) ([]domain.Banner, error) {
	banners, err := u.GetBanners(ctx, queryParams)
	if err != nil {
		return nil, err
	}
	
	return banners, nil
}
