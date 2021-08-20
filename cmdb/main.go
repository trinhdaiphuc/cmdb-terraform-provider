package main

import (
	"github.com/gin-gonic/gin"
	"github.com/whatvn/denny"
	"github.com/whatvn/denny/middleware/http"
	"terraform-provider-example/cmdb/controller"
)

var (
	apiVersion = "v1"
)

func setupHTTPHandler(server *denny.Denny) {
	//Register http handler
	apiController := server.Group("/api/" + apiVersion)
	apiController.POST("/names", controller.CreateName)
	apiController.GET("/names", controller.GetName)
	apiController.PUT("/names", controller.UpdateName)
	apiController.DELETE("/names", controller.DeleteName)
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
