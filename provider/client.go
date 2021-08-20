package provider

import (
	"context"
	"fmt"
	builder "github.com/phuc1998/http-builder"
)

type RequestParams struct {
	Name  string `http:"name,form"`
	Value string `http:"value,form"`
}

type Config struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	CreatedAt string `json:"createdAt"`
	UpdateAt  string `json:"updateAt"`
}

type Client struct {
	*builder.APIClient
}

// NewClient initializes a new Client instance to communicate with the CMDB api
func NewClient(port int, protocol, hostname, version string) *Client {
	httpCfg := builder.NewConfiguration()
	httpCfg.BasePath = fmt.Sprintf("%s://%s:%d/api/%s", protocol, hostname, port, version)
	httpClient := builder.NewAPIClient(httpCfg)

	return &Client{httpClient}
}

func (c *Client) CreateName(name, value string) (*Config, error) {
	var (
		req = RequestParams{
			Name:  name,
			Value: value,
		}
		resp = &Config{}
	)
	_, err := c.Builder("/names").
		Post().
		BuildForm(req).
		UseMultipartFormData().
		Call(context.Background(), resp)
	return resp, err
}

func (c *Client) GetName(name string) (*Config, error) {
	var (
		resp = &Config{}
	)
	_, err := c.Builder("/names").
		Get().
		SetQuery("name", name).
		Call(context.Background(), resp)
	return resp, err
}

func (c *Client) UpdateName(name, value string) (*Config, error) {
	var (
		req = RequestParams{
			Name:  name,
			Value: value,
		}
		resp = &Config{}
	)
	_, err := c.Builder("/names").
		Put().
		BuildForm(req).
		UseXFormURLEncoded().
		Call(context.Background(), resp)
	return resp, err
}
func (c *Client) DeleteName(name string) error {
	var (
		resp interface{}
	)
	_, err := c.Builder("/names").
		Delete().
		SetQuery("name", name).
		Call(context.Background(), &resp)
	return err
}
