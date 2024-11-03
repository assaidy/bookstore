package handlers

import (
	"fmt"
	"time"

	"github.com/assaidy/bookstore/internals/database"
	"github.com/assaidy/bookstore/internals/models"
	"github.com/assaidy/bookstore/internals/utils"
	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	db *database.DBService
}

func NewBookHandler(db *database.DBService) *BookHandler {
	return &BookHandler{db: db}
}

func (h *BookHandler) HandleCreateBook(c *fiber.Ctx) error {
	req := models.BookCreateRequest{}
	if err := parseAndValidateReq(c, &req); err != nil {
		return err
	}

	book := models.Book{
		Title:       req.Title,
		Description: req.Description,
		CategoryId:  req.CategoryId,
		CoverId:     req.CoverId,
		Price:       req.Price,
		Quantity:    req.Quantity,
		Discount:    req.Discount,
		AddedAt:     time.Now().UTC(),
	}

	if err := h.db.CreateBook(&book); err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusCreated).JSON(utils.ApiResponse{
		Message: "created successfully",
	})
}

func getSortingTechnique(c *fiber.Ctx) (string, error) {
	st := c.Query("sorting")
	if st == "" {
		return "", fmt.Errorf("plz set value for param 'sorting'")
	}

	if st == "popularity" || st == "latest" || st == "price_asc" || st == "price_desc" {
		return st, nil
	}

	return "", fmt.Errorf("'sorting' param takes only values {popularity, latest, price_asc, price_desc}")
}

func getPaginationData(c *fiber.Ctx) (page, limit int) {
	defaultPage := 1
	defaultLimit := 8
	page = c.QueryInt("page", defaultPage)
	limit = c.QueryInt("limit", defaultLimit)

	if page < 1 {
		page = defaultPage
	}

	if limit < 1 {
		limit = defaultLimit
	} else if limit > 100 {
		limit = 100
	}

	return page, limit
}

func (h *BookHandler) HandleGetAllBooks(c *fiber.Ctx) error {
	sorting, err := getSortingTechnique(c)
	if err != nil {
		return utils.BadRequestError(err.Error())
	}
	page, limit := getPaginationData(c)

	books, err := h.db.GetAllBooks(sorting, page, limit)
	if err != nil {
		return utils.InternalServerError(err)
	}

	totalBooks, err := h.db.GetTotalBooks()
	if err != nil {
		return utils.InternalServerError(err)
	}
	totalPages := (totalBooks + limit - 1) / limit

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "retrieved successfully",
		Data: fiber.Map{
			"books":      books,
			"page":       page,
			"limit":      limit,
			"totalPages": totalPages,
		},
	})
}

func (h *BookHandler) HnadleGetBookById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	book, err := h.db.GetBookById(id)
	if err != nil {
		return utils.InternalServerError(err)
	}
	if book == nil {
		return utils.NotFoundError(fmt.Sprintf("book with id %d not found", id))
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "retrieved successfully",
		Data:    fiber.Map{"book": book},
	})
}

func (h *BookHandler) HnadleUpdateBookById(c *fiber.Ctx) error {
	req := models.BookUpdateRequest{}
	if err := parseAndValidateReq(c, &req); err != nil {
		return err
	}

	id, _ := c.ParamsInt("id")

	book, err := h.db.GetBookById(id)
	if err != nil {
		return utils.InternalServerError(err)
	}
	if book == nil {
		return utils.NotFoundError(fmt.Sprintf("book with id %d not found", id))
	}

	book.Title = req.Title
	book.Description = req.Description
	book.CategoryId = req.CategoryId
	book.Price = req.Price
	book.Quantity = req.Quantity
	book.Discount = req.Discount

	if err := h.db.UpdateBook(book); err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "deleted successfully",
	})
}

func (h *BookHandler) HnadleDeleteBookById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	if ok, err := h.db.CheckIfBookExists(id); err != nil {
		return utils.InternalServerError(err)
	} else if !ok {
		return utils.NotFoundError(fmt.Sprintf("book with id %d not found", id))
	}

	if err := h.db.DeleteBook(id); err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "deleted successfully",
	})
}
