package main

import (
	"github.com/bocodigitalmedia/go-transloadit/transloadit_template_service"
	"github.com/hashicorp/terraform/helper/schema"
)

func NewTemplateResource() *schema.Resource {
	return &schema.Resource{
		Exists:   templateResourceExists,
		Create:   templateResourceCreate,
		Read:     templateResourceRead,
		Update:   templateResourceUpdate,
		Delete:   templateResourceDelete,
		Importer: &schema.ResourceImporter{},
		Schema:   templateResourceSchema(),
	}
}

func templateResourceExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	templateService := MetaTemplateService(meta)
	return templateService.Exists(d.Id())
}

func templateResourceCreate(d *schema.ResourceData, meta interface{}) error {
	templateService := MetaTemplateService(meta)

	params := &transloadit_template_service.CreateParams{
		Name:    d.Get("name").(string),
		Content: d.Get("content").(string),
	}

	if result, _, err := templateService.Create(params); err != nil {
		return err
	} else {
		d.SetId(result.Id)
		return nil
	}
}

func templateResourceRead(d *schema.ResourceData, meta interface{}) error {
	templateService := MetaTemplateService(meta)

	if result, _, err := templateService.Read(d.Id()); err != nil {
		return err
	} else {
		d.Set("name", result.Name)
		d.Set("content", result.Content)
		return nil
	}
}

func templateResourceUpdate(d *schema.ResourceData, meta interface{}) error {
	templateService := MetaTemplateService(meta)

	params := &transloadit_template_service.UpdateParams{
		Name:    d.Get("name").(string),
		Content: d.Get("content").(string),
	}

	if _, _, err := templateService.Update(d.Id(), params); err != nil {
		return err
	} else {
		return nil
	}
}

func templateResourceDelete(d *schema.ResourceData, meta interface{}) error {
	templateService := MetaTemplateService(meta)
	if _, _, err := templateService.Delete(d.Id()); err != nil {
		return err
	} else {
		return nil
	}
}
