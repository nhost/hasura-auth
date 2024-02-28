package cmd

import (
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/urfave/cli/v2"
)

func getDBPool(cCtx *cli.Context) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(cCtx.String(flagPostgresConnection))
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}

	config.MaxConns = 4
	config.MinConns = 1
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = time.Minute * 30 //nolint:gomnd
	config.HealthCheckPeriod = time.Minute

	pool, err := pgxpool.NewWithConfig(cCtx.Context, config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return pool, nil
}
