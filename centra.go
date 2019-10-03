package centra

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// Client hold the information required to contact Centra.
type Client struct {
	url    string
	secret string
	client *http.Client
}

// Init initiates the credentials for communicating with Centra.
func Init(url string, secret string) *Client {
	return &Client{
		url:    url,
		secret: secret,
		client: &http.Client{},
	}
}

func (c *Client) get(path string, target interface{}) (statusCode int, e error) {
	statusCode, e = c.request(http.MethodGet, path, &target)
	return
}

func (c *Client) request(method, path string, target interface{}) (statusCode int, e error) {
	req, e := http.NewRequest(method, c.url+path, nil)
	if e != nil {
		return
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Bearer "+c.secret)

	response, e := c.client.Do(req)
	statusCode = response.StatusCode
	if e != nil || statusCode >= 300 {
		e = errors.New(response.Status)
		return
	}

	defer response.Body.Close()

	body, e := ioutil.ReadAll(response.Body)
	if e != nil {
		return
	}

	e = json.Unmarshal(body, &target)
	return
}
