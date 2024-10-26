package handlers

import (
	"fmt"

	"github.com/assaidy/bookstore/internals/database"
	"github.com/assaidy/bookstore/internals/utils"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	db *database.DBService
}

func NewOrderHandler(db *database.DBService) *OrderHandler {
	return &OrderHandler{db: db}
}

func (h *OrderHandler) HandleApplyOrder(c *fiber.Ctx) error {
	// TODO: handle third party shipment
	// TODO: handle third party payment

	uid, _ := c.ParamsInt("uid")

	if ok, err := h.db.CheckIfUserExists(uid); err != nil {
		return utils.InternalServerError(err)
	} else if !ok {
		return utils.NotFoundError(fmt.Sprintf("user with id %d not found", uid))
	}

	if err := h.db.MakeOrder(uid); err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "ordered successfully",
	})
}

func (h *OrderHandler) HandleGetAllOrderByUser(c *fiber.Ctx) error {
	uid, _ := c.ParamsInt("uid")

	if ok, err := h.db.CheckIfUserExists(uid); err != nil {
		return utils.InternalServerError(err)
	} else if !ok {
		return utils.NotFoundError(fmt.Sprintf("user with id %d not found", uid))
	}

	orders, err := h.db.GetAllOrdersByUser(uid)
	if err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "retrieved successfully",
		Data:    fiber.Map{"orders": orders},
	})
}

func (h *OrderHandler) HandleGetAllOrders(c *fiber.Ctx) error {
	orders, err := h.db.GetAllOrders()
	if err != nil {
		return utils.InternalServerError(err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "retrieved successfully",
		Data:    fiber.Map{"orders": orders},
	})
}

func (h *OrderHandler) HandleGetOrderById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	order, err := h.db.GetOrderById(id)
	if err != nil {
		return utils.InternalServerError(err)
	}
	if order == nil {
		return utils.NotFoundError(fmt.Sprintf("order with id %d not found", id))
	}

	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Message: "retrieved successfully",
		Data:    fiber.Map{"order": order},
	})

}
