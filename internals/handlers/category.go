package handlers

import (
	"fmt"
	"strings"

	"github.com/assaidy/bookstore/internals/database"
	"github.com/assaidy/bookstore/internals/models"
	"github.com/assaidy/bookstore/internals/utils"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	db *database.DBService
}

func NewCategoryHandler(db *database.DBService) *CategoryHandler {
	return &CategoryHandler{db: db}
}

func (h *CategoryHandler) HandleCreateCategory(c *fiber.Ctx) error {
	req := models.CategoryCreateOrUpdateReq{}
	if err := parseAndValidateReq(c, &req); err != nil {
		return err
	}
	req.Name = strings.TrimSpace(strings.ToLower(req.Name))

	if ok, err := h.db.CheckCategoryConflict(req.Name); err != nil {
		return utils.InternalServerError(err)
	} else if ok {
		return utils.ConflictError(fmt.Sprintf("category %s already exists", req.Name))
	}

	cat := models.Category{Name: req.Name}
	if err := h.db.CreateCategory(&cat); err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusCreated).JSON(utils.ApiResponse{
		Message: "created successfully",
		Data:    fiber.Map{"category": cat},
	})
}

func (h *CategoryHandler) HandleGetAllCategories(c *fiber.Ctx) error {
	cats, err := h.db.GetAllCategories()
	if err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "retrieved successfully",
		Data:    fiber.Map{"categories": cats},
	})
}

func (h *CategoryHandler) HandleUpdateCategoryById(c *fiber.Ctx) error {
	req := models.CategoryCreateOrUpdateReq{}
	if err := parseAndValidateReq(c, &req); err != nil {
		return err
	}
	req.Name = strings.TrimSpace(strings.ToLower(req.Name))

	id, _ := c.ParamsInt("id")

	cat, err := h.db.GetCategoryById(id)
	if err != nil {
	}
	if cat == nil {
		return utils.NotFoundError(fmt.Sprintf("category with id %d not found", id))
	}

	if req.Name != cat.Name {
		if ok, err := h.db.CheckCategoryConflict(req.Name); err != nil {
			return utils.InternalServerError(err)
		} else if ok {
			return utils.ConflictError(fmt.Sprintf("category %s already exists", req.Name))
		}
	}

	cat.Name = req.Name
	if err := h.db.UpdateCategory(cat); err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "updated successfully",
		Data:    fiber.Map{"category": cat},
	})
}

func (h *CategoryHandler) HandleDeleteCategoryById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	if ok, err := h.db.CheckIfCategoryExists(id); err != nil {
		return utils.InternalServerError(err)
	} else if !ok {
		return utils.NotFoundError(fmt.Sprintf("category with id %d not found", id))
	}

	if err := h.db.DeleteCategory(id); err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "deleted successfully",
	})
}
