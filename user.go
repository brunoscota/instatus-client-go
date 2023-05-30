package instatus

type User struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Avatar string `json:"avatar"`
}

func (client *Client) GetUser() (*User, error) {
	var u User
	err := readResourceCustomURL(client, "/user", "user", &u)
	return &u, err
}
