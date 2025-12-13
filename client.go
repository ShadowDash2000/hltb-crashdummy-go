package hltb_crashdummy_go

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"resty.dev/v3"
)

const defaultTimeout = 10 * time.Second

type Client struct {
	clientOpts
}

type Option func(opts *clientOpts)

type clientOpts struct {
	client  *resty.Client
	timeout time.Duration
}

func defaultOpts() clientOpts {
	return clientOpts{
		client:  resty.New(),
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

func (c *Client) get(ctx context.Context, url string, output any) error {
	res, err := c.client.R().
		SetHeader("Accept", "application/json").
		SetContext(ctx).
		SetTimeout(c.timeout).
		SetResult(output).
		Execute(http.MethodGet, url)
	if err != nil {
		return err
	}

	if res.IsError() {
		return fmt.Errorf("hltb.get(): %s, status code = %d", res.String(), res.StatusCode())
	}

	return nil
}

func (c *Client) post(ctx context.Context, url string, input any, output any) error {
	res, err := c.client.R().
		SetHeader("Accept", "application/json").
		SetContext(ctx).
		SetTimeout(c.timeout).
		SetBody(input).
		SetResult(output).
		Execute(http.MethodPost, url)
	if err != nil {
		return err
	}

	if res.IsError() {
		return fmt.Errorf("hltb.post(): %s, status code = %d", res.String(), res.StatusCode())
	}

	return nil
}
