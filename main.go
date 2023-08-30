package main

import (
	"fmt"
	"log"

	conf "games-js/config"
	ctlr "games-js/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	pathView := fmt.Sprintf("%vviews", conf.GetPathView())
	log.Printf("pathView : %v", pathView)
	// Create a new engine
	engine := html.New(pathView, ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./static")

	//main
	app.Get("/", ctlr.Main)
	app.Get("/main", ctlr.Main)

	//game-1
	app.Get("/game-js-1", ctlr.Game1)

	log.Fatal(app.Listen(":3500"))
}
