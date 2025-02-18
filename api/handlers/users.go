package handlers

import (
	"log/slog"

	"github.com/devs-group/driplet/api/di"
	"github.com/devs-group/driplet/api/repositories"
	"github.com/devs-group/godi"
	"github.com/go-faster/errors"
	"github.com/gofiber/fiber/v2"
)

type UsersHandler struct {
	usersRepository *repositories.UsersRepository
}

func NewUsersHandler() (*UsersHandler, error) {
	usersRepository, err := godi.Resolve[*repositories.UsersRepository](di.Container)
	if err != nil {
		return nil, errors.Wrap(err, "unable to resolve users repository")
	}
	return &UsersHandler{
		usersRepository: usersRepository,
	}, nil
}

type GetUserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Credits   int    `json:"credits"`
	PublicKey string `json:"public_key"`
}

func (h *UsersHandler) GET_User(c *fiber.Ctx) error {
	u, ok := c.Locals("user").(*repositories.User)
	if !ok {
		return fiber.ErrUnauthorized
	}
	return c.JSON(&GetUserResponse{
		ID:        u.ID,
		Email:     u.Email,
		Credits:   u.Credits,
		PublicKey: u.PublicKey.String,
	})
}

func (h *UsersHandler) PUT_UpdateUsersPublicKey(c *fiber.Ctx) error {
	u, ok := c.Locals("user").(*repositories.User)
	if !ok {
		slog.Error("unable to get user from context")
		return fiber.ErrUnauthorized
	}

	payload := struct {
		PublicKey string `json:"public_key"`
	}{}
	err := c.BodyParser(&payload)
	if err != nil {
		slog.Error("unable to parse request body", "err", err)
		return fiber.ErrBadRequest
	}

	err = h.usersRepository.UpdatePublicKey(u.ID, payload.PublicKey)
	if err != nil {
		slog.Error("unable to update user public key", "err", err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{
		"message": "public key updated successfully",
	})
}
