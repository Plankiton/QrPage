package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func (h *Handler) GetInviteCard(c echo.Context) error {
	logger := c.Logger()
	ctx := c.Request().Context()

	slugCode := c.Param("slugCode")
	inviteCard, err := h.rps.inviteCardsRp.BySlugCode(ctx, slugCode)
	if err != nil {
		logger.Errorj(log.JSON{
			"message": "error searching for inviteCard",
			"err":     err.Error(),
		})

		return c.JSON(404, echo.Map{
			"ok":  false,
			"err": "error searching for inviteCard",
		})
	}

	return c.JSON(200, echo.Map{
		"ok":   true,
		"data": inviteCard,
	})
}
