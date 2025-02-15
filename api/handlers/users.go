package handlers

import (
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
	ID      string `json:"id"`
	Email   string `json:"email"`
	Credits int    `json:"credits"`
}

func (h *UsersHandler) GET_User(c *fiber.Ctx) error {
	user, err := h.usersRepository.FindByEmail("") // TODO: Get user email from context
	if err != nil {
		return fiber.ErrNotFound
	}
	return c.JSON(&GetUserResponse{
		ID:      user.ID,
		Email:   user.Email,
		Credits: user.Credits,
	})
}
