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

	//game-2
	app.Get("/game-js-2", ctlr.Game2)

	//game-3
	app.Get("/game-js-3", ctlr.Game3)

	//game-4
	app.Get("/game-js-4", ctlr.Game4)

	//game-5
	app.Get("/game-js-5", ctlr.Game5)

	log.Fatal(app.Listen(":3500"))
}
