package handler

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gotired/notification-feature/worker/model"
	"github.com/gotired/notification-feature/worker/services"
	"github.com/segmentio/kafka-go"
)

func StartConsumer(service *services.NotificationService, broker, topic, groupID string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: groupID,
	})
	defer reader.Close()
	log.Println("Starting consume message")

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("error reading message: %v", err)
			continue
		}

		var payload model.NotificationPayload
		if err := json.Unmarshal(m.Value, &payload); err != nil {
			log.Printf("invalid payload: %v", err)
			continue
		}
		tenantID, err := uuid.Parse(payload.TenantID)
		if err != nil {
			log.Printf("invalid tenant: %v", err)
			continue
		}
		err = service.NotifyAll(tenantID, payload.Message, time.Now())
		if err != nil {
			log.Printf("failed to process notification: %v", err)
		}
	}
}
