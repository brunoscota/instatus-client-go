package instatus

const templateName = "template"

type Template struct {
	Subdomain  *string             `json:"subdomain"`
	Name       *string             `json:"name"`
	Type       *string             `json:"type"`
	Message    *string             `json:"message"`
	Status     *string             `json:"status"`
	Components []TemplateComponent `json:"components"`
	Notify     *bool               `json:"notify"`
}

type TemplateComponent struct {
	ID          *string `json:"id"`
	ComponentID *string `json:"componentId,omitempty"`
	Status      *string `json:"status"`
}

type TemplateFull struct {
	Template
	ID          *string `json:"id"`
	CreatedAt   *string `json:"createdAt"`
	MessageHtml *string `json:"messageHtml,omitempty"`
}

func (client *Client) CreateTemplate(pageID string, template *Template) (*TemplateFull, error) {
	var i TemplateFull
	err := createResource(
		client,
		pageID,
		templateName,
		template,
		&i,
	)

	return &i, err
}

func (client *Client) GetTemplate(pageID, templateID string) (*TemplateFull, error) {
	var i TemplateFull
	err := readResource(client, pageID, templateID, templateName, &i)

	return &i, err
}

func (client *Client) UpdateTemplate(pageID, templateID string, template *Template) (*TemplateFull, error) {
	var i TemplateFull

	err := updateResource(
		client,
		pageID,
		templateName,
		templateID,
		template,
		&i,
	)

	return &i, err
}

func (client *Client) DeleteTemplate(pageID, templateID string) (err error) {
	return deleteResource(client, pageID, templateName, templateID)
}
