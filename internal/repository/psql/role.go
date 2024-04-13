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

func (repo *RolePsqlRepo) GetRoleNameById(ctx context.Context, id uint) (string, error) {
	query := `SELECT
				r."name"
			FROM
				public.role r
			WHERE
				r.id = $1`

	var roleName string
	err := repo.db.QueryRow(ctx, query, id).Scan(&roleName)

	if err != nil {
		return "", err
	}
	return roleName, nil
}
