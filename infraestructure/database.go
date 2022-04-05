package infraestructure

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(env Env) *gorm.DB {
	dbConfig := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		env.POSTGRES_USERNAME,
		env.POSTGRES_PASSWORD,
		env.POSTGRES_HOST,
		env.POSTGRES_PORT,
		env.POSTGRES_DATABASE,
	)

	db, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})

	if err != nil {
		panic(errors.Wrap(err, "Error to connect at Database"))
	}

	return db
}
