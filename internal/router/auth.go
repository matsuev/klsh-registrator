package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// authSignIn function
func (r *Router) authSignIn(ctx *gin.Context) {
	if err := r.service.AuthSignIn(); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.String(http.StatusOK, "Sign-In")
}

// authSignUp function
func (r *Router) authSignUp(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Sign-Up")
}
