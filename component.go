package instatus

const componentName = "component"

type Component struct {
	Name        *string `json:"name"`
	Description *string `json:"description,omitempty"`
	Status      *string `json:"status,omitempty"`
	Order       *int64  `json:"order,omitempty"`
	GroupID     *string `json:"group_id,omitempty"`
	ShowUptime  *bool   `json:"show_uptime,omitempty"`
	Grouped     *bool   `json:"grouped,omitempty"`
	Group       *string `json:"group,omitempty"`
}

type ComponentFull struct {
	Component
	ID          *string `json:"id"`
	UniqueEmail *string `json:"unique_email,omitempty"`
}

func (client *Client) CreateComponent(pageID string, component *Component) (*ComponentFull, error) {
	var c ComponentFull
	err := createResource(
		client,
		1,
		pageID,
		componentName,
		component,
		&c,
	)

	return &c, err
}

func (client *Client) GetComponent(pageID string, componentID string) (*ComponentFull, error) {
	var c ComponentFull
	err := readResource(client, 1, pageID, componentID, componentName, &c)

	return &c, err
}

func (client *Client) UpdateComponent(pageID string, componentID string, component *Component) (*ComponentFull, error) {
	var c ComponentFull

	err := updateResource(
		client,
		1,
		pageID,
		componentName,
		componentID,
		component,
		&c,
	)

	return &c, err
}

func (client *Client) DeleteComponent(pageID string, componentID string) (err error) {
	return deleteResource(client, 1, pageID, componentName, componentID)
}
