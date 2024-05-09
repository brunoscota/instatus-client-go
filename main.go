package instatus

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const apiRoot = "https://api.instatus.com/v2"

// HTTPClient is the http wrapper for the application
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type IClient interface {
	doHTTPRequest(method, endpoint string, item interface{}) (resp *http.Response, err error)
}

type Client struct {
	apiKey     string
	httpClient HTTPClient
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
}

// Allows overriding the HTTP Client, leaving the choice
// of using a retry library to the user
func (client *Client) UseHTTPClient(httpClient HTTPClient) {
	client.httpClient = httpClient
}

func (client *Client) doHTTPRequest(method, endpoint string, item interface{}) (resp *http.Response, err error) {
	url := apiRoot + endpoint

	var body io.Reader

	if item != nil {
		data, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}

		body = strings.NewReader(string(data))
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+client.apiKey)
	req.Header.Set("Content-Type", "application/json")

	maxRetries := 10
	retryInterval := 30 * time.Second

	// Basic Retry logic around rate limiting
	resp, err = client.httpClient.Do(req)
	retries := 0
	for retries = 1; resp != nil && resp.StatusCode == 429 && retries <= maxRetries; retries = retries + 1 {
		time.Sleep(retryInterval)
		resp, err = client.httpClient.Do(req)
	}

	return resp, err
}

func createPageResource(client IClient, resource, result interface{}) error {
	return createResourceCustomURL(client, "/pages", resource, result)
}

func createResource(client IClient, pageID, resourceType string, resource, result interface{}) error {
	return createResourceCustomURL(client, "/"+pageID+"/"+resourceType+"s", resource, result)
}

func createResourceCustomURL(client IClient, URL string, resource, result interface{}) error {
	resp, err := client.doHTTPRequest(
		"POST",
		URL,
		resource,
	)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return json.Unmarshal(bodyBytes, &result)
	}

	return fmt.Errorf("failed creating resource, request returned %d, full response: %+v", resp.StatusCode, resp)
}

func readResourceCustomURL(client IClient, url string, errorMessage string, target interface{}) error {
	resp, err := client.doHTTPRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		return err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	switch resp.StatusCode {
	case http.StatusOK:
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return json.Unmarshal(bodyBytes, target)

	case http.StatusNotFound:
		return nil

	default:
		return fmt.Errorf("could not find %s, http status %d", errorMessage, resp.StatusCode)
	}
}

func readPageResource(client IClient, ID string, target interface{}) error {
	return readResourceCustomURL(
		client, "/pages?page="+ID,
		fmt.Sprintf("Page with ID: %s", ID),
		target,
	)
}

func readResource(client IClient, pageID string, ID string, resourceType string, target interface{}) error {
	return readResourceCustomURL(
		client, "/"+pageID+"/"+resourceType+"s/"+ID,
		fmt.Sprintf("%s with ID: %s", resourceType, ID),
		target,
	)
}

func updateResourceCustomURL(client IClient, url string, errorMessage string, resource interface{}, result interface{}) error {
	resp, err := client.doHTTPRequest(
		"PUT",
		url,
		resource,
	)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return json.Unmarshal(bodyBytes, &result)
	}

	return fmt.Errorf("failed updating %s, request returned %d", errorMessage, resp.StatusCode)
}

func updateResource(client IClient, pageID string, resourceType string, ID string, resource interface{}, result interface{}) error {
	return updateResourceCustomURL(
		client,
		"/"+pageID+"/"+resourceType+"s/"+ID,
		fmt.Sprintf("%s with ID: %s", resourceType, ID),
		resource,
		result,
	)
}

func deleteResourceCustomURL(client IClient, url string, errorMessage string) error {
	resp, err := client.doHTTPRequest(
		"DELETE",
		url,
		nil,
	)
	if err != nil {
		return err
	}

	// StatusGroup deletion returns StatusOK instead of StatusNoContent like other resources
	if resp.StatusCode == http.StatusNoContent || resp.StatusCode == http.StatusOK {
		return nil
	}

	return fmt.Errorf("failed deleting %s, request returned %d", errorMessage, resp.StatusCode)
}

func deleteResource(client IClient, pageID string, resourceType string, ID string) error {
	return deleteResourceCustomURL(client, "/"+pageID+"/"+resourceType+"s/"+ID, resourceType)
}
