package http

import (
	"belajar-golang/internal/entity"
	"belajar-golang/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// UserHandler represent the httphandler for user
type UserHandler struct {
	userUsecase usecase.UserUsecase
}

// NewUserHandler will initialize the user/ resources endpoint
func NewUserHandler(f *fiber.App, userUsecase usecase.UserUsecase) {
	handler := UserHandler{userUsecase}

	api := f.Group("/api")
	v1 := api.Group("/v1")

	userGroup := v1.Group("/users")
	userGroup.Post("/", handler.RegisterUser)
	userGroup.Get("/", handler.FindAllUser)
	userGroup.Get("/:id", handler.FindUserByID)
	userGroup.Get("/username/:username", handler.FindUserByUsername)
	userGroup.Put("/:id", handler.UpdateUser)
	userGroup.Delete("/:id", handler.DeleteUser)
}

// RegisterUser will create a new user
func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	var user entity.User

	// catch error if the body is not valid
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	// validate input with go validator
	if err := user.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	// check if username already exists
	existingUserName, err := h.userUsecase.FindUserByUsername(user.Username)
	if existingUserName.ID != 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  "error",
			"message": "username already exists",
		})
	}

	// check if email already exists
	existingEmail, err := h.userUsecase.FindUserByUsername(user.Email)
	if existingEmail.ID != 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  "error",
			"message": "email already exists",
		})
	}

	// hash password before save to database
	hasPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	user.Password = string(hasPassword)

	// create new user
	result, err := h.userUsecase.RegisterUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   result,
	})
}

// FindAllUser will return all user
func (h *UserHandler) FindAllUser(c *fiber.Ctx) error {
	result, err := h.userUsecase.FindAllUser()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   result,
	})
}

// FindUserByID will find user by given id
func (h *UserHandler) FindUserByID(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	result, err := h.userUsecase.FindUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   result,
	})
}

// FindUserByUsername will find user by given username
func (h *UserHandler) FindUserByUsername(c *fiber.Ctx) error {
	username := c.Params("username")

	result, err := h.userUsecase.FindUserByUsername(username)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   result,
	})
}

// UpdateUser will update user by given id
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	var user entity.User

	// catch error if the body is not valid
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	// validate input with go validator
	if err := user.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	// check if user exists
	_, err = h.userUsecase.FindUserByID(userID)
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": "user not found",
			})
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	// hash password if password is not empty
	if user.Password != "" {
		hasPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		}

		user.Password = string(hasPassword)

	}

	// update user
	result, err := h.userUsecase.UpdateUser(userID, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   result,
	})
}

// DeleteUser will delete user by given id
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	// check if user exists
	_, err = h.userUsecase.FindUserByID(userID)
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": "user not found",
			})
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	err = h.userUsecase.DeleteUser(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "user deleted",
	})
}
