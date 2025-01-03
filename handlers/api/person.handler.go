
package api_handlers

import (
	"github.com/MarcelArt/family_tree/models"
	"github.com/MarcelArt/family_tree/repositories"
	"github.com/gofiber/fiber/v2"
)

type PersonHandler struct {
	BaseCrudHandler[models.Person, models.PersonDTO, models.PersonPage]
}

func NewPersonHandler(repo repositories.IPersonRepo) *PersonHandler {
	return &PersonHandler{
		BaseCrudHandler: BaseCrudHandler[models.Person, models.PersonDTO, models.PersonPage]{
			repo: repo,
		},
	}
}

// Create creates a new person
// @Summary Create a new person
// @Description Create a new person
// @Tags Person
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Person body models.PersonDTO true "Person data"
// @Success 201 {object} models.PersonDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /person [post]
func (h *PersonHandler) Create(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Create(c)
}

// Read retrieves a list of persons
// @Summary Get a list of persons
// @Description Get a list of persons
// @Tags Person
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param sort query string false "Sort"
// @Param filters query string false "Filter"
// @Success 200 {array} models.PersonPage
// @Router /person [get]
func (h *PersonHandler) Read(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Read(c)
}

// Update updates an existing person
// @Summary Update an existing person
// @Description Update an existing person
// @Tags Person
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Person ID"
// @Param Person body models.PersonDTO true "Person data"
// @Success 200 {object} models.PersonDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /person/{id} [put]
func (h *PersonHandler) Update(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Update(c)
}

// Delete deletes an existing person
// @Summary Delete an existing person
// @Description Delete an existing person
// @Tags Person
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Person ID"
// @Success 200 {object} models.Person
// @Failure 500 {object} string
// @Router /person/{id} [delete]
func (h *PersonHandler) Delete(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Delete(c)
}

// GetByID retrieves a person by ID
// @Summary Get a person by ID
// @Description Get a person by ID
// @Tags Person
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Person ID"
// @Success 200 {object} models.Person
// @Failure 500 {object} string
// @Router /person/{id} [get]
func (h *PersonHandler) GetByID(c *fiber.Ctx) error {
	return h.BaseCrudHandler.GetByID(c)
}
