package transloadit

import "gopkg.in/transloadit/go-sdk.v1"

type Config struct {
	AuthKey    string
	AuthSecret string
}

func (c *Config) Client() (*Client, error) {
	options := transloadit.DefaultConfig

	options.AuthKey = c.AuthKey
	options.AuthSecret = c.AuthSecret

	tl := transloadit.NewClient(options)
	client := Client{&tl}

	if err := client.Validate(); err != nil {
		return nil, err
	} else {
		return &client, nil
	}
}
