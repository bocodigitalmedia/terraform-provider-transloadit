package main

import (
	terraform_transloadit "transloadit-terraform-provider/src/transloadit"

	"github.com/hashicorp/terraform/plugin"
)

func main() {
	opts := plugin.ServeOpts{
		ProviderFunc: terraform_transloadit.Provider,
	}
	plugin.Serve(&opts)
}
