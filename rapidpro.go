package rapidpro

import (
	"os"

	"github.com/rasoro/rapidpro-api-go/client"
	"github.com/rasoro/rapidpro-api-go/v2/flows"
	"github.com/rasoro/rapidpro-api-go/v2/flowstarts"
)

const apiURL = "https://localhost:8000/api"

type RestClient struct {
	*client.RequestHandler
	Flows      *flows.ApiService
	FlowStarts *flowstarts.ApiService
	baseURL    string
}

type ClientParams struct {
	Client client.BaseClient
	Token  string
	ApiURL string
}

func NewRestClient() *RestClient {
	return NewRestClientWithParams(ClientParams{ApiURL: apiURL})
}

func NewRestClientWithParams(params ClientParams) *RestClient {
	requestHandler := client.NewRequestHandler(params.Client)

	if params.Client == nil {
		token := params.Token
		if token == "" {
			token = os.Getenv("RAPIDPRO_API_GO_TOKEN")
		}
		defaultClient := &client.Client{
			Credentials: &client.Credentials{Token: token},
		}
		requestHandler = client.NewRequestHandler(defaultClient)
	}
	c := &RestClient{
		RequestHandler: requestHandler,
	}

	c.Flows = flows.NewService(c.RequestHandler, params.ApiURL)
	c.FlowStarts = flowstarts.NewService(c.RequestHandler, params.ApiURL)
	return c
}
