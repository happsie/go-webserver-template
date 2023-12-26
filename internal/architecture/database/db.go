package database

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/happsie/go-webserver-template/internal/architecture"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

func Init(log *slog.Logger, config architecture.Config) (*sqlx.DB, error) {
	log.Info("connecting to database")
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@(%s)/%s?parseTime=true", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Database))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})
	if err != nil {
		return nil, err
	}
	log.Info("database connection established")
	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", config.Database.MigrationSrc), "mysql", driver)
	if err != nil {
		return nil, err
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, err
	}
	log.Info("database migrations completed")
	return db, nil
}
