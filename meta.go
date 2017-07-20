package main

import transloadit_template_service "github.com/bocodigitalmedia/go-transloadit/transloadit_template_service"

type Meta struct {
	TransloaditTemplateService *transloadit_template_service.Service
}

func NewMeta(templateService *transloadit_template_service.Service) *Meta {
	return &Meta{
		TransloaditTemplateService: templateService,
	}
}

func MetaTemplateService(meta interface{}) *transloadit_template_service.Service {
	return meta.(*Meta).TransloaditTemplateService
}
