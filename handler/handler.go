package handler

import (
	"net/http"

	"github.com/agungcandra/gojira/entity"
	"github.com/agungcandra/gojira/service"
	"github.com/labstack/echo/v4"
)

const (
	MergeRequestType = "Merge Request Hook"
	PushRequesType   = "Push Hook"
)

type GojiraHandler struct {
	hookService service.HookInterface
}

func NewGojiraHandler(hookService service.HookInterface) *GojiraHandler {
	return &GojiraHandler{
		hookService: hookService,
	}
}

func (handle *GojiraHandler) MergeRequest(c echo.Context) error {
	mergeRequest := new(entity.MergeRequestRequest)
	if err := c.Bind(mergeRequest); err != nil {
		return err
	}

	if err := handle.hookService.MergeRequest(mergeRequest); len(err) > 0 {
		return err[0]
	}

	c.JSON(http.StatusOK, "ok")
	return nil
}

func (handle *GojiraHandler) PushRequest(c echo.Context) error {
	pushRequest := new(entity.PushRequest)
	if err := c.Bind(pushRequest); err != nil {
		return err
	}

	if err := handle.hookService.Push(pushRequest); len(err) > 0 {
		return err[0]
	}

	c.JSON(http.StatusOK, "ok")
	return nil
}

func (handle *GojiraHandler) Hook(c echo.Context) error {
	event := c.Request().Header.Get("X-Gitlab-Event")
	switch event {
	case MergeRequestType:
		return handle.MergeRequest(c)
	case PushRequesType:
		return handle.PushRequest(c)
	}

	return nil
}
