package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/blkst8/client-service/internal/app"
	"github.com/blkst8/client-service/internal/log"
	"github.com/blkst8/client-service/internal/models"
)

type CreateClientRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type CreateClientResponse struct {
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateClient handles POST /api/v1/clients.
func CreateClient(ctx echo.Context) error {
	var req CreateClientRequest
	if err := ctx.Bind(&req); err != nil {
		log.Logger.Error("failed to bind create client request", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	client := models.Client{
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}

	if err := app.A.Service.Client.Create(ctx.Request().Context(), client); err != nil {
		log.Logger.Error("failed to create client", zap.Error(err))
		return err
	}

	return ctx.JSON(http.StatusCreated, CreateClientResponse{
		Status:    "created",
		CreatedAt: client.CreatedAt,
	})
}
