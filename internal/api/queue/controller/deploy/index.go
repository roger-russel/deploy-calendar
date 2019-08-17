package deploy

import (
	"encoding/json"

	"github.com/roger-russel/deploy-calendar/pkg/api"
	"github.com/roger-russel/deploy-calendar/pkg/logger"
	"github.com/valyala/fasthttp"
)

type createRequest struct {
	ProjectHash    string `json:"project_hash,omitempty"`
	Context        string `json:"context,omitempty"`
	Version        string `json:"version,omitempty"`
	PullRequestURL string `json:"pull_request_url,omitempty"`
}

func Create(ctx *fasthttp.RequestCtx) {

	status, message, err := createHandler(ctx.PostBody())

	response := api.SimpleResponse{
		Message: message,
		Error:   err,
	}

	ctx.SetStatusCode(status)

	if status != fasthttp.StatusAccepted {
		api.JsonOut(ctx, response)
	}

}

func createHandler(data []byte) (int, string, error) {

	var request = &createRequest{}

	err := json.Unmarshal(data, request)

	if err != nil {
		logger.Error("error unmarshal create deploy request", err)
		return fasthttp.StatusBadRequest, "body data could not be parsed as json", err
	}

	//ToDO Validar request

	// jsonRequest, _ := json.Marshal(request)

	//nc := nats.Pub(nats.SubEventDeploy, *jsonRequest)

	return fasthttp.StatusAccepted, "", nil

}
