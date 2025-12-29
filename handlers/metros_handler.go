package handlers

import (
	"net/http"
	"tudo-tools/views"

	"github.com/gin-gonic/gin"
)

type MetrosHandler struct{}

func (h *MetrosHandler) GetMetros(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "/Metros", views.Metros(ctx.MustGet("htmxRequest").(bool)))
}
