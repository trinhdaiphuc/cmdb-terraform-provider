package provider

import (
	"context"
	"fmt"
	builder "github.com/phuc1998/http-builder"
)

type Client struct {
	*builder.APIClient
}

// NewClient initializes a new Client instance to communicate with the CMDB api
func NewClient(host, version string) *Client {
	httpCfg := builder.NewConfiguration()
	httpCfg.BasePath = fmt.Sprintf("%s/api/%s", host, version)
	httpClient := builder.NewAPIClient(httpCfg)

	return &Client{httpClient}
}

func (c *Client) CreateConfig(config Config) (Config, error) {
	var (
		req = RequestParams{
			Name:  config.Name,
			Value: config.Value,
		}
		resp = Config{}
	)
	_, err := c.Builder("/configs").
		Post().
		BuildForm(req).
		UseMultipartFormData().
		Call(context.Background(), &resp)
	return resp, err
}

func (c *Client) GetConfig(name string) (Config, error) {
	var (
		resp = Config{}
	)
	_, err := c.Builder("/configs").
		Get().
		SetQuery("name", name).
		Call(context.Background(), &resp)
	return resp, err
}

func (c *Client) UpdateConfig(config Config) (Config, error) {
	var (
		req = RequestParams{
			Name:  config.Name,
			Value: config.Value,
		}
		resp = Config{}
	)
	_, err := c.Builder("/configs").
		Put().
		BuildForm(req).
		UseXFormURLEncoded().
		Call(context.Background(), &resp)
	return resp, err
}

func (c *Client) DeleteConfig(name string) error {
	var (
		resp interface{}
	)
	_, err := c.Builder("/configs").
		Delete().
		SetQuery("name", name).
		Call(context.Background(), &resp)
	return err
}

func (c *Client) GetHistory(name string) (History, error) {
	var (
		resp = History{}
	)
	_, err := c.Builder("/histories").
		Get().
		SetQuery("name", name).
		Call(context.Background(), &resp)
	return resp, err
}
