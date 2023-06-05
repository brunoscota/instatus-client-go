package instatus

const componentName = "component"

type Component struct {
	Name        *string `json:"name"`
	Description *string `json:"description,omitempty"`
	Status      *string `json:"status,omitempty"`
	Order       *int64  `json:"order,omitempty"`
	ShowUptime  *bool   `json:"showUptime,omitempty"`
	Grouped     *bool   `json:"grouped,omitempty"`
	Group       *string `json:"group,omitempty"`
}

type Group struct {
	Name *string `json:"name,omitempty"`
}

type ComponentFull struct {
	Component
	ID          *string `json:"id"`
	UniqueEmail *string `json:"uniqueEmail,omitempty"`
	GroupID     *string `json:"groupId,omitempty"`
	Group       Group   `json:"group,omitempty"`
}

func (client *Client) CreateComponent(pageID string, component *Component) (*ComponentFull, error) {
	var c ComponentFull
	err := createResource(
		client,
		pageID,
		componentName,
		component,
		&c,
	)

	return &c, err
}

func (client *Client) GetComponent(pageID string, componentID string) (*ComponentFull, error) {
	var c ComponentFull
	err := readResource(client, pageID, componentID, componentName, &c)

	return &c, err
}

func (client *Client) UpdateComponent(pageID string, componentID string, component *Component) (*ComponentFull, error) {
	var c ComponentFull

	err := updateResource(
		client,
		pageID,
		componentName,
		componentID,
		component,
		&c,
	)

	return &c, err
}

func (client *Client) DeleteComponent(pageID string, componentID string) (err error) {
	return deleteResource(client, pageID, componentName, componentID)
}
