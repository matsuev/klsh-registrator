package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// usersGet function
func (r *Router) usersGet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Users handler",
	})
}

// usersCreate function
func (r *Router) usersCreate(ctx *gin.Context) {
	ctx.String(http.StatusNoContent, "Users create")
}
