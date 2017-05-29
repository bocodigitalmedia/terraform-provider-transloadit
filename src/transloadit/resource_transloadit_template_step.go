package transloadit

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceTransloaditTemplateStep() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        resourceTransloaditTemplateStepCreate,
		Read:          resourceTransloaditTemplateStepRead,
		Update:        resourceTransloaditTemplateStepUpdate,
		Delete:        resourceTransloaditTemplateStepDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"params": {
				Type:     schema.TypeMap,
				Required: true,
			},
		},
	}
}

func resourceTransloaditTemplateStepExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	log.Println("[INFO] resourceTransloaditTemplateStepExists called")
	return true, nil
}

func resourceTransloaditTemplateStepCreate(d *schema.ResourceData, meta interface{}) error {
	log.Println("[INFO] resourceTransloaditTemplateStepCreate called")
	return nil
}

func resourceTransloaditTemplateStepRead(d *schema.ResourceData, meta interface{}) error {
	log.Println("[INFO] resourceTransloaditTemplateStepRead called")
	return nil
}

func resourceTransloaditTemplateStepUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Println("[INFO] resourceTransloaditTemplateStepRead called")
	return nil
}

func resourceTransloaditTemplateStepDelete(d *schema.ResourceData, meta interface{}) error {
	log.Println("[INFO] resourceTransloaditTemplateStepDelete called")
	return nil
}
