package handler

import (
	"eticketing/entity"
	"eticketing/handler/request"
	"eticketing/handler/response"
	"eticketing/service"
	"eticketing/validate"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TerminalHandler struct {
	terminalService *service.TerminalService
}

func NewTerminalHandler(
	terminalService *service.TerminalService,
) *TerminalHandler {
	return &TerminalHandler{
		terminalService: terminalService,
	}
}

func (t *TerminalHandler) Create(c echo.Context) error {

	var tReq request.Terminal
	err := c.Bind(&tReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to read json request",
					Code:    "BAD_REQUEST",
				},
			},
		})
	}
	err = validate.Validate(tReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: err.Error(),
					Code:    "TERMINAL_INVALID",
				},
			},
		})
	}
	terminal := entity.Terminal{
		Name: tReq.Name,
	}
	err = t.terminalService.AddTerminal(c.Request().Context(), &terminal)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to create product",
					Code:    "TERMINAL_CREATE-ERROR",
				},
			},
		})
	}
	res := response.BuildTerminal(terminal)
	return c.JSON(http.StatusCreated, res)
}

func (t *TerminalHandler) GetAll(c echo.Context) error {
	terminals, err := t.terminalService.GetAllTerminal(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to read all product",
					Code:    "TERMINAL_READ-ALL-ERROR",
				},
			},
		})
	}
	fmt.Println(terminals)
	res := response.BuildTerminals(terminals)

	return c.JSON(http.StatusOK, res)
}
