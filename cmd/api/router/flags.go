package router

import (
	"main/services/flags"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type FlagsRouter struct {
	flagService *flags.Service
}

func (r *FlagsRouter) register(e *echo.Echo) {
	e.POST("/flags", r.createFlag)
	e.GET("/flags/:id", r.getFlag)
}

func (fr *FlagsRouter) getFlag(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, WithReason("invalid id"))
	}

	response, err := fr.flagService.Get(ctx, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, WithReason(err.Error()))
	}

	return c.JSON(http.StatusOK, response)
}

type createFlagRequest struct {
	Name string `json:"name"`
}

func (fr *FlagsRouter) createFlag(c echo.Context) error {
	ctx := c.Request().Context()

	var payload createFlagRequest
	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, WithReason("invalid payload"))
	}

	flag, err := fr.flagService.Create(ctx, payload.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, WithReason("failed to create tag"))
	}

	return c.JSON(http.StatusCreated, flag)
}
