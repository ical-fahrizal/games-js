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

var Game2 = func(c *fiber.Ctx) error {
	return c.Render("game-js-2", fiber.Map{
		"Title": "Games Js",
	})
}

var Game3 = func(c *fiber.Ctx) error {
	return c.Render("game-js-3", fiber.Map{
		"Title": "Games Js",
	})
}

var Game4 = func(c *fiber.Ctx) error {
	return c.Render("game-js-4", fiber.Map{
		"Title": "Games Js",
	})
}

var Game5 = func(c *fiber.Ctx) error {
	return c.Render("game-js-5", fiber.Map{
		"Title": "Games Js",
	})
}
