package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/gotired/notification-feature/worker/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type NotificationService struct {
	repo model.NotificationRepo
}

func NewNotificationService(repo model.NotificationRepo) *NotificationService {
	return &NotificationService{repo: repo}
}

func (s *NotificationService) NotifyAll(tenantID uuid.UUID, message string, createdAt time.Time) error {
	tenantIDBson := bson.Binary{
		Subtype: 4,
		Data:    tenantID[:],
	}

	userIDs, err := s.repo.GetUserIDsByTenantID(tenantIDBson)
	if err != nil {
		return err
	}

	notifications := make([]model.Notification[bson.Binary], len(userIDs))

	for i, userID := range userIDs {
		notiID := uuid.New()
		notifications[i] = model.Notification[bson.Binary]{
			ID: bson.Binary{
				Subtype: 4,
				Data:    notiID[:],
			},
			TenantID:  tenantIDBson,
			UserID:    userID,
			Message:   message,
			CreatedAt: createdAt,
		}
	}
	return s.repo.StoreNotifications(notifications)
}
