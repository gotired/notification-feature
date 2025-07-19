package services

import (
	"github.com/google/uuid"
	"github.com/gotired/notification-feature/app/model"
	"github.com/gotired/notification-feature/app/utils"
)

type userService struct {
	repo model.UserRepository
}

func NewUserService(repo model.UserRepository) model.UserService {
	return &userService{repo}
}

func (r *userService) Check(name string) (*model.User[uuid.UUID], error) {
	res, err := r.repo.Check(name)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}

	userID, err := utils.ConvertBinarytoUUID(res.ID)
	if err != nil {
		return nil, err
	}

	tenantID, err := utils.ConvertBinarytoUUID(res.Tenant)
	if err != nil {
		return nil, err
	}

	return &model.User[uuid.UUID]{
		ID:        userID,
		Name:      res.Name,
		Tenant:    tenantID,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (r *userService) Insert(name string, tenant uuid.UUID) error {
	return r.repo.Insert(name, utils.ConvertUUIDToBinary(tenant))
}

func (r *userService) Detail(id uuid.UUID) (*model.User[uuid.UUID], error) {
	res, err := r.repo.Detail(utils.ConvertUUIDToBinary(id))
	if err != nil {
		return nil, err
	}

	userID, err := utils.ConvertBinarytoUUID(res.ID)
	if err != nil {
		return nil, err
	}

	tenantID, err := utils.ConvertBinarytoUUID(res.Tenant)
	if err != nil {
		return nil, err
	}

	return &model.User[uuid.UUID]{
		ID:        userID,
		Name:      res.Name,
		Tenant:    tenantID,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (r *userService) List(limit, page int, keyword, order, orderKey string) ([]model.User[uuid.UUID], error) {

	res, err := r.repo.List(limit, page, keyword, order, orderKey)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}

	users := make([]model.User[uuid.UUID], 0, len(res))
	for _, ub := range res {
		u, err := uuid.FromBytes(ub.ID.Data)
		if err != nil {
			return nil, err
		}
		t, err := uuid.FromBytes(ub.Tenant.Data)
		if err != nil {
			return nil, err
		}
		users = append(users, model.User[uuid.UUID]{
			ID:        u,
			Name:      ub.Name,
			Tenant:    t,
			CreatedAt: ub.CreatedAt,
			UpdatedAt: ub.UpdatedAt,
		})
	}
	return users, nil
}

func (r *userService) Update(id uuid.UUID, name string) error {

	return r.repo.Update(utils.ConvertUUIDToBinary(id), name)
}

func (r *userService) Delete(id uuid.UUID) error {
	return r.repo.Delete(utils.ConvertUUIDToBinary(id))
}
