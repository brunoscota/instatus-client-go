package instatus

const templateName = "template"

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
	Template
	ID        *string `json:"id"`
	CreatedAt *string `json:"created_at"`
}

func (client *Client) CreateTemplate(pageID string, template *Template) (*TemplateFull, error) {
	var i TemplateFull
	err := createResource(
		client,
		1,
		pageID,
		templateName,
		template,
		&i,
	)

	return &i, err
}

func (client *Client) GetTemplate(pageID, templateID string) (*TemplateFull, error) {
	var i TemplateFull
	err := readResource(client, 1, pageID, templateID, templateName, &i)

	return &i, err
}

func (client *Client) UpdateTemplate(pageID, templateID string, template *Template) (*TemplateFull, error) {
	var i TemplateFull

	err := updateResource(
		client,
		1,
		pageID,
		templateName,
		templateID,
		template,
		&i,
	)

	return &i, err
}

func (client *Client) DeleteTemplate(pageID, templateID string) (err error) {
	return deleteResource(client, 1, pageID, templateName, templateID)
}
