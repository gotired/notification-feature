package repositories

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/gotired/notification-feature/app/model"
	"github.com/gotired/notification-feature/app/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type tenantRepo struct {
	collection *mongo.Collection
}

func NewTenantRepository(database *mongo.Database) model.TenantRepository {
	collection := database.Collection("tenants")
	return &tenantRepo{collection}
}

func (r *tenantRepo) Check(name string) (*model.Tenant[bson.Binary], error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res := r.collection.FindOne(ctx, bson.M{"name": name})

	var response model.Tenant[bson.Binary]
	if err := res.Decode(&response); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &response, nil
}
func (r *tenantRepo) CheckByID(id bson.Binary) (*model.Tenant[bson.Binary], error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res := r.collection.FindOne(ctx, bson.M{"_id": id})

	var response model.Tenant[bson.Binary]
	if err := res.Decode(&response); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &response, nil
}

func (r *tenantRepo) Insert(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	insertDoc := model.Tenant[bson.Binary]{
		ID:        utils.ConvertUUIDToBinary(uuid.New()),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err := r.collection.InsertOne(ctx, insertDoc)
	if err != nil {
		return err
	}
	return nil
}

func (r *tenantRepo) Detail(id bson.Binary) (*model.Tenant[bson.Binary], error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res := r.collection.FindOne(ctx, bson.M{"_id": id})

	var response model.Tenant[bson.Binary]
	if err := res.Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (r *tenantRepo) List(limit, page int, keyword, order, orderKey string) ([]model.Tenant[bson.Binary], error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if limit == 0 {
		limit = 10
	}
	if page == 0 {
		page = 1
	}
	if orderKey == "" {
		orderKey = "updated_at"
	}
	direction := 1
	if order == "desc" {
		direction = -1
	}

	var pipeline []bson.M

	if keyword != "" {
		pipeline = append(pipeline, bson.M{
			"$match": bson.M{
				"name": bson.M{
					"$regex":   keyword,
					"$options": 'i',
				},
			},
		})
	}
	pipeline = append(
		pipeline,
		bson.M{
			"$sort": bson.M{
				orderKey: direction,
			},
		},
		bson.M{
			"$skip": (page - 1) * limit,
		},
		bson.M{
			"$limit": limit,
		},
	)

	res, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var response []model.Tenant[bson.Binary]
	if err = res.All(ctx, &response); err != nil {
		return nil, err
	}
	return response, nil

}

func (r *tenantRepo) Update(id bson.Binary, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{
			"$set": bson.M{"name": name, "updated_at": time.Now()},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *tenantRepo) Delete(id bson.Binary) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil

}
