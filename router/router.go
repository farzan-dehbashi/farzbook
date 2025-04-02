package router

import (
	"net/http"

	"farzbook.com/m/handler"
	"github.com/gin-gonic/gin"
)

type Router struct {
        engine *gin.Engine
}

func New() *Router {
        return &Router{
                engine: gin.Default(),
        }
}

func (r *Router) RegisterRoutes(farzbook *handler.FarzbookHandler) {
        r.engine.GET("/", farzbook.Home)
        r.engine.GET("/login", farzbook.Login)
        r.engine.GET("/profile", farzbook.Profile)
        r.engine.GET("/logout", farzbook.Logout)
        r.engine.GET("/add", farzbook.Add)
        r.engine.GET("/sub", farzbook.Sub)
}

func (r *Router) Handler() http.Handler {
        return r.engine
}
