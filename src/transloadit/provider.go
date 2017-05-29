package transloadit

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"auth_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TRANSLOADIT_AUTH_KEY", nil),
			},
			"auth_secret": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TRANSLOADIT_AUTH_SECRET", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"transloadit_template": resourceTransloaditTemplate(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		AuthKey:    d.Get("auth_key").(string),
		AuthSecret: d.Get("auth_secret").(string),
	}

	return config.Client()
}
