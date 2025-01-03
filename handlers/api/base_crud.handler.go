package api_handlers

import (
	"github.com/MarcelArt/family_tree/models"
	"github.com/MarcelArt/family_tree/repositories"
	"github.com/gofiber/fiber/v2"
)

type BaseCrudHandler[TModel any, TDto models.IDTO, TPage any] struct {
	repo repositories.IBaseCrudRepo[TModel, TDto, TPage]
}

func (h *BaseCrudHandler[TModel, TDto, TPage]) Create(c *fiber.Ctx) error {
	var dto TDto

	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	id, err := h.repo.Create(dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"ID": id})
}

// Read retrieves a list of resources
// @Summary Get a list of resources
// @Description Get a list of resources
// @Tags CRUDs
// @Accept json
// @Produce json
// @Param resources path string true "Resource route" Enums(template, process)
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param sort query string false "Sort"
// @Param filters query string false "Filter"
// @Success 200 {array} interface{}
// @Router /{resources} [get]
func (h *BaseCrudHandler[TModel, TDto, TPage]) Read(c *fiber.Ctx) error {
	var dest []TPage

	page := h.repo.Read(c, dest)

	return c.Status(fiber.StatusOK).JSON(page)
}

func (h *BaseCrudHandler[TModel, TDto, TPage]) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var dto TDto

	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := h.repo.Update(id, &dto); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(dto)
}

func (h *BaseCrudHandler[TModel, TDto, TPage]) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	model, err := h.repo.Delete(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(model)
}

func (h *BaseCrudHandler[TModel, TDto, TPage]) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	model, err := h.repo.GetByID(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(model)
}
