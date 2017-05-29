package transloadit

import (
	context "context"
	"errors"

	transloadit "gopkg.in/transloadit/go-sdk.v1"
)

type Client struct {
	tl *transloadit.Client
}

func (s *Client) Validate() error {
	options := &transloadit.ListOptions{PageSize: 1}

	if _, err := s.tl.ListTemplates(context.Background(), options); err != nil {
		return err
	} else {
		return nil
	}
}

func (s *Client) AssertTemplateIdIsValid(id string) error {
	if len(id) == 32 {
		return nil
	} else {
		message := "Not a valid template id: " + id
		return errors.New(message)
	}
}

func (s *Client) TemplateExists(id string) (bool, error) {
	if err := s.AssertTemplateIdIsValid(id); err != nil {
		return false, err
	}

	if _, err := s.tl.GetTemplate(context.Background(), id); err != nil {
		switch err.(transloadit.RequestError).Code {
		case "TEMPLATE_NOT_FOUND":
			return false, nil
		default:
			return false, err
		}
	} else {
		return true, nil
	}
}

type TemplateStepMap *map[string]TemplateStepParams
type TemplateStepParams *map[string]interface{}

func (s *Client) CreateTemplate(name string, steps map[string]interface{}) (*string, error) {

	template := transloadit.NewTemplate()
	template.Name = name

	for name, params := range steps {
		template.AddStep(name, params.(map[string]interface{}))
	}

	if id, err := s.tl.CreateTemplate(context.Background(), template); err != nil {
		return nil, err
	} else {
		return &id, nil
	}
}

func (s *Client) ReadTemplate(id string) (*transloadit.Template, error) {
	if template, err := s.tl.GetTemplate(context.Background(), id); err != nil {
		return nil, err
	} else {
		return &template, nil
	}
}

func (s *Client) UpdateTemplate(id string, name string, steps map[string]interface{}) error {
	template := transloadit.NewTemplate()
	template.Name = name

	for name, params := range steps {
		template.AddStep(name, params.(map[string]interface{}))
	}

	return s.tl.UpdateTemplate(context.Background(), id, template)
}

func (s *Client) DeleteTemplate(id string) error {
	return s.tl.DeleteTemplate(context.Background(), id)
}
