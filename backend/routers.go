package main

import (
	"github.com/gin-gonic/gin"

	"github.com/toxtashev/LogisticApp/middlewares"

	admin		"github.com/toxtashev/LogisticApp/modules/aadmin"
	declaration "github.com/toxtashev/LogisticApp/modules/declaration/controllers"
	new 		"github.com/toxtashev/LogisticApp/modules/neworder/controllers"
	active		"github.com/toxtashev/LogisticApp/modules/activeorder/controllers"
	ok			"github.com/toxtashev/LogisticApp/modules/okorder/controllers"
	delete 		"github.com/toxtashev/LogisticApp/modules/deleteorder/controllers"
	statistic 	"github.com/toxtashev/LogisticApp/modules/statistic/controllers"
)

func InitRoutes (r *gin.Engine) {

	r.Use(middlewares.Cors())

	r.POST("/declaration", declaration.Create)
	
	r.POST("/admin", admin.Login)
	
	news := r.Group("/neworders")
	news.POST("/", new.Complete)
	news.DELETE("/", new.Delete)
	news.GET("/", new.Get)
	news.GET("/:itemId", new.GetById)
	
	actives := r.Group("/activeorders")
	actives.POST("/", active.Complete)
	actives.GET("/", active.Get)
	actives.GET("/:itemId", active.GetById)

	oks := r.Group("/okgetorder")
	oks.GET("/", ok.Get)
	oks.GET("/:itemId", ok.GetById)

	deletes := r.Group("/delorders")
	deletes.GET("/", delete.Get)
	deletes.GET("/:itemId", delete.GetById)

	statistics := r.Group("/statistic")
	statistics.GET("/new", statistic.New)
	statistics.GET("/active", statistic.Active)
	statistics.GET("/ok", statistic.Ok)
	statistics.GET("/delete", statistic.Delete)

}
