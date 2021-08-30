package main

import (
	"github.com/gin-gonic/gin"
	"github.com/trinhdaiphuc/terraform-provider-cmdb/cmdb/controller"
	"github.com/whatvn/denny"
	"github.com/whatvn/denny/middleware/http"
)

var (
	apiVersion = "v1"
)

func setupHTTPHandler(server *denny.Denny) {
	//Register http handler
	apiController := server.Group("/api/" + apiVersion)
	apiController.POST("/configs", controller.CreateConfig)
	apiController.GET("/configs", controller.GetConfig)
	apiController.PUT("/configs", controller.UpdateConfig)
	apiController.DELETE("/configs", controller.DeleteConfig)
	apiController.GET("/histories", controller.GetHistory)
}

func main() {
	server := denny.NewServer(true)
	server.WithMiddleware(gin.Recovery(), http.Logger())
	server.RedirectTrailingSlash = false

	setupHTTPHandler(server)

	if err := server.GraceFulStart(); err != nil {
		panic(err)
	}
}
