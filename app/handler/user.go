package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/gotired/notification-feature/app/model"
)

type UserHandler struct {
	userService   model.UserService
	tenantService model.TenantService
}

func NewUserHandler(userService model.UserService, tenantService model.TenantService) UserHandler {
	return UserHandler{userService, tenantService}
}

func (h *UserHandler) Create(ctx *fiber.Ctx) error {
	var body model.CreateUser
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": fmt.Sprintf("invalid input: %s", err.Error())})
	}

	// Validate duplicate user
	duplicate_user, err := h.userService.Check(body.Name)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": fmt.Sprintf("validate user error: %s", err.Error())})
	}
	if duplicate_user != nil {
		return ctx.Status(402).JSON(fiber.Map{"error": fmt.Sprintf("user %s is already created", body.Name)})
	}

	// Check Tenant
	tenant, err := h.tenantService.CheckByID(body.Tenant)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	if tenant == nil {
		return ctx.Status(402).JSON(fiber.Map{"error": fmt.Sprintf("tenant %s is not exists", body.Tenant.String())})
	}

	// Insert user
	err = h.userService.Insert(body.Name, body.Tenant)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "error while insert user"})
	}
	return ctx.JSON(fiber.Map{"status": fmt.Sprintf("user %s created successfully", body.Name)})
}

func (h *UserHandler) Get(ctx *fiber.Ctx) error {
	user_id := ctx.Params("user_id")
	user_uuid, err := uuid.Parse(user_id)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "invalid user uuid"})
	}

	detail, err := h.userService.Detail(user_uuid)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	if detail == nil {
		return ctx.Status(404).JSON(fiber.Map{"error": fmt.Sprintf("user %s not found", user_id)})
	}

	return ctx.JSON(fiber.Map{"data": detail})
}

func (h *UserHandler) Delete(ctx *fiber.Ctx) error {
	user_id := ctx.Params("user_id")
	user_uuid, err := uuid.Parse(user_id)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "invalid user uuid"})
	}

	if err := h.userService.Delete(user_uuid); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err})
	}

	return ctx.JSON(fiber.Map{"status": fmt.Sprintf("user %s deleted successfully", user_id)})
}

func (h *UserHandler) Update(ctx *fiber.Ctx) error {
	user_id := ctx.Params("user_id")
	user_uuid, err := uuid.Parse(user_id)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "invalid user uuid"})
	}

	var body struct {
		Name string `json:"name"`
	}
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": fmt.Sprintf("invalid input: %s", err)})
	}

	// Insert user
	err = h.userService.Update(user_uuid, body.Name)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "error while edit user"})
	}
	return ctx.JSON(fiber.Map{"status": fmt.Sprintf("user %s edited successfully", body.Name)})
}

func (h *UserHandler) List(ctx *fiber.Ctx) error {
	param := model.SearchOptions{}
	param.Default()
	ctx.ParamsParser(&param)

	detail, err := h.userService.List(param.Limit, param.Page, param.Keyword, string(param.Order), param.OrderKey)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	if detail == nil {
		return ctx.JSON(fiber.Map{"data": []model.User[uuid.UUID]{}})
	}
	return ctx.JSON(fiber.Map{"data": detail})
}
