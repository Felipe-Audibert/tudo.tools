package handlers

import (
	"net/http"
	"strconv"
	"time"
	"tudo-tools/views/convert"

	"github.com/gin-gonic/gin"
	"github.com/tkuchiki/go-timezone"
)

type ConvertHandler struct{}

func (handler ConvertHandler) GetWeight(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "", views_convert.Weight(ctx.MustGet("htmxRequest").(bool)))
}

func (handler ConvertHandler) GetDistance(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "", views_convert.Distance(ctx.MustGet("htmxRequest").(bool)))
}

func (handler ConvertHandler) GetTime(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "", views_convert.Time(ctx.MustGet("htmxRequest").(bool)))
}

func (handler ConvertHandler) Convert(ctx *gin.Context) {
	fromInput := ctx.PostForm("from-input")
	fromSelect := ctx.PostForm("from-select")
	toSelect := ctx.PostForm("to-select")
	conversion := ctx.PostForm("conversion")

	output := ""
	var err *error = nil

	switch conversion {
	case "multiply":
		output, err = handler.Multiply(fromInput, fromSelect, toSelect)
		break

	case "timezone":
		output, err = handler.Timezone(fromInput, fromSelect, toSelect)
		break
	}

	ctx.Header("Content-Type", "text/plain; charset=utf-8")

	if err != nil {
		ctx.String(http.StatusBadRequest, output)
		return
	}

	ctx.String(http.StatusOK, output)
}

func (handler ConvertHandler) Multiply(fromInput string, fromSelect string, toSelect string) (string, *error) {
	fromValue, err := strconv.ParseFloat(fromInput, 64)
	if err != nil {
		return "", nil
	}

	fromMultiplier, err := strconv.ParseFloat(fromSelect, 64)
	if err != nil {
		return "Multiplier cannot be empty", &err
	}

	toMultiplier, err := strconv.ParseFloat(toSelect, 64)
	if err != nil {
		return "Multiplier cannot be null", &err
	}

	output := fromMultiplier * fromValue / toMultiplier

	return strconv.FormatFloat(output, 'f', 4, 64), nil
}

func (handler ConvertHandler) Timezone(fromInput string, fromSelect string, toSelect string) (string, *error) {
	tz := timezone.New()

	fromTime, err := time.Parse("15:04", fromInput)
	if err != nil {
		return "Could not parse time", &err
	}

	fromTimezone, err := tz.GetTzInfo(fromSelect)
	if err != nil {
		return "Timezone not found", &err
	}

	toTimezone, err := tz.GetTzInfo(toSelect)
	if err != nil {
		return "Timezone not found", &err
	}

	output := fromTime.Add(time.Second * time.Duration(-1*fromTimezone.StandardOffset())).Add(time.Second * time.Duration(toTimezone.StandardOffset()))

	return output.Format("15:04"), nil
}
