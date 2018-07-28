package client

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
	bytes2 "bytes"
)

const apiUrl = "https://partners-api.999.md"

type Client struct {
	key    string
	client *http.Client
}

func BuildClient(key string) *Client {
	return &Client{
		key: key,
		client: &http.Client{},
	}
}

func (client *Client) execute(method string, resource string, body interface{}, response interface{}) (error) {
	bytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(method, fmt.Sprintf("%s/%s", apiUrl, resource), bytes2.NewReader(bytes))
	if err != nil {
		return err
	}
	request.SetBasicAuth(client.key, "")
	request.Header.Add("Content-Type", "application/json")

	httpResponse, err := client.client.Do(request)

	if err != nil {
		return err
	}

	if httpResponse.StatusCode == http.StatusOK {
		all, err := ioutil.ReadAll(httpResponse.Body)
		if err != nil {
			return err
		}
		json.Unmarshal(all, &response)

		return nil
	}

	return fmt.Errorf("Status code %v", httpResponse.StatusCode)

}

func (client *Client) Post(resource string, body interface{}, result interface{}) error {
	return client.execute("POST", resource, body, result)
}

func (client *Client) Get(resource string, result interface{}) error {
	return client.execute("GET", resource, nil, result)
}
