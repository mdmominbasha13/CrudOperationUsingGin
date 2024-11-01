package router

import (
	"github.com/CrudOperationUsingGin/pkg/config"
	"github.com/CrudOperationUsingGin/pkg/controller"
	"github.com/CrudOperationUsingGin/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := NewRouter()
	router.Run(config.Appconfig.GetString("server.port"))
}

func NewRouter() *gin.Engine {
	router := gin.New()
	resource := router.Group("/api")
	resource.Use(middleware.LogRequestInfo())
	{
		resource.GET("/book/", controller.GetBook)
		resource.GET("/book/:bookid", controller.GetBookById)
		resource.POST("/book/", controller.CreateBook)
		resource.PUT("/book/:bookid", controller.UpdateBookById)
		resource.DELETE("/book/:bookid", controller.DeleteBookById)

	}
	return router
}
