package main

import (
	terraform_transloadit "terraform-transloadit-provider/src/transloadit"

	"github.com/hashicorp/terraform/plugin"
)

func main() {
	opts := plugin.ServeOpts{
		ProviderFunc: terraform_transloadit.Provider,
	}
	plugin.Serve(&opts)
}
