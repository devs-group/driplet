package main

import (
	"github.com/devs-group/driplet/api/auth"
	"github.com/devs-group/driplet/api/di"
	"github.com/devs-group/driplet/api/handlers"
	"github.com/devs-group/driplet/api/middlewares"
	"github.com/devs-group/driplet/api/repositories"
	"github.com/devs-group/godi"
	"github.com/go-faster/errors"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) error {
	tokenValidator, err := godi.Resolve[*auth.TokenValidator](di.Container)
	if err != nil {
		return errors.Wrap(err, "unable to resolve token validator")
	}
	userRepository, err := godi.Resolve[*repositories.UsersRepository](di.Container)
	if err != nil {
		return errors.Wrap(err, "unable to resolve users repository")
	}

	usersHandler, err := handlers.NewUsersHandler()
	if err != nil {
		return errors.Wrap(err, "unable to create new users handler")
	}
	healthHandler, err := handlers.NewHealthHandler()
	if err != nil {
		return errors.Wrap(err, "unable to create new health handler")
	}

	v1 := app.Group(
		"/api/v1",
		middlewares.RequireAuth(middlewares.AuthConfig{
			TokenValidator:  tokenValidator,
			UsersRepository: userRepository,
		}),
	)
	app.Get("/health", healthHandler.GET_health)
	v1.Get("/user", usersHandler.GET_User)

	return nil
}
