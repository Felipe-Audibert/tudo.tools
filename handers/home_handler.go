package handers

import (
	"net/http"
	"tudo-tools/views"

	"github.com/gin-gonic/gin"
)

type HomeHandler struct{}

func (h *HomeHandler) GetHome(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "", views.Home())
}
