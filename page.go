package instatus

const pageName = "page"

type Page struct {
	ID          *string `json:"id"`
	Subdomain   *string `json:"subdomain"`
	Name        *string `json:"name"`
	PublicEmail *string `json:"publicEmail"`
}

func (client *Client) CreatePage(pageID string, page *Page) (*Page, error) {
	var c Page
	err := createResource(
		client,
		pageID,
		pageName,
		page,
		&c,
	)

	return &c, err
}

func (client *Client) GetPage(pageID string) (*Page, error) {
	var c Page
	err := readPageResource(client, pageName, &c)
	return &c, err
}

func (client *Client) UpdatePage(pageID string, page *Page) (*Page, error) {
	var c Page

	err := updateResource(
		client,
		pageID,
		pageName,
		pageID,
		page,
		&c,
	)

	return &c, err
}

func (client *Client) DeletePage(pageID string) (err error) {
	return deleteResource(client, pageID, pageName, pageID)
}
