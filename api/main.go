package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/devs-group/driplet/api/config"
	"github.com/devs-group/driplet/api/di"
	"github.com/devs-group/driplet/api/migrations"
	"github.com/devs-group/driplet/pkg/db"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "api",
		Usage: "api for driplet",
		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "executes the api",
				Action: func(c *cli.Context) error {
					app := fiber.New()
					di.Init()       // initializing dependency injection container
					InitRoutes(app) // initializing http routes
					return app.Listen(getPort())
				},
			},
			{
				Name:  "migrate",
				Usage: "database migration commands",
				Subcommands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create a new migration file",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "name",
								Aliases:  []string{"n"},
								Usage:    "name of the migration",
								Required: true,
							},
						},
						Action: func(c *cli.Context) error {
							return migrations.CreateMigrationFile(c.String("name"))
						},
					},
					{
						Name:  "up",
						Usage: "run all pending migrations",
						Action: func(c *cli.Context) error {
							database, err := db.Connect(db.DefaultConfig())
							if err != nil {
								return fmt.Errorf("failed to connect to database: %w", err)
							}
							defer database.Close()
							return migrations.RunMigrations(database.DB, "up")
						},
					},
					{
						Name:  "down",
						Usage: "rollback the last migration",
						Action: func(c *cli.Context) error {
							database, err := db.Connect(db.DefaultConfig())
							if err != nil {
								return fmt.Errorf("failed to connect to database: %w", err)
							}
							defer database.Close()
							return migrations.RunMigrations(database.DB, "down")
						},
					},
					{
						Name:  "reset",
						Usage: "rollback all migrations",
						Action: func(c *cli.Context) error {
							database, err := db.Connect(db.DefaultConfig())
							if err != nil {
								return fmt.Errorf("failed to connect to database: %w", err)
							}
							defer database.Close()
							return migrations.RunMigrations(database.DB, "reset")
						},
					},
					{
						Name:  "status",
						Usage: "print migrations status",
						Action: func(c *cli.Context) error {
							database, err := db.Connect(db.DefaultConfig())
							if err != nil {
								return fmt.Errorf("failed to connect to database: %w", err)
							}
							defer database.Close()
							return migrations.RunMigrations(database.DB, "status")
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getPort() string {
	port := 9000
	if config.PORT != "" {
		var err error
		port, err = strconv.Atoi(config.PORT)
		if err != nil {
			log.Fatal("PORT env var must be a number parsable to an int")
		}
	}
	return fmt.Sprintf(":%d", port)
}
