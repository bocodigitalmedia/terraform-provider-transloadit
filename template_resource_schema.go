package main

import "github.com/hashicorp/terraform/helper/schema"

func templateResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"content": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}
