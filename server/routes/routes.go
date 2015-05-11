package routes

import (
	"github.com/ericmdantas/random-lunch/server/api/places/routes"
	_ "github.com/ericmdantas/random-lunch/server/api/users/routes"
	"github.com/ericmdantas/random-lunch/server/common/static"
	"github.com/julienschmidt/httprouter"
)

func Init(router *httprouter.Router) {
	//userroutes.Init(router)
	placeroutes.Init(router)
	static.Init(router)
}
