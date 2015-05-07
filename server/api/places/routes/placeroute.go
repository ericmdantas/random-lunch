package placeroutes

import (
	"github.com/ericmdantas/random-lunch/server/api/places//controller"
	"github.com/julienschmidt/httprouter"
)

func Init(router *httprouter.Router) {
	router.GET("/api/place", placecontroller.GetAll)
	router.GET("/api/place/:id", placecontroller.GetById)
	router.POST("/api/place", placecontroller.New)
	router.PUT("/api/place/:id", placecontroller.Update)
	router.DELETE("/api/place/:id", placecontroller.Remove)
}
