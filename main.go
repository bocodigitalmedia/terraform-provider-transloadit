package main

import "github.com/hashicorp/terraform/plugin"
import "transloadit-terraform-provider/transloadit"

func main() {
	opts := plugin.ServeOpts{
		ProviderFunc: transloadit.Provider,
	}
	plugin.Serve(&opts)
}
