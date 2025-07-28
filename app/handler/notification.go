package handler

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gofiber/fiber/v2"
	"github.com/gotired/notification-feature/app/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type NotificationHandler struct {
	producer *kafka.Producer
	database *mongo.Database
}

func NewNotificationHandler(producer *kafka.Producer, database *mongo.Database) NotificationHandler {
	return NotificationHandler{producer, database}
}

func (h *NotificationHandler) CreateAlert(ctx *fiber.Ctx) error {
	var job model.NotificationBody
	if err := ctx.BodyParser(&job); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}

	payload, _ := json.Marshal(job)
	err := h.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &[]string{"notifications"}[0], Partition: kafka.PartitionAny},
		Value:          payload,
	}, nil)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": fmt.Sprintf("failed to produce message : %v", err.Error())})
	}

	return ctx.JSON(fiber.Map{"status": "notification job queued"})
}
