package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/plankiton/qrpage/pkg/db"
)

// GetInviteCard Godoc
// @Summary get a invite card by SlugCode
// @ID get-invite-card-by-slugcode
// @Produce json
// @Param id path string true "todo ID"
// @Success 200 {object}
// @Failure 404 {object} message
// @Router /api/i/{SlugCode} [get]
func (h *Handler) GetInviteCard(c echo.Context) error {
	ctx := c.Request().Context()
	logger := c.Logger()

	slugCode := c.Param("slugCode")
	inviteCard, err := h.rps.inviteCardsRp.BySlugCode(ctx, slugCode)
	if err != nil {
		logger.Errorj(log.JSON{
			"message": "error searching for inviteCard",
			"err":     err.Error(),
		})

		return c.JSON(404, GetInviteCardPayload{Error: "error searching for inviteCard"})
	}

	return c.JSON(200, GetInviteCardPayload{
		Data: inviteCard,
	})
}

type GetInviteCardPayload struct {
	Data  *db.InviteCard `json:"data"`
	Error string         `json:"error,omitempty"`
}
