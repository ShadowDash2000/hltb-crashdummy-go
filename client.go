package hltb

import (
	"context"
	"net/http"
	"time"

	"golang.org/x/time/rate"
	"resty.dev/v3"
)

const (
	apiUrl         = "https://hltbapi.codepotatoes.de"
	defaultTimeout = 30 * time.Second
)

type Client struct {
	clientOpts
}

type Option func(opts *clientOpts)

type clientOpts struct {
	client  *resty.Client
	baseUrl string
	timeout time.Duration
}

func defaultOpts() clientOpts {
	return clientOpts{
		client:  resty.New(),
		baseUrl: apiUrl,
		timeout: defaultTimeout,
	}
}

func WithTimeout(timeout int) Option {
	return func(opts *clientOpts) {
		if timeout > 0 {
			opts.timeout = time.Duration(timeout) * time.Second
		}
	}
}

func WithRetryCount(count int) Option {
	return func(opts *clientOpts) {
		opts.client.SetRetryCount(count)
	}
}

func WithRateLimit(r time.Duration, b int) Option {
	limiter := rate.NewLimiter(rate.Every(r), b)
	return func(opts *clientOpts) {
		opts.client.AddRequestMiddleware(func(client *resty.Client, request *resty.Request) error {
			return limiter.Wait(request.Context())
		})
	}
}

func WithBaseUrl(baseUrl string) Option {
	return func(opts *clientOpts) {
		if baseUrl != "" {
			opts.baseUrl = baseUrl
		}
	}
}

func New(opts ...Option) *Client {
	c := &Client{
		clientOpts: defaultOpts(),
	}

	for _, opt := range opts {
		opt(&c.clientOpts)
	}

	return c
}

// Client returns the resty client so you can configure it
func (c *Client) Client() *resty.Client {
	return c.client
}

func (c *Client) get(ctx context.Context, url string, output any) (*resty.Response, error) {
	return c.client.R().
		SetHeader("Accept", "application/json").
		SetContext(ctx).
		SetTimeout(c.timeout).
		SetResult(output).
		Execute(http.MethodGet, url)
}

func (c *Client) post(ctx context.Context, url string, input any, output any) (*resty.Response, error) {
	return c.client.R().
		SetHeader("Accept", "application/json").
		SetContext(ctx).
		SetTimeout(c.timeout).
		SetBody(input).
		SetResult(output).
		Execute(http.MethodPost, url)
}
