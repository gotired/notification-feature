package main

import (
	"github.com/gotired/notification-feature/worker/config"
	"github.com/gotired/notification-feature/worker/database"
	"github.com/gotired/notification-feature/worker/handler"
	"github.com/gotired/notification-feature/worker/repositories"
	"github.com/gotired/notification-feature/worker/services"
)

func main() {
	config := config.Load("./config/config.yaml")
	db := database.NewDatabase(config.Database.URL, config.Database.Name)

	notiRepo := repositories.NewMongoNotificationRepo(db)
	notiService := services.NewNotificationService(notiRepo)

	handler.StartConsumer(notiService, config.Kafka.Server, "notifications", config.Kafka.Group)
}
