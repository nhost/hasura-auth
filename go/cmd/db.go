package cmd

import (
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/urfave/cli/v2"
)

const (
	poolMinMaxConns          = 4
	poolMinMinConns          = 1
	poolMinMaxConnLifetime   = time.Hour
	poolMinMaxConnIdleTime   = time.Minute * 30
	poolMinHealthCheckPeriod = time.Minute
)

func getDBPool(cCtx *cli.Context) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(cCtx.String(flagPostgresConnection))
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}

	if config.MaxConns < poolMinMaxConns { //
		config.MaxConns = poolMinMaxConns
	}
	if config.MinConns < poolMinMinConns {
		config.MinConns = poolMinMinConns
	}
	if config.MaxConnLifetime < poolMinMaxConnLifetime {
		config.MaxConnLifetime = poolMinMaxConnLifetime
	}
	if config.MaxConnIdleTime < poolMinMaxConnIdleTime {
		config.MaxConnIdleTime = poolMinMaxConnIdleTime
	}
	if config.HealthCheckPeriod < poolMinHealthCheckPeriod {
		config.HealthCheckPeriod = poolMinHealthCheckPeriod
	}

	pool, err := pgxpool.NewWithConfig(cCtx.Context, config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return pool, nil
}
