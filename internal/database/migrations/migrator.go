package migrations

import (
	"context"
	"database/sql"
	migrationssql "github.com/bitmagnet-io/bitmagnet/migrations"
	goose "github.com/pressly/goose/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Params struct {
	fx.In
	DB     *sql.DB
	Logger *zap.SugaredLogger
}

func New(p Params) Migrator {
	initGoose(p.Logger)
	return &migrator{
		db: p.DB,
	}
}

func initGoose(logger *zap.SugaredLogger) {
	goose.SetLogger(gooseLogger{logger.Named("migrator")})
	goose.SetBaseFS(migrationssql.FS)
	err := goose.SetDialect("postgres")
	if err != nil {
		panic(err)
	}
}

type Migrator interface {
	Up(ctx context.Context) error
	UpTo(ctx context.Context, version int64) error
	Down(ctx context.Context) error
	DownTo(ctx context.Context, version int64) error
}

type migrator struct {
	db *sql.DB
}

func (m *migrator) Up(ctx context.Context) error {
	return goose.UpContext(ctx, m.db, ".")
}

func (m *migrator) UpTo(ctx context.Context, version int64) error {
	return goose.UpToContext(ctx, m.db, ".", version)
}

func (m *migrator) Down(ctx context.Context) error {
	return goose.DownContext(ctx, m.db, ".")
}

func (m *migrator) DownTo(ctx context.Context, version int64) error {
	return goose.DownToContext(ctx, m.db, ".", version)
}