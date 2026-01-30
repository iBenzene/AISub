package repository

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Wei-Shaw/sub2api/internal/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

func openPostgresDB(cfg *config.Config) (*sql.DB, error) {
	dsn := cfg.Database.DSNWithTimezone(cfg.Timezone)
	if useSimpleProtocol(cfg) {
		pgxConfig, err := pgx.ParseConfig(dsn)
		if err != nil {
			return nil, err
		}
		pgxConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
		pgxConfig.StatementCacheCapacity = 0
		pgxConfig.DescriptionCacheCapacity = 0
		log.Printf("[Database] Using pgx simple protocol for pooler compatibility")
		return stdlib.OpenDB(*pgxConfig), nil
	}

	return sql.Open("postgres", dsn)
}

func useSimpleProtocol(cfg *config.Config) bool {
	if cfg == nil {
		return false
	}
	if raw, ok := os.LookupEnv("DATABASE_SIMPLE_PROTOCOL"); ok {
		if v, err := strconv.ParseBool(strings.TrimSpace(raw)); err == nil {
			return v
		}
	}
	host := strings.ToLower(strings.TrimSpace(cfg.Database.Host))
	return strings.Contains(host, "pooler") || strings.Contains(host, "pgbouncer")
}
