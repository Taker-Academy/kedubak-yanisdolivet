package main

import (
	"context"

	"example/kedubak-yanisdolivet/connect_db"
	"example/kedubak-yanisdolivet/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	Client := connect_db.ConnectDb()
	defer func() {
		if err := Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	app.Use(cors.New())
	controllers.GetPostRequest(app, Client)
	app.Listen(":8080")
}
