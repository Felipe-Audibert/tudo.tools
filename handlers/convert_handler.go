package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"tudo-tools/views/convert"

	"github.com/gin-gonic/gin"
)

type ConvertHandler struct{}

func (handler ConvertHandler) GetDistance(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "", views_convert.Distance(ctx.MustGet("htmxRequest").(bool)))
}

func (handler ConvertHandler) Convert(ctx *gin.Context) {
	fromInput := ctx.PostForm("from-input")
	fromSelect := ctx.PostForm("from-select")
	toSelect := ctx.PostForm("to-select")
	conversion := ctx.PostForm("conversion")

	fmt.Println(fromInput, fromSelect, toSelect, conversion)

	output := ""

	switch conversion {
	case "multiply":
		output = handler.Multiply(fromInput, fromSelect, toSelect)
	}

	ctx.Header("Content-Type", "text/plain; charset=utf-8")
	ctx.String(http.StatusOK, output)
}

func (handler ConvertHandler) Multiply(fromInput string, fromSelect string, toSelect string) string {
	fromValue, err := strconv.ParseFloat(fromInput, 64)
	if err != nil {
		return ""
	}

	fromMultiplier, err := strconv.ParseFloat(fromSelect, 64)
	if err != nil {
		return ""
	}

	toMultiplier, err := strconv.ParseFloat(toSelect, 64)
	if err != nil {
		return ""
	}
	fmt.Println(fromInput, fromSelect, toSelect)

	output := fromMultiplier * fromValue / toMultiplier

	return strconv.FormatFloat(output, 'f', 4, 64)
}
