package storage

import (
	"context"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/tern/migrate"
	"github.com/pkg/errors"
	"github.com/tarent/logrus"
)

type PostgreSqlStorage struct {
	pool *pgxpool.Pool
}

type Config struct {
	Username     string
	Password     string
	DatabaseName string
	Host         string
	Port         int
}

const maxPoolConns = 10

func NewPostgreSqlStorage(c Config) (*PostgreSqlStorage, error) {
	var pool *pgxpool.Pool
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable pool_max_conns=%d",
		c.Username, c.Password, c.Host, c.Port, c.DatabaseName, maxPoolConns)
	err := backoff.Retry(func() (err error) {
		pool, err = pgxpool.Connect(context.Background(), dsn)
		return
	}, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 5))
	if err != nil {
		return nil, errors.Wrapf(err, "unable to connect to database")
	}

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, errors.Wrapf(err, "unable to connect to database")
	}

	err = migrateDatabase(conn.Conn())
	if err != nil {
		return nil, errors.Wrapf(err, "unable to migrate database")
	}
	conn.Release()

	return &PostgreSqlStorage{
		pool: pool,
	}, nil
}

func migrateDatabase(conn *pgx.Conn) error {
	migrator, err := migrate.NewMigrator(context.Background(), conn, "schema_version")
	if err != nil {
		return errors.Wrapf(err, "unable to create a migrator")
	}

	err = migrator.LoadMigrations("./storage/migrations")
	if err != nil {
		return errors.Wrapf(err, "unable to load migrations")
	}

	err = migrator.Migrate(context.Background())
	if err != nil {
		return errors.Wrapf(err, "unable to migrate")
	}

	ver, err := migrator.GetCurrentVersion(context.Background())
	if err != nil {
		return errors.Wrapf(err, "unable to get current schema version")
	}

	logrus.Infof("Migration done. Current schema version: %v", ver)
	return nil
}
