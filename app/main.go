package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gotired/notification-feature/app/config"
	"github.com/gotired/notification-feature/app/database"
	"github.com/gotired/notification-feature/app/handler"
	"github.com/gotired/notification-feature/app/repositories"
	"github.com/gotired/notification-feature/app/services"
)

func main() {
	config := config.Load("./config/config.yaml")
	database := database.NewDatabase(config.Database.URL, config.Database.Name)

	tenantRepo := repositories.NewTenantRepository(database)
	userRepo := repositories.NewUserRepository(database)

	tenantService := services.NewTenantService(tenantRepo)
	userService := services.NewUserService(userRepo)

	tenantHandler := handler.NewTenantHandler(tenantService)
	userHandler := handler.NewUserHandler(userService, tenantService)

	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/api")
	tenantRouter := api.Group("/tenants")
	userRouter := api.Group("/users")

	tenantRouter.Post("", tenantHandler.Create)
	tenantRouter.Get("", tenantHandler.List)
	tenantRouter.Delete(":tenant_id", tenantHandler.Delete)
	tenantRouter.Get(":tenant_id", tenantHandler.Get)
	tenantRouter.Put(":tenant_id", tenantHandler.Update)

	userRouter.Post("", userHandler.Create)
	userRouter.Delete(":user_id", userHandler.Delete)
	userRouter.Get(":user_id", userHandler.Get)
	userRouter.Get("", userHandler.List)
	userRouter.Put(":user_id", userHandler.Update)

	log.Fatal(app.Listen(":3000"))
}
