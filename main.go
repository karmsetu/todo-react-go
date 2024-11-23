package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello world")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error { // Ctx for context
		return c.Status(200).JSON(fiber.Map{"msg": "hello world"})
	})

	log.Fatal(app.Listen(":4000"))
}