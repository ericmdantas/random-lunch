package userroutes

import (
	"github.com/ericmdantas/random-lunch/server/api/users/controller"
	"github.com/julienschmidt/httprouter"
)

func Init(router *httprouter.Router) {
	router.GET("/api/user", usercontroller.GetAll)
	router.GET("/api/user/:id", usercontroller.GetById)
	router.POST("/api/user", usercontroller.New)
	router.PUT("/api/user/:id", usercontroller.Update)
	router.DELETE("/api/user/:id", usercontroller.Remove)
}
