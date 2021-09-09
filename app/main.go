package main

import (
	"fmt"
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

	steps, err := workflowServ.Find("SUBS-4791", "InProgress", cred)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(steps)

	e := echo.New()
	e.Use(middleware.Logger())
	e.POST("/hook", handle.Hook)
	e.POST("/merge_request", handle.MergeRequest)
	e.POST("/push", handle.PushRequest)

	log.Fatal(e.Start(":8000"))
}
