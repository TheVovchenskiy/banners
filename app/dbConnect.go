package app

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetPostgres() (conn *pgxpool.Pool, err error) {
	for _, host := range []string{"pg_db", "localhost"} {
		pgConnStr := fmt.Sprintf(
			"user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
			os.Getenv("PG_USER"),
			os.Getenv("PG_DBNAME"),
			os.Getenv("PG_PASSWORD"),
			host,
			os.Getenv("PG_PORT"),
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
