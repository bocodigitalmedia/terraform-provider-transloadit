package transloadit

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	transloadit "gopkg.in/transloadit/go-sdk.v1"
)

func resourceTransloaditTemplate() *schema.Resource {

	return &schema.Resource{
		Create: resourceTransloaditTemplateCreate,
		Read:   resourceTransloaditTemplateRead,
		Update: resourceTransloaditTemplateUpdate,
		Delete: resourceTransloaditTemplateDelete,
		Exists: resourceTransloaditTemplateExists,
		// Importer: &schema.ResourceImporter{
		// 	State: resourceTransloaditTemplateImport,
		// },

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"content": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceTransloaditTemplateExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	client := meta.(*transloadit.Client)
	_, err := client.GetTemplate(context.Background(), d.Get("id").(string))
	//
	// // TODO: Check error codes
	// if err != nil {
	// 	return false, nil
	// }
	if err != nil {
		fmt.Println(err)
	}

	return err == nil, nil
}

func resourceTransloaditTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceTransloaditTemplateRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceTransloaditTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceTransloaditTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
