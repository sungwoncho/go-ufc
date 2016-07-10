package ufc

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Client struct {
	BaseUrl    string
	HTTPClient *http.Client
}

func New() *Client {
	c := &Client{
		BaseUrl:    "http://ufc-data-api.ufc.com/api/v3/us/",
		HTTPClient: http.DefaultClient,
	}

	return c
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
	url := c.BaseUrl + "fighters"
	var fighters []Fighter

	err := c.GetResponse(url, &fighters)
	if err != nil {
		return nil, err
	}

	return fighters, nil
}

func (c *Client) Fighter(id int) (Fighter, error) {
	url := c.BaseUrl + "fighters/" + strconv.Itoa(id) + ".json"
	var fighter Fighter

	err := c.GetResponse(url, &fighter)
	if err != nil {
		return fighter, err
	}

	return fighter, nil
}

func (c *Client) Events() ([]Event, error) {
	url := c.BaseUrl + "events"
	var events []Event

	err := c.GetResponse(url, &events)
	if err != nil {
		return events, err
	}

	return events, nil
}

func (c *Client) Event(id int) (Event, error) {
	url := c.BaseUrl + "events/" + strconv.Itoa(id) + ".json"
	var event Event

	err := c.GetResponse(url, &event)
	if err != nil {
		return event, err
	}

	return event, nil
}
