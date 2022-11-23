package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Config ...
type Config struct {
	IsDebug bool `yaml:"is_debug" env-default:"false"`
}

// DefaultConfig function
func DefaultConfig() *Config {
	return &Config{
		IsDebug: false,
	}
}

// ServiceInterface ...
type ServiceInterface interface {
	AuthSignIn() error
	AuthSignUp() error
	UsersGetAll() error
	UsersCreate() error
}

// Router ...
type Router struct {
	router  *gin.Engine
	service ServiceInterface
}

// New function
func New(cfg *Config, service ServiceInterface) http.Handler {
	if !cfg.IsDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	handler := &Router{
		router:  gin.New(),
		service: service,
	}

	handler.router.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	handler.registerHandlers()

	return handler
}

// ServeHTTP function
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
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
