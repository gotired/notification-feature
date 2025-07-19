package services

import (
	"github.com/google/uuid"
	"github.com/gotired/notification-feature/app/model"
	"github.com/gotired/notification-feature/app/utils"
)

type tenantService struct {
	repo model.TenantRepository
}

func NewTenantService(repo model.TenantRepository) model.TenantService {
	return &tenantService{repo}
}

func (r *tenantService) Check(name string) (*model.Tenant[uuid.UUID], error) {
	res, err := r.repo.Check(name)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	id, _ := utils.ConvertBinarytoUUID(res.ID)
	return &model.Tenant[uuid.UUID]{
		Name: res.Name,
		ID:   id,
	}, nil
}

func (r *tenantService) CheckByID(id uuid.UUID) (*model.Tenant[uuid.UUID], error) {
	res, err := r.repo.CheckByID(utils.ConvertUUIDToBinary(id))
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	uuidID, _ := utils.ConvertBinarytoUUID(res.ID)
	return &model.Tenant[uuid.UUID]{
		Name: res.Name,
		ID:   uuidID,
	}, nil
}

func (r *tenantService) Insert(name string) error {
	return r.repo.Insert(name)
}

func (r *tenantService) Detail(id uuid.UUID) (*model.Tenant[uuid.UUID], error) {
	res, err := r.repo.Detail(utils.ConvertUUIDToBinary(id))
	if err != nil {
		return nil, err
	}
	resID, _ := utils.ConvertBinarytoUUID(res.ID)
	return &model.Tenant[uuid.UUID]{
		Name: res.Name,
		ID:   resID,
	}, nil
}

func (r *tenantService) List(limit, page int, keyword, order, orderKey string) ([]model.Tenant[uuid.UUID], error) {

	res, err := r.repo.List(limit, page, keyword, order, orderKey)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}

	tenants := make([]model.Tenant[uuid.UUID], 0, len(res))
	for _, tb := range res {
		u, err := uuid.FromBytes(tb.ID.Data)
		if err != nil {
			return nil, err
		}
		tenants = append(tenants, model.Tenant[uuid.UUID]{
			ID:        u,
			Name:      tb.Name,
			CreatedAt: tb.CreatedAt,
			UpdatedAt: tb.UpdatedAt,
		})
	}
	return tenants, nil
}

func (r *tenantService) Update(id uuid.UUID, name string) error {
	return r.repo.Update(utils.ConvertUUIDToBinary(id), name)
}

func (r *tenantService) Delete(id uuid.UUID) error {
	return r.repo.Delete(utils.ConvertUUIDToBinary(id))
}
