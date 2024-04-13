package psql

import (
	"context"
	"time"

	"github.com/TheVovchenskiy/banners/internal/domain"
	"github.com/TheVovchenskiy/banners/pkg/queryManager"
	"github.com/jackc/pgx/v5"
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
		LEFT JOIN public.banner_tag bt ON b.id = bt.banner_id
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

func (repo *BannerPsqlRepo) AddBanner(ctx context.Context, banner domain.CreateBanner) (newBanner domain.Banner, err error) {
	tx, err := repo.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			return
		}

		if err = tx.Rollback(ctx); err != nil && err != pgx.ErrTxClosed {
			return
		}

		err = nil
	}()

	insertBannerQuery := `INSERT INTO
				public.banner (feature_id, "content", is_active)
			VALUES
				($1, $2, $3)
			RETURNING id, created_at, updated_at`

	var createdAt, updatedAt time.Time
	err = tx.QueryRow(ctx, insertBannerQuery, banner.FeatureId, banner.Content, banner.IsActive).Scan(
		&newBanner.Id,
		&createdAt,
		&updatedAt,
	)

	newBanner.CreatedAt = createdAt.Format(time.RFC1123Z)
	newBanner.UpdatedAt = updatedAt.Format(time.RFC1123Z)
	if err != nil {
		return
	}
	newBanner.FeatureId = banner.FeatureId
	newBanner.Content = banner.Content
	newBanner.IsActive = banner.IsActive

	insertBannerTagQuery := `INSERT INTO 
		public.banner_tag (banner_id, tag_id)
	VALUES
		($1, $2)
	RETURNING tag_id, (SELECT t."name" FROM public.tag t WHERE t.id = $2)`

	tags := []domain.Tag{}

	for _, tagId := range banner.TagIds {
		var tag domain.Tag
		err := tx.QueryRow(ctx, insertBannerTagQuery, newBanner.Id, tagId).Scan(
			&tag.Id,
			&tag.Name,
		)
		if err != nil {
			return domain.Banner{}, err
		}

		tags = append(tags, tag)
	}

	newBanner.Tags = tags

	err = tx.Commit(ctx)
	if err != nil {
		return
	}

	return
}
