package app

import (
	"context"
	"fmt"

	"github.com/TheVovchenskiy/banners/configs"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetPostgres(pgConfig configs.PgConfig) (conn *pgxpool.Pool, err error) {
	for _, host := range []string{"pg_db", "localhost"} {
		pgConnStr := fmt.Sprintf(
			"user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
			pgConfig.PgUser,
			pgConfig.PgDBName,
			pgConfig.PgPassword,
			host,
			pgConfig.PgPort,
		)
		conn, err = pgxpool.New(context.Background(), pgConnStr)
		if err != nil {
			continue
		}
		err = conn.Ping(context.Background())
		if err != nil {
			continue
		}

		break
	}

	return conn, err
}
