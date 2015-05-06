package routes

import (
	"github.com/ericmdantas/random-lunch/server/api/todo/routes"
	"github.com/ericmdantas/random-lunch/server/common/static"
	"github.com/julienschmidt/httprouter"
)

func Init(router *httprouter.Router) {
	todoroutes.Init(router)
	static.Init(router)
}
