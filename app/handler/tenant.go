package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/gotired/notification-feature/app/model"
)

type TenantHandler struct {
	service model.TenantService
}

func NewTenantHandler(service model.TenantService) TenantHandler {
	return TenantHandler{service}
}

func (h *TenantHandler) Create(ctx *fiber.Ctx) error {
	var body struct {
		Name string `json:"name"`
	}
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": fmt.Sprintf("invalid input: %s", err)})
	}

	// Validate duplicate tenant
	duplicate_tenant, err := h.service.Check(body.Name)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	if duplicate_tenant != nil {
		return ctx.Status(402).JSON(fiber.Map{"error": fmt.Sprintf("tenant %s is already created", body.Name)})
	}

	// Insert Tenant
	err = h.service.Insert(body.Name)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "error while insert tenant"})
	}
	return ctx.JSON(fiber.Map{"status": fmt.Sprintf("tenant %s created successfully", body.Name)})
}

func (h *TenantHandler) Get(ctx *fiber.Ctx) error {
	tenant_id := ctx.Params("tenant_id")
	tenant_uuid, err := uuid.Parse(tenant_id)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "invalid tenant uuid"})
	}

	detail, err := h.service.Detail(tenant_uuid)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err})
	}
	if detail == nil {
		return ctx.Status(404).JSON(fiber.Map{"error": fmt.Sprintf("tenant %s not found", tenant_id)})
	}

	return ctx.JSON(fiber.Map{"data": detail})
}

func (h *TenantHandler) Delete(ctx *fiber.Ctx) error {
	tenant_id := ctx.Params("tenant_id")
	tenant_uuid, err := uuid.Parse(tenant_id)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "invalid tenant uuid"})
	}

	if err := h.service.Delete(tenant_uuid); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err})
	}

	return ctx.JSON(fiber.Map{"status": fmt.Sprintf("tenant %s deleted successfully", tenant_id)})
}

func (h *TenantHandler) Update(ctx *fiber.Ctx) error {
	tenant_id := ctx.Params("tenant_id")
	tenant_uuid, err := uuid.Parse(tenant_id)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "invalid tenant uuid"})
	}

	var body struct {
		Name string `json:"name"`
	}
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": fmt.Sprintf("invalid input: %s", err)})
	}

	err = h.service.Update(tenant_uuid, body.Name)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "error while edit tenant"})
	}
	return ctx.JSON(fiber.Map{"status": fmt.Sprintf("tenant %s edited successfully", body.Name)})
}

func (h *TenantHandler) List(ctx *fiber.Ctx) error {
	param := model.SearchOptions{}
	param.Default()
	ctx.ParamsParser(&param)

	detail, err := h.service.List(param.Limit, param.Page, param.Keyword, string(param.Order), param.OrderKey)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	if detail == nil {
		return ctx.JSON(fiber.Map{"data": []model.Tenant[uuid.UUID]{}})
	}
	return ctx.JSON(fiber.Map{"data": detail})
}
