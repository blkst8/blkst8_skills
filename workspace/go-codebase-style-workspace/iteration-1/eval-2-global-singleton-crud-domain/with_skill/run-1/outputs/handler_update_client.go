package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/blkst8/client-service/internal/app"
	"github.com/blkst8/client-service/internal/log"
	"github.com/blkst8/client-service/internal/models"
	"github.com/blkst8/client-service/internal/repository"
)

type UpdateClientRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// UpdateClient handles PUT /api/v1/clients/:id.
func UpdateClient(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid client id")
	}

	var req UpdateClientRequest
	if err := ctx.Bind(&req); err != nil {
		log.Logger.Error("failed to bind update client request", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	client := models.Client{
		ID:    uint32(id),
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}

	if err := app.A.Service.Client.Update(ctx.Request().Context(), client); err != nil {
		if errors.Is(err, repository.ErrClientNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "client not found")
		}
		log.Logger.Error("failed to update client", zap.Uint64("id", id), zap.Error(err))
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{"status": "updated"})
}
