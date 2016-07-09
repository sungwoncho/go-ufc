package ufc

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
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

func (c *Client) GetResponse(url string, i interface{}) error {
	body, err := c.Get(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, i)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Fighters() ([]Fighter, error) {
	url := getEndpoint("fighters")
	var fighters []Fighter

	err := c.GetResponse(url, &fighters)
	if err != nil {
		return nil, err
	}

	return fighters, nil
}

func (c *Client) Fighter(id int) (Fighter, error) {
	url := getEndpoint("fighters/" + strconv.Itoa(id) + ".json")
	var fighter Fighter

	err := c.GetResponse(url, &fighter)
	if err != nil {
		return fighter, err
	}

	return fighter, nil
}
