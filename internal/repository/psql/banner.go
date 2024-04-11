package psql

import (
	"context"

	"github.com/TheVovchenskiy/banners/internal/domain"
	"github.com/TheVovchenskiy/banners/pkg/queryManager"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BannerPsqlRepo struct {
	db *pgxpool.Pool
}

func NewBannerPg(db *pgxpool.Pool) *BannerPsqlRepo {
	return &BannerPsqlRepo{
		db: db,
	}
}

func (repo *BannerPsqlRepo) GetBanners(ctx context.Context, queryParams queryManager.ParsedQueryParams) ([]domain.Banner, error) {
	query, args := queryManager.BuildSQLQuery(
		`SELECT 
			b.id,
			b.feature_id,
			b."content",
			b.is_active,
			b.created_at,
			b.updated_at,
			t.id,
			t."name",
			t.created_at,
			t.updated_at
		FROM
			public.banner b
		LEFT JOIN public.banner_tag bt ON b.id = bt.bunner_id
		LEFT JOIN public.tag t ON bt.tag_id = t.id`,
		queryParams,
		[]string{"feature_id", "tag_id"},
	)

	rows, err := repo.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	banners := []domain.Banner{}
	bannerTags := map[uint][]domain.Tag{}
	for rows.Next() {
		var banner domain.Banner
		var tag domain.Tag

		err := rows.Scan(
			&banner.Id,
			&banner.FeatureId,
			&banner.Content,
			&banner.IsActive,
			&banner.CreatedAt,
			&banner.UpdatedAt,
			&tag.Id,
			&tag.Name,
			&tag.CreatedAt,
			&tag.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		if _, exist := bannerTags[banner.Id]; exist {
			bannerTags[banner.Id] = append(bannerTags[banner.Id], tag)
		} else {
			bannerTags[banner.Id] = append(bannerTags[banner.Id], tag)
		}

		banners = append(banners, banner)
	}

	for i, banner := range banners {
		banners[i].Tags = bannerTags[banner.Id]
	}

	return banners, nil
}
