package router

import (
	"main/services/flags"

	"github.com/labstack/echo/v4"
)

type Router struct {
	engine *echo.Echo

	flags FlagsRouter
}

func NewRouter(flagsService *flags.Service) *Router {
	e := echo.New()
	return &Router{
		engine: e,
		flags: FlagsRouter{
			flagService: flagsService,
		},
	}
}

func (r *Router) Register() {
	r.flags.register(r.engine)
	r.listen(":1323")
}

func (r *Router) listen(port string) {
	r.engine.Logger.Fatal(r.engine.Start(port))
}

type Reason struct {
	Reason string `json:"reason"`
}

func WithReason(reason string) Reason {
	return Reason{
		Reason: reason,
	}
}
