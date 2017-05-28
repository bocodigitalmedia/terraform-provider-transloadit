package transloadit

import (
	"log"

	"gopkg.in/transloadit/go-sdk.v1"
)

type Config struct {
	AuthKey    string
	AuthSecret string
}

func (c *Config) Client() *transloadit.Client {
	options := transloadit.DefaultConfig
	options.AuthKey = c.AuthKey
	options.AuthSecret = c.AuthSecret
	client := transloadit.NewClient(options)
	log.Printf("[INFO] Transloadit Client configured ")
	return &client
}
