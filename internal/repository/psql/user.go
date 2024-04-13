package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/TheVovchenskiy/banners/internal/domain"
	"github.com/TheVovchenskiy/banners/internal/repository"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserPg struct {
	db *pgxpool.Pool
}

func NewUserPg(db *pgxpool.Pool) *UserPg {
	return &UserPg{
		db: db,
	}
}

func (repo *UserPg) StoreUser(ctx context.Context, user *domain.User) (uint, uint, error) {
	query := `INSERT
				INTO public.user_profile (
					role_id,
					username,
					password_hash,
					salt
				)
			VALUES (
				(
					SELECT 
						id
					FROM 
						public.role 
					WHERE 
						name = $1
				),
				$2,
				$3,
				$4
			)
			RETURNING id, role_id;`

	var id, roleId uint
	err := repo.db.QueryRow(
		ctx,
		query,
		user.Role.Name,
		user.Username,
		user.PasswordHash,
		user.Salt,
	).
		Scan(&id, &roleId)

	if err != nil {
		if pgErr, ok := err.(pgx.PgError); ok {
			switch pgErr.Code {
			case "23505":
				return 0, 0, repository.ErrAccountAlreadyExists
			}
		}
		return 0, 0, err
	}

	return id, roleId, nil
}

func (repo *UserPg) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	query := `SELECT
					up.id,
					up.role_id,
					up.username,
					up.password_hash,
					up.salt
				FROM
					public.user_profile up
				WHERE
					up.username = $1
				`
	var user domain.User
	var passwordHash []byte
	var salt []byte
	err := repo.db.QueryRow(
		ctx,
		query,
		username,
	).
		Scan(
			&user.Id,
			&user.Role.Id,
			&user.Username,
			&passwordHash,
			&salt,
		)

	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("%w: %s", repository.ErrNoUserFound, username)
		}
		return domain.User{}, err
	}
	user.PasswordHash = string(passwordHash)
	user.Salt = string(salt)

	return user, nil
}
