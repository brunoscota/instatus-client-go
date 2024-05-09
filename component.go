package instatus

import "time"

const componentName = "component"

type Component struct {
	ID                         *string            `json:"id"`
	Name                       *string            `json:"name,omitempty"`
	NameTranslationID          *string            `json:"nameTranslationId,omitempty"`
	Description                *string            `json:"description,omitempty"`
	DescriptionTranslationID   *string            `json:"descriptionTranslationId,omitempty"`
	Status                     *string            `json:"status,omitempty"`
	InternalStatus             *string            `json:"internalStatus,omitempty"`
	Order                      *int               `json:"order,omitempty"`
	ShowUptime                 *bool              `json:"showUptime,omitempty"`
	CreatedAt                  *time.Time         `json:"createdAt,omitempty"`
	UpdatedAt                  *time.Time         `json:"updatedAt,omitempty"`
	ArchivedAt                 *time.Time         `json:"archivedAt,omitempty"`
	SiteID                     *string            `json:"siteId,omitempty"`
	UniqueEmail                *string            `json:"uniqueEmail,omitempty"`
	OldGroup                   *string            `json:"oldGroup,omitempty"`
	GroupID                    *string            `json:"groupId,omitempty"`
	IsParent                   *bool              `json:"isParent,omitempty"`
	IsCollapsed                *bool              `json:"isCollapsed,omitempty"`
	MonitorID                  *string            `json:"monitorId,omitempty"`
	NameHTML                   *string            `json:"nameHtml,omitempty"`
	NameHTMLTranslationID      *string            `json:"nameHtmlTranslationId,omitempty"`
	DescriptionHTML            *string            `json:"descriptionHtml,omitempty"`
	DescriptionHTMLTranslation *string            `json:"descriptionHtmlTranslationId,omitempty"`
	IsThirdParty               *bool              `json:"isThirdParty,omitempty"`
	ThirdPartyStatus           *string            `json:"thirdPartyStatus,omitempty"`
	ThirdPartyComponentID      *string            `json:"thirdPartyComponentId,omitempty"`
	ThirdPartyComponentService *string            `json:"thirdPartyComponentServiceId,omitempty"`
	ImportedFromStatuspage     *bool              `json:"importedFromStatuspage,omitempty"`
	StartDate                  *time.Time         `json:"startDate,omitempty"`
	Group                      *string            `json:"group,omitempty"`
	Translations               *ComponentLanguage `json:"translations,omitempty"`
}

type ComponentLanguage struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type Group struct {
	Name *string `json:"name,omitempty"`
}

func (client *Client) CreateComponent(pageID string, component *Component) (*Component, error) {
	var c Component
	err := createResource(
		client,
		pageID,
		componentName,
		component,
		&c,
	)

	return &c, err
}

func (client *Client) GetComponent(pageID string, componentID string) (*Component, error) {
	var c Component
	err := readResource(client, pageID, componentID, componentName, &c)

	return &c, err
}

func (client *Client) UpdateComponent(pageID string, componentID string, component *Component) (*Component, error) {
	var c Component

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
