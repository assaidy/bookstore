package handlers

import (
	"fmt"
	"time"

	"github.com/assaidy/bookstore/internals/database"
	"github.com/assaidy/bookstore/internals/models"
	"github.com/assaidy/bookstore/internals/utils"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	db *database.DBService
}

func NewUserHandler(db *database.DBService) *UserHandler {
	return &UserHandler{db: db}
}

// TODO: trim spaces

func parseAndValidateReq(c *fiber.Ctx, out any) error {
	if err := c.BodyParser(&out); err != nil {
		return utils.InvalidJsonRequestError()
	}
	if errs := utils.ValidateRequest(out); errs != nil {
		return utils.ValidationError(errs)
	}
	return nil
}

func (h *UserHandler) HandleRegisterUser(c *fiber.Ctx) error {
	req := models.UserRegisterOrUpdateReq{}
	if err := parseAndValidateReq(c, &req); err != nil {
		return err
	}

	if ok, err := h.db.CheckUsernameAndEmailConflict(req.Username, req.Email); err != nil {
		return utils.InternalServerError(err)
	} else if ok {
		return utils.ConflictError("username or email already exists")
	}

	hashedPassword, err := utils.HashPassword([]byte(req.Password))
	if err != nil {
		return utils.InternalServerError(err)
	}
	user := models.User{
		Name:     req.Name,
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
		Address:  req.Address,
		JoinedAt: time.Now().UTC(),
	}
	if err := h.db.CreateUser(&user); err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusCreated).JSON(utils.ApiResponse{
		Message: "created successfully",
		Data:    fiber.Map{"user": user},
	})
}

func (h *UserHandler) HandleLoginUser(c *fiber.Ctx) error {
	req := models.UserLoginReq{}
	if err := parseAndValidateReq(c, &req); err != nil {
		return err
	}

	user, err := h.db.GetUserByUsername(req.Username)
	if err != nil {
		return utils.InternalServerError(err)
	}
	if user == nil || !utils.VerifyPasswrod([]byte(req.Password), []byte(user.Password)) {
		return utils.UnauthorizedError()
	}

	tokenStr, err := utils.GenerateJwtToken(user.Id, user.Username)
	if err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "logged in successfully",
		Data:    fiber.Map{"token": tokenStr, "user": user},
	})
}

func (h *UserHandler) HandleGetAllUsers(c *fiber.Ctx) error {
	users, err := h.db.GetAllUsers()
	if err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "retrieved successfully",
		Data:    fiber.Map{"users": users},
	})
}

func (h *UserHandler) HandleGetUserById(c *fiber.Ctx) error {
	// id, ok := utils.GetUserIdFromContext(c)
	// if !ok {
	// 	return utils.UnauthorizedError()
	// }
	id, _ := c.ParamsInt("id")

	user, err := h.db.GetUserById(id)
	if err != nil {
		return utils.InternalServerError(err)
	}
	if user == nil {
		return utils.NotFoundError(fmt.Sprintf("user with id %d not found", id))
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "user found",
		Data:    fiber.Map{"user": user},
	})
}

func (h *UserHandler) HandleUpdateUserById(c *fiber.Ctx) error {
	req := models.UserRegisterOrUpdateReq{}
	if err := parseAndValidateReq(c, &req); err != nil {
		return err
	}

	id, _ := c.ParamsInt("id")

	user, err := h.db.GetUserById(id)
	if err != nil {
		return utils.InternalServerError(err)
	}
	if user == nil {
		return utils.NotFoundError(fmt.Sprintf("user with id %d not found", id))
	}

	if req.Username != user.Username {
		if ok, err := h.db.CheckUsernameConflict(req.Username); err != nil {
			return utils.InternalServerError(err)
		} else if ok {
			return utils.ConflictError("username already exists")
		}
	}
	if req.Email != user.Email {
		if ok, err := h.db.CheckEmailConflict(req.Email); err != nil {
			return utils.InternalServerError(err)
		} else if ok {
			return utils.ConflictError("email already exists")
		}
	}

	hashedPassword, err := utils.HashPassword([]byte(req.Password))
	if err != nil {
		return utils.InternalServerError(err)
	}
	user.Password = hashedPassword
	user.Username = req.Username
	user.Name = req.Name
	user.Email = req.Email
	user.Address = req.Address

	if err := h.db.UpdateUser(user); err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "updated successfully",
		Data:    fiber.Map{"user": user},
	})
}

func (h *UserHandler) HandleDeleteUserById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	if ok, err := h.db.CheckIfUserExists(id); err != nil {
		return utils.InternalServerError(err)
	} else if !ok {
		return utils.UnauthorizedError()
	}

	if err := h.db.DeleteUser(id); err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "deleted successfully",
	})
}
