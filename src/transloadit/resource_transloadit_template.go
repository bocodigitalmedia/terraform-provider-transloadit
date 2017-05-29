package transloadit

import (
	"sort"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceTransloaditTemplate() *schema.Resource {

	return &schema.Resource{
		SchemaVersion: 1,
		Create:        resourceTransloaditTemplateCreate,
		Read:          resourceTransloaditTemplateRead,
		Update:        resourceTransloaditTemplateUpdate,
		Delete:        resourceTransloaditTemplateDelete,
		Exists:        resourceTransloaditTemplateExists,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"step": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     resourceTransloaditTemplateStep(),
			},
		},
	}

}

func resourceTransloaditTemplateStep() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
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

func resourceTransloaditTemplateExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	client := meta.(*Client)
	return client.TemplateExists(d.Id())
}

func resourceTransloaditTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)

	name := d.Get("name").(string)
	stepList := getResourceDataTemplateStepList(d)
	stepMap := createTemplateStepMapFromList(stepList)

	if id, err := client.CreateTemplate(name, stepMap); err != nil {
		return err
	} else {
		d.SetId(*id)
		d.Set("name", name)
		setResourceDataTemplateStepList(d, stepList)

		return nil
	}
}

func resourceTransloaditTemplateRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	id := d.Id()

	if template, err := client.ReadTemplate(id); err != nil {
		return err
	} else {
		stepList := createTemplateStepListFromMap(template.Content.Steps)

		d.SetId(template.ID)
		d.Set("name", template.Name)
		setResourceDataTemplateStepList(d, stepList)

		return nil
	}
}

func resourceTransloaditTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)

	id := d.Id()

	name := d.Get("name").(string)
	stepList := getResourceDataTemplateStepList(d)
	stepMap := createTemplateStepMapFromList(stepList)

	if err := client.UpdateTemplate(id, name, stepMap); err != nil {
		return err
	} else {
		return resourceTransloaditTemplateRead(d, meta)
	}
}

func resourceTransloaditTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)

	id := d.Id()

	if err := client.DeleteTemplate(id); err != nil {
		return err
	} else {
		return nil
	}
}

type TemplateStepListByName []interface{}

func (s TemplateStepListByName) Len() int {
	return len(s)
}

func (s TemplateStepListByName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s TemplateStepListByName) GetName(i int) string {
	data := s[i].(map[string]interface{})
	return data["name"].(string)
}

func (s TemplateStepListByName) Less(i, j int) bool {
	a, b := s.GetName(i), s.GetName(j)
	return strings.Compare(a, b) < 0
}

func getResourceDataTemplateStepList(d *schema.ResourceData) []interface{} {
	stepList := d.Get("step").([]interface{})
	sortTemplateStepListByName(stepList)
	return stepList
}

func setResourceDataTemplateStepList(d *schema.ResourceData, stepList []interface{}) {
	sortTemplateStepListByName(stepList)
	d.Set("step", stepList)
}

func sortTemplateStepListByName(stepList []interface{}) {
	sort.Sort(TemplateStepListByName(stepList))
}

func createTemplateStepListFromMap(stepMap map[string]interface{}) []interface{} {
	i := 0
	stepList := make([]interface{}, len(stepMap))

	for name, params := range stepMap {
		stepData := map[string]interface{}{}
		stepData["name"] = name
		stepData["params"] = params.(map[string]interface{})
		stepList[i] = stepData
		i += 1
	}

	return stepList
}

func createTemplateStepMapFromList(stepList []interface{}) map[string]interface{} {
	stepMap := map[string]interface{}{}
	for _, v := range stepList {
		data := v.(map[string]interface{})
		name := data["name"].(string)
		params := data["params"].(map[string]interface{})
		stepMap[name] = params
	}
	return stepMap
}
