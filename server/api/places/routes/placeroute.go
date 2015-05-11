package placeroutes

import (
	"github.com/ericmdantas/random-lunch/server/api/places/controller"
	"github.com/julienschmidt/httprouter"
)

func Init(router *httprouter.Router) {
	router.GET("/api/places", placecontroller.GetAll)
	router.POST("/api/places", placecontroller.New)
	router.PUT("/api/places/:id", placecontroller.Update)
}
