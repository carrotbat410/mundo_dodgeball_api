package routes

import (
	"fiber_prac/controllers"

	"github.com/gofiber/fiber/v2"
)

// api := api.Group("/v1", middleware) 필요시 middleware + c.Next() 사용 가능

// app.Get("/", func(c *fiber.Ctx) error {
// 	return c.SendString("hello world")
// })

// app.Get("/sample/error", func(c *fiber.Ctx) error {
// 	return fiber.NewError(782, "Custom Error Message")
// })

// app.Get("/sample/:name", func(c *fiber.Ctx) error {
// 	return c.SendString("name: " + c.Params("name"))
// })

// AuthRoutes는 /auth 그룹 경로를 설정합니다.
func AuthRoutes(app *fiber.App) {
	authGroup := app.Group("/auth")

	authGroup.Post("/register", controllers.RegisterUser)
	authGroup.Post("/login", controllers.Login)
}
