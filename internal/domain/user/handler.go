package user

import (
	"fmt"
	"infra-base-go/pkg/util"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetAll(c echo.Context) error {
	users, err := h.Service.GetAllUsers()

	if err != nil {
		return util.NewJsonError(c, http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, users)
}

func (h *Handler) GetById(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return util.NewJsonError(c, http.StatusBadGateway, "Param id in path is required")
	}

	user, err := h.Service.GetUserById(id)
	if err != nil {
		return util.NewJsonError(c, http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, user)
}

func (h *Handler) Create(c echo.Context) error {
	var createUserDTO CreateUserDTO
	if err := c.Bind(createUserDTO); err != nil {
		return util.NewJsonError(c, http.StatusBadRequest, "Invalid body format")
	}

	if err := util.Validate.Struct(createUserDTO); err != nil {
		errors := err.(validator.ValidationErrors)
		return util.NewJsonError(c, http.StatusBadRequest, fmt.Sprintf("Invalid params: %v", errors))
	}

	id, err := h.Service.CreateUser(&createUserDTO)
	if err != nil {
		return util.NewJsonError(c, http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusCreated, map[string]string{"userId": id})
}

func (h *Handler) Update(c echo.Context) error {
	var updateUserDTO UpdateUserDTO
	if err := c.Bind(updateUserDTO); err != nil {
		return util.NewJsonError(c, http.StatusBadRequest, "Invalid body format")
	}

	if err := util.Validate.Struct(updateUserDTO); err != nil {
		errors := err.(validator.ValidationErrors)
		return util.NewJsonError(c, http.StatusBadRequest, fmt.Sprintf("Invalid params: %v", errors))
	}

	err := h.Service.UpdateUser(&updateUserDTO)
	if err != nil {
		return util.NewJsonError(c, http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "user updated successfully",
	})
}
