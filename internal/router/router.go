package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceInterface interface {
	AuthSignIn() error
	AuthSignUp() error
	UsersGetAll() error
	UsersCreate() error
}

type Router struct {
	router  *gin.Engine
	service ServiceInterface
}

// NewRouter function
func NewRouter(service ServiceInterface) *Router {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	router := &Router{
		router:  r,
		service: service,
	}

	router.registerHandlers()

	return router
}

// GetHandler function
func (r *Router) GetHandler() http.Handler {
	return r.router
}

// registerHandlers function
func (r *Router) registerHandlers() {
	const EMPTY_PATH = ""

	auth := r.router.Group("/auth")
	{
		auth.GET("/sign-in", r.authSignIn)
		auth.GET("/sign-up", r.authSignUp)
	}

	api := r.router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET(EMPTY_PATH, r.usersGet)
			users.POST(EMPTY_PATH, r.usersCreate)
		}
	}
}
