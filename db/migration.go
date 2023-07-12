package db

import (
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

type config struct {
	src   string
	db    string
	steps int
	Error error
}

type option interface {
	apply(*config)
}

type optionFunc func(cfg *config)

func (f optionFunc) apply(cfg *config) {
	f(cfg)
}

func WithSrc(src string) option {
	return optionFunc(func(cfg *config) {
		cfg.src = src
	})
}

func WithDB(db string) option {
	return optionFunc(func(cfg *config) {
		cfg.db = db
	})
}

func WithSteps(steps int) option {
	return optionFunc(func(cfg *config) {
		cfg.steps = steps
	})
}

func defaultCfg() *config {
	cfg := &config{
		src:   "",
		db:    os.Getenv("DB_URL"),
		steps: 0,
		Error: nil,
	}
	src, err := filepath.Abs(os.Getenv("MIGRATION_SRC_URL"))
	cfg.src = "file://" + src
	cfg.Error = err
	return cfg
}

func Migrate(opts ...option) error {
	cfg := defaultCfg()
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if cfg.Error != nil {
		return cfg.Error
	}

	m, err := migrate.New(cfg.src, cfg.db)
	if err != nil {
		return err
	}
	
	if cfg.steps != 0 {
		return m.Steps(cfg.steps)
	}
	return m.Up()
}
