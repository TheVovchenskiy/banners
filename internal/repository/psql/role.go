package psql

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RolePsqlRepo struct {
	db *pgxpool.Pool
}

func NewRolePg(db *pgxpool.Pool) *RolePsqlRepo {
	return &RolePsqlRepo{
		db: db,
	}
}

func (repo *RolePsqlRepo) GetAllRoles(ctx context.Context) ([]string, error) {
	query := `SELECT "name" FROM public.role`

	rows, err := repo.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := []string{}

	for rows.Next() {
		var role string

		err := rows.Scan(
			&role,
		)

		if err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	return roles, nil
}
