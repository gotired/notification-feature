package model

type NotificationBody struct {
	TenantID string `json:"tenant_id"`
	Message  string `json:"message"`
}
