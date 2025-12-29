package handlers

import (
	"net/http"
	"tudo-tools/views/convert"

	"github.com/gin-gonic/gin"
)

type ConvertHandler struct{}

func (handler ConvertHandler) GetDistance(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "", views_convert.Distance(ctx.MustGet("htmxRequest").(bool)))
}

func (handler ConvertHandler) ConvertDistance(ctx *gin.Context) {
	//inputUnit := ctx.PostForm("input_unit")
	//outputUnit := ctx.PostForm("output_unit")
	//inputValue := ctx.PostForm("input_value")

	ctx.Header("Content-Type", "text/plain; charset=utf-8")
	ctx.String(http.StatusOK, "20")
}
