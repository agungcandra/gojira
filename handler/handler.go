package handler

import (
	"net/http"

	"github.com/agungcandra/gojira/entity"
	"github.com/agungcandra/gojira/service"
	"github.com/labstack/echo/v4"
)

type GojiraHandler struct {
	hookService service.HookInterface
}

func NewGojiraHandler(hookService service.HookInterface) *GojiraHandler {
	return &GojiraHandler{
		hookService: hookService,
	}
}

func (handler *GojiraHandler) MergeRequest(c echo.Context) error {
	mergeRequest := new(entity.MergeRequestRequest)
	if err := c.Bind(mergeRequest); err != nil {
		return err
	}

	if err := handler.hookService.MergeRequest(mergeRequest); err != nil {
		return err
	}

	c.JSON(http.StatusOK, "ok")
	return nil
}

func (handler *GojiraHandler) PushRequest(c echo.Context) error {
	pushRequest := new(entity.PushRequest)
	if err := c.Bind(pushRequest); err != nil {
		return err
	}

	if err := handler.hookService.Push(pushRequest); len(err) > 0 {
		return err[0]
	}

	c.JSON(http.StatusOK, "ok")
	return nil
}
