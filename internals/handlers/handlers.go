package handlers

import (
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
	return &UserHandler{
		db: db,
	}
}

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
	req := models.UserRegisterOrUpdateRequest{}
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
		Message: "user created successfully",
	})
}

func (h *UserHandler) HandleLoginUser(c *fiber.Ctx) error {
	req := models.UserLoginRequest{}
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
		Data:    fiber.Map{"token": tokenStr},
	})
}
