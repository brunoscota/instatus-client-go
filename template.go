package instatus

const resourceName = "template"

type Template struct {
	Name        *string      `json:"name"`
	Type        *string      `json:"type"`
	Message     *string      `json:"message"`
	MessageHtml *string      `json:"message_html,omitempty"`
	Status      *string      `json:"status"`
	Subdomain   *string      `json:"subdomain,omitempty"`
	Components  []Components `json:"components"`
	Notify      *bool        `json:"notify"`
}

type Components struct {
	ID     *string `json:"id"`
	Status *string `json:"status"`
}

type TemplateFull struct {
	IncidentTemplate
	ID        *string `json:"id"`
	CreatedAt *string `json:"created_at"`
}

func (client *Client) CreateTemplate(pageID string, template *Template) (*TemplateFull, error) {
	var i TemplateFull
	err := createResource(
		client,
		pageID,
		resourceName,
		template,
		&i,
	)

	return &i, err
}

func (client *Client) GetTemplate(pageID, templateID string) (*TemplateFull, error) {
	var i TemplateFull
	err := readResource(client, pageID, templateID, resourceName, &i)

	return &i, err
}

func (client *Client) UpdateTemplate(pageID, templateID string, template *Template) (*TemplateFull, error) {
	var i TemplateFull

	err := updateResource(
		client,
		pageID,
		resourceName,
		templateID,
		template,
		&i,
	)

	return &i, err
}

func (client *Client) DeleteTemplate(pageID, templateID string) (err error) {
	return deleteResource(client, pageID, resourceName, templateID)
}
