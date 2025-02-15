package di

import (
	"log"

	"github.com/devs-group/driplet/api/auth"
	"github.com/devs-group/driplet/api/config"
	"github.com/devs-group/driplet/api/repositories"
	"github.com/devs-group/driplet/pkg/db"
	"github.com/devs-group/godi"
	"github.com/jmoiron/sqlx"
)

var Container = godi.New()

func Init() {
	godi.Register(Container, func() *sqlx.DB {
		database, err := db.Connect(db.DefaultConfig())
		if err != nil {
			log.Fatal(err)
		}
		return database
	}, godi.Singleton)

	godi.Register(Container, func() *auth.TokenValidator {
		return auth.NewTokenValidator(config.GOOGLE_CLIENT_ID, config.ALLOWED_EXTENSION_CLIENT_IDS)
	}, godi.Singleton)

	godi.Register(Container, func() *repositories.UsersRepository {
		db, _ := godi.Resolve[*sqlx.DB](Container)
		return &repositories.UsersRepository{DB: db}
	}, godi.Singleton)
}
