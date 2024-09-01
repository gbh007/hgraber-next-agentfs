package masterAPI

import (
	"context"
	"hgnextfs/open_api/serverAPI"
	"net/http"
	"time"

	"github.com/ogen-go/ogen/ogenerrors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

const agentTimeout = time.Minute * 1

type Client struct {
	rawClient *serverAPI.Client
}

func New(baseURL string, token string) (*Client, error) {
	httpClient := http.Client{
		Transport: otelPropagationRT{next: http.DefaultTransport},
		Timeout:   agentTimeout,
	}

	rawClient, err := serverAPI.NewClient(
		baseURL,
		securitySource{
			token: token,
		},
		serverAPI.WithClient(&httpClient),
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		rawClient: rawClient,
	}, nil
}

type otelPropagationRT struct {
	next http.RoundTripper
}

func (rt otelPropagationRT) RoundTrip(req *http.Request) (*http.Response, error) {
	req = req.Clone(req.Context())
	otel.GetTextMapPropagator().Inject(req.Context(), propagation.HeaderCarrier(req.Header))

	return rt.next.RoundTrip(req)
}

type securitySource struct {
	token string
}

func (s securitySource) HeaderAuth(ctx context.Context, operationName string) (serverAPI.HeaderAuth, error) {
	return serverAPI.HeaderAuth{
		APIKey: s.token,
	}, nil
}

func (s securitySource) Cookies(ctx context.Context, operationName string) (serverAPI.Cookies, error) {
	// FIXME: убрать если не будет проблем.
	// return serverAPI.Cookies{
	// 	APIKey: s.token,
	// }, nil

	return serverAPI.Cookies{}, ogenerrors.ErrSkipClientSecurity
}
