package repositories

import (
	"context"
	"log"

	"github.com/gotired/notification-feature/worker/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoNotificationRepo struct {
	database *mongo.Database
}

func NewMongoNotificationRepo(database *mongo.Database) model.NotificationRepo {
	return &MongoNotificationRepo{database}
}

func (r *MongoNotificationRepo) GetUserIDsByTenantID(id bson.Binary) ([]bson.Binary, error) {
	collection := r.database.Collection("users")

	ctx := context.Background()

	res, err := collection.Aggregate(ctx, bson.A{
		bson.M{
			"$match": bson.M{
				"tenant": id,
			},
		},
		bson.M{
			"$project": bson.M{
				"_id": 1,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	listUsers := make([]bson.Binary, 0)
	for res.Next(context.TODO()) {
		var user struct {
			ID bson.Binary `bson:"_id"`
		}
		if err := res.Decode(&user); err != nil {
			log.Fatal(err)
		}
		listUsers = append(listUsers, user.ID)
	}

	return listUsers, nil
}

func (r *MongoNotificationRepo) StoreNotifications(notifications []model.Notification[bson.Binary]) error {
	collection := r.database.Collection("notifications")
	ctx := context.Background()

	docs := make([]any, len(notifications))
	for i, n := range notifications {
		docs[i] = bson.M{
			"tenant_id":  n.TenantID,
			"user_id":    n.UserID,
			"message":    n.Message,
			"created_at": n.CreatedAt,
		}
	}
	_, err := collection.InsertMany(ctx, docs)
	return err
}
