package handlers

import (
	"fmt"

	"github.com/assaidy/bookstore/internals/database"
	"github.com/assaidy/bookstore/internals/models"
	"github.com/assaidy/bookstore/internals/utils"
	"github.com/gofiber/fiber/v2"
)

type CartHandler struct {
	db *database.DBService
}

func NewCartHandler(db *database.DBService) *CartHandler {
	return &CartHandler{db: db}
}

func (h *CartHandler) HandleAddToCart(c *fiber.Ctx) error {
	uid, _ := c.ParamsInt("uid")

	req := models.CartAddBookReq{}
	if err := parseAndValidateReq(c, &req); err != nil {
		return err
	}

	book, err := h.db.GetBookById(req.BookId)
	if err != nil {
		return err
	}
	if book == nil {
		return utils.NotFoundError(fmt.Sprintf("book with id %d not found", req.BookId))
	}

	if book.Quantity < req.Quantity {
		return utils.InvalidDataError("given quantity is greated than book quantity")
	}

	if err := h.db.AddBookToCart(uid, req.BookId, req.Quantity); err != nil {
		return utils.InternalServerError(err)
	}

	book.Quantity -= req.Quantity
	if err := h.db.UpdateBook(book); err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusCreated).JSON(utils.ApiResponse{
		Message: "created successfully",
	})
}

func (h *CartHandler) HandleGetBooksInCart(c *fiber.Ctx) error {
	uid, _ := c.ParamsInt("uid")
	books, err := h.db.GetBooksInCart(uid)
	if err != nil {
		return utils.InternalServerError(err)
	}
	total := getTotalPrice(books)
	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "retrieved successfully",
		Data:    fiber.Map{"books": books, "total": total},
	})
}

func (h *CartHandler) HandleDeleteBookFromCart(c *fiber.Ctx) error {
	uid, _ := c.ParamsInt("uid")
	bid, _ := c.ParamsInt("bid")

    cartBook, err := h.db.GetBookFromCart(uid, bid)
	if err != nil {
		return err
	}
	if cartBook == nil {
		return utils.NotFoundError(fmt.Sprintf("book with id %d not found in cart", bid))
	}

    if err := h.db.DeleteBookFromCart(uid, bid); err != nil {
        return err
    }

    book, err := h.db.GetBookById(bid)
    if err != nil {
        return err
    }
    if book == nil {
        return utils.NotFoundError(fmt.Sprintf("book with id %d not found", bid))
    }

	book.Quantity += cartBook.Quantity
	if err := h.db.UpdateBook(book); err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "deleted successfully",
	})
}

func getTotalPrice(books []*models.CartBook) float64 {
	total := 0.0
	for _, book := range books {
		total += book.PricePerUnite
	}
	return total
}
