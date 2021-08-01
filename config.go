package main

import (
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type appData struct {
	config *config

	db *pgxpool.Pool
}

type config struct {
	db *dbConfig
}

type dbConfig struct {
	host string
	user string
	pass string
}

func initAppData(log *logrus.Logger) (*appData, error) {
	c, err := initConfig(log)
	if err != nil {
		return nil, fmt.Errorf("cannot initialize digital trainer config: %w", err)
	}

	db, err := initDatabase(log, c)
	if err != nil {
		return nil, fmt.Errorf("cannot initialize digital trainer database: %w", err)
	}

	return &appData{
		config: c,
		db:     db,
	}, nil
}

func initConfig(log *logrus.Logger) (*config, error) {
	viper.SetEnvPrefix("DTB")
	viper.AutomaticEnv()

	config := &config{
		db: &dbConfig{
			host: viper.GetString("db_host"),
			user: viper.GetString("db_user"),
			pass: viper.GetString("db_pass"),
		},
	}

	return config, nil
}
