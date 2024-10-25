package handlers

import (
	"fmt"

	"github.com/assaidy/bookstore/internals/database"
	"github.com/assaidy/bookstore/internals/utils"
	"github.com/gofiber/fiber/v2"
)

type FavouritesHandler struct {
	db *database.DBService
}

func NewFavouritesHandler(db *database.DBService) *FavouritesHandler {
	return &FavouritesHandler{db: db}
}

func (h *FavouritesHandler) HandleAddBookToFavourites(c *fiber.Ctx) error {
	uid, _ := c.ParamsInt("uid")
	bid, _ := c.ParamsInt("bid")

	if ok, err := h.db.CheckIfBookExists(bid); err != nil {
		return utils.InternalServerError(err)
	} else if !ok {
		return utils.NotFoundError(fmt.Sprintf("book with id %d not found", bid))
	}

	if err := h.db.AddBookToFavourites(uid, bid); err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusCreated).JSON(utils.ApiResponse{
		Message: "created successfully",
	})
}

func (h *FavouritesHandler) HandleGetAllUserFavourites(c *fiber.Ctx) error {
	uid, _ := c.ParamsInt("uid")
	books, err := h.db.GetAllBooksInFavourites(uid)
	if err != nil {
		return utils.InternalServerError(err)
	}
	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "retrieved successfully",
		Data:    fiber.Map{"books": books},
	})
}

func (h *FavouritesHandler) HandleDeleteBookFromFavourites(c *fiber.Ctx) error {
	uid, _ := c.ParamsInt("uid")
	bid, _ := c.ParamsInt("bid")

	if ok, err := h.db.CheckIfBookExists(bid); err != nil {
		return utils.InternalServerError(err)
	} else if !ok {
		return utils.NotFoundError(fmt.Sprintf("book with id %d not found", bid))
	}

	if err := h.db.DeleteBookFromFavourites(uid, bid); err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "deleted successfully",
	})
}
