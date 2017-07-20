package main

import (
	"github.com/bocodigitalmedia/go-transloadit/transloadit_api"
	"github.com/bocodigitalmedia/go-transloadit/transloadit_template_service"
	"github.com/hashicorp/terraform/helper/schema"
)

func Configure(d *schema.ResourceData) (interface{}, error) {

	params := transloadit_api.NewApiParams{
		AuthKey:    d.Get("auth_key").(string),
		AuthSecret: d.Get("auth_secret").(string),
	}

	api := transloadit_api.New(&params)
	templateService := transloadit_template_service.New(api)

	return NewMeta(templateService), nil
}
