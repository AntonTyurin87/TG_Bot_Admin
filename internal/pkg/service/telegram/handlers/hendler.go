package telegram

import (
	"TG_Bot_Admin/internal/pkg/domain/presenter"
	"TG_Bot_Admin/internal/pkg/service/admin"
)

// Handler обрабатывает команды Telegram бота
type Handler struct {
	adminService admin.Interface
	presenter    presenter.Interface
}

// NewHandler создает новый обработчик
func NewHandler(
	adminService admin.Interface,
	presenter presenter.Interface,
) *Handler {
	return &Handler{
		adminService: adminService,
		presenter:    presenter,
	}
}
