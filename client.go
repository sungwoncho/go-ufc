package ufc

import (
	"encoding/json"
	"io/ioutil"
)

type Client struct {
	*Config
}

func New(config *Config) *Client {
	c := &Client{Config: config}
	return c
}

func getEndpoint(path string) string {
	return "http://ufc-data-api.ufc.com/api/v3/us/" + path
}

func (c *Client) Get(url string) ([]byte, error) {
	res, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, err
}

func (c *Client) Fighters() ([]Fighter, error) {
	url := getEndpoint("fighters")

	var fighters []Fighter
	body, err := c.Get(url)
	if err != nil {
		return fighters, err
	}

	err = json.Unmarshal(body, &fighters)
	if err != nil {
		return fighters, err
	}

	return fighters, nil
}
