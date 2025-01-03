
package api_handlers

import (
	"github.com/MarcelArt/family_tree/models"
	"github.com/MarcelArt/family_tree/repositories"
	"github.com/gofiber/fiber/v2"
)

type ParentHandler struct {
	BaseCrudHandler[models.Parent, models.ParentDTO, models.ParentPage]
}

func NewParentHandler(repo repositories.IParentRepo) *ParentHandler {
	return &ParentHandler{
		BaseCrudHandler: BaseCrudHandler[models.Parent, models.ParentDTO, models.ParentPage]{
			repo: repo,
		},
	}
}

// Create creates a new parent
// @Summary Create a new parent
// @Description Create a new parent
// @Tags Parent
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Parent body models.ParentDTO true "Parent data"
// @Success 201 {object} models.ParentDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /parent [post]
func (h *ParentHandler) Create(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Create(c)
}

// Read retrieves a list of parents
// @Summary Get a list of parents
// @Description Get a list of parents
// @Tags Parent
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param sort query string false "Sort"
// @Param filters query string false "Filter"
// @Success 200 {array} models.ParentPage
// @Router /parent [get]
func (h *ParentHandler) Read(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Read(c)
}

// Update updates an existing parent
// @Summary Update an existing parent
// @Description Update an existing parent
// @Tags Parent
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Parent ID"
// @Param Parent body models.ParentDTO true "Parent data"
// @Success 200 {object} models.ParentDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /parent/{id} [put]
func (h *ParentHandler) Update(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Update(c)
}

// Delete deletes an existing parent
// @Summary Delete an existing parent
// @Description Delete an existing parent
// @Tags Parent
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Parent ID"
// @Success 200 {object} models.Parent
// @Failure 500 {object} string
// @Router /parent/{id} [delete]
func (h *ParentHandler) Delete(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Delete(c)
}

// GetByID retrieves a parent by ID
// @Summary Get a parent by ID
// @Description Get a parent by ID
// @Tags Parent
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Parent ID"
// @Success 200 {object} models.Parent
// @Failure 500 {object} string
// @Router /parent/{id} [get]
func (h *ParentHandler) GetByID(c *fiber.Ctx) error {
	return h.BaseCrudHandler.GetByID(c)
}
