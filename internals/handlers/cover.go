package handlers

import (
	"encoding/base64"
	"fmt"
	"github.com/assaidy/bookstore/internals/database"
	"github.com/assaidy/bookstore/internals/models"
	"github.com/assaidy/bookstore/internals/utils"
	"github.com/gofiber/fiber/v2"
)

type CoverHandler struct {
	db *database.DBService
}

func NewCoverHandler(db *database.DBService) *CoverHandler {
	return &CoverHandler{db: db}
}

// FIX: should not be used, creaet cover in create book handler
func (h *CoverHandler) HandleCreateCover(c *fiber.Ctx) error {
	req := models.CoverCreateOrUpdateReq{}
	if err := parseAndValidateReq(c, &req); err != nil {
		return err
	}

	if ok := utils.CheckEncodingMatchesContent(req.Encoding, req.Content); !ok {
		return utils.InvalidDataError(fmt.Sprintf("content does not match %s encoding", req.Encoding))
	}

	cov := models.Cover{
		Encoding: req.Encoding,
		Content:  req.Content,
	}

	if err := h.db.CreateCover(&cov); err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusCreated).JSON(utils.ApiResponse{
		Message: "created successfully",
		Data:    fiber.Map{"coverId": cov.Id},
	})
}

func (h *CoverHandler) HandleGetCoverById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	cov, err := h.db.GetCoverById(id)
	if err != nil {
		return utils.InternalServerError(err)
	}
	if cov == nil {
		return utils.NotFoundError(fmt.Sprintf("cover with id %d not found", id))
	}

	decodedImage, err := base64.StdEncoding.DecodeString(cov.Content)
	if err != nil {
		return utils.InternalServerError(err)
	}
	c.Set("Content-Type", cov.Encoding)

	return c.Send(decodedImage)

	// return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
	// 	Message: "retrieved successfully",
	// 	Data:    fiber.Map{"cover": cov},
	// })
}

func (h *CoverHandler) HandleUpdateCoverById(c *fiber.Ctx) error {
	req := models.CoverCreateOrUpdateReq{}
	if err := parseAndValidateReq(c, &req); err != nil {
		return err
	}

	if ok := utils.CheckEncodingMatchesContent(req.Encoding, req.Content); !ok {
		return utils.InvalidDataError(fmt.Sprintf("content does not match %s encoding", req.Encoding))
	}

	id, _ := c.ParamsInt("id")

	cov, err := h.db.GetCoverById(id)
	if err != nil {
		return utils.InternalServerError(err)
	}
	if cov == nil {
		return utils.NotFoundError(fmt.Sprintf("cover with id %d not found", id))
	}

	cov.Encoding = req.Encoding
	cov.Content = req.Content

	if err := h.db.UpdateCover(cov); err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "updated successfully",
	})
}

// FIX: should not be used, creaet cover in delete book handler
func (h *CoverHandler) HandleDeleteCoverById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	if ok, err := h.db.CheckIfCoverExists(id); err != nil {
		return utils.InternalServerError(err)
	} else if !ok {
		return utils.NotFoundError(fmt.Sprintf("cover with id %d not found", id))
	}

	if err := h.db.DeleteCover(id); err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "deleted successfully",
	})
}
