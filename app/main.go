package main

import (
	"log"
	"os"

	"github.com/agungcandra/gojira/config"
	"github.com/agungcandra/gojira/handler"
	"github.com/agungcandra/gojira/request"
	"github.com/agungcandra/gojira/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	_ = godotenv.Load()

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")

	cred := config.NewCredential(username, password)
	workflow := config.LoadWorklow()

	req := request.NewRequest(host)

	workflowServ := service.NewWorkflowService(workflow, req)
	extractorServ := service.NewExtractorService()
	transitionServ := service.NewTransitionService(req, workflowServ)
	hookServ := service.NewHookService(transitionServ, extractorServ, *workflowServ, cred)
	handle := handler.NewGojiraHandler(hookServ)

	server := echo.New()
	server.Use(
		middleware.Recover(),
		middleware.Logger(),
		middleware.RequestID(),
	)
	server.HideBanner = true

	server.POST("/hook", handle.Hook)
	server.POST("/merge_request", handle.MergeRequest)
	server.POST("/push", handle.PushRequest)

	log.Fatal(server.Start(":8000"))
}
