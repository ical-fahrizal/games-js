package controller

import "github.com/gofiber/fiber/v2"

var Main = func(c *fiber.Ctx) error {
	// Render index within layouts/main
	return c.Render("main", fiber.Map{
		"Title": "Games Js",
	})
}

var Game1 = func(c *fiber.Ctx) error {
	return c.Render("game-js-1", fiber.Map{
		"Title": "Games Js",
	})
}
