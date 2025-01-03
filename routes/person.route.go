
package routes

import (
	"github.com/MarcelArt/family_tree/database"
	api_handlers "github.com/MarcelArt/family_tree/handlers/api"
	"github.com/MarcelArt/family_tree/repositories"
	"github.com/gofiber/fiber/v2"
)

func SetupPersonRoutes(api fiber.Router) {
	h := api_handlers.NewPersonHandler(repositories.NewPersonRepo(database.GetDB()))

	g := api.Group("/person")
	g.Get("/", h.Read)
	g.Get("/:id", h.GetByID)
	g.Post("/", h.Create)
	g.Put("/:id", h.Update)
	g.Delete("/:id", h.Delete)
}
