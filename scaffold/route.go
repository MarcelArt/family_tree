package scaffold

import (
	"log"
	"os"
	"strings"
)

const routeTemplate = `
package routes

import (
	"${moduleName}/database"
	api_handlers "github.com/MarcelArt/family_tree/handlers/api"
	"github.com/MarcelArt/family_tree/repositories"
	"github.com/gofiber/fiber/v2"
)

func Setup${modelName}Routes(api fiber.Router) {
	h := api_handlers.New${modelName}Handler(repositories.New${modelName}Repo(database.GetDB()))

	g := api.Group("/${handlerRoute}")
	g.Get("/", h.Read)
	g.Get("/:id", h.GetByID)
	g.Post("/", h.Create)
	g.Put("/:id", h.Update)
	g.Delete("/:id", h.Delete)
}
`

func ScaffoldRoute(modelName string, handlerRoute string) {
	filename := "routes/" + ToSeparateByCharLowered(modelName, '_') + ".route.go"
	log.Printf("Generating route file: %s", filename)

	newRoute := strings.ReplaceAll(routeTemplate, "${modelName}", modelName)
	newRoute = strings.ReplaceAll(newRoute, "${handlerRoute}", handlerRoute)
	newRoute = strings.ReplaceAll(newRoute, "${moduleName}", moduleName)

	if err := os.WriteFile(filename, []byte(newRoute), 0644); err != nil {
		panic("Failed writing route file")
	}
}
