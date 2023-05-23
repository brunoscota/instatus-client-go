package instatus

const resourceName = "component"

type Component struct {
	Name        *string `json:"name"`
	Description *string `json:"description,omitempty"`
	Status      *string `json:"status,omitempty"`
	Order       *int    `json:"order,omitempty"`
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

func CreateComponent(client *Client, pageID string, component *Component) (*ComponentFull, error) {
	var c ComponentFull
	err := createResource(
		client,
		pageID,
		resourceName,
		component,
		&c,
	)

	return &c, err
}

func GetComponent(client *Client, pageID string, componentID string) (*ComponentFull, error) {
	var c ComponentFull
	err := readResource(client, pageID, componentID, resourceName, &c)

	return &c, err
}

func UpdateComponent(client *Client, pageID, componentID string, component *Component) (*ComponentFull, error) {
	var c ComponentFull

	err := updateResource(
		client,
		pageID,
		resourceName,
		componentID,
		component,
		&c,
	)

	return &c, err
}

func DeleteComponent(client *Client, pageID, componentID string) (err error) {
	return deleteResource(client, pageID, resourceName, componentID)
}
