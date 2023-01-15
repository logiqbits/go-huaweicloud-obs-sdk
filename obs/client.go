package obs

import (
	"github.com/logiqbits/go-huaweicloud-obs-sdk/internal/sdk"
)

type Client struct {
	endpoint    string
	obsClient   *sdk.ObsClient
	aliasDomain string
}

func NewClient(endpoint, accessKey, secretKey string) (*Client, error) {
	obsClient, err := sdk.New(accessKey, secretKey, endpoint)
	if err != nil {
		return nil, err
	}
	return &Client{obsClient: obsClient, endpoint: endpoint}, nil
}

func (c *Client) WithAliasDomain(aliasDomain string) *Client {
	c.aliasDomain = aliasDomain
	return c
}
