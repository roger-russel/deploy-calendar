package api

import (
	"encoding/json"
	"fmt"

	"github.com/roger-russel/deploy-calendar/pkg/logger"
	"github.com/valyala/fasthttp"
)

func setJsonHeader(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json; charset=utf-8")
}

func JsonOut(ctx *fasthttp.RequestCtx, v interface{}) {

	setJsonHeader(ctx)

	json, err := json.Marshal(v)

	if err != nil {
		logger.Error("An error ocourrer on JsonOut: %s", err)
	}

	fmt.Fprintf(ctx, string(json))

}
