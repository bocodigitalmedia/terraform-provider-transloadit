package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func NewProvider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema:        providerSchema(),
		ResourcesMap:  providerResourcesMap(),
		ConfigureFunc: Configure,
	}
}

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"auth_key": {
			Type:     schema.TypeString,
			Required: true,
		},
		"auth_secret": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func providerResourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"transloadit_template": NewTemplateResource(),
	}
}
