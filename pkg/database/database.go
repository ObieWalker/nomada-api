package database

import (
    "log"
    "os"
    "github.com/joho/godotenv"
    "github.com/obiewalker/nomada-api/pkg/postgres"
		"github.com/obiewalker/nomada-api/pkg/database/models"
    "gorm.io/gorm"
)

type Dbinstance struct {
	Db *gorm.DB
}

var Instance Dbinstance

func ConnectDb() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &postgres.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := postgres.NewConnection(config)

	if err != nil {
		log.Fatal("could not load the database")
	}

	err = model.MigrateUsers(db)
	if err != nil {
		log.Fatal("could not migrate user")
	}

  err = model.MigrateBikes(db)
	if err != nil {
		log.Fatal("could not migrate bikes")
	}

	err = model.MigrateGroups(db)
	if err != nil {
		log.Fatal("could not migrate bikes")
	}

	Instance = Dbinstance{
		Db: db,
	}
}