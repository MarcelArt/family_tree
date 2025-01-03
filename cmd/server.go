package cmd

import (
	"fmt"
	"log"

	"github.com/MarcelArt/family_tree/config"
	"github.com/MarcelArt/family_tree/database"
	"github.com/MarcelArt/family_tree/routes"
	"github.com/gofiber/fiber/v2"
)

func Serve() {
	database.ConnectDB()

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Printf("Listening to http://localhost:%s", config.Env.PORT)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Env.PORT)))
}
