package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"hotel-reservations/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5001", "Listen address of the API server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	app.Get("/foo", handleFoo)
	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)

	log.Fatal(app.Listen(*listenAddr)) // asterisk dereferences the pointer
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "working fine!"})
}
