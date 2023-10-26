package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Implementation of our underlying API access

const apiRoot = "https://www.bankofcanada.ca/valet/"

type ValetClient struct {
	apiRoot string
}

func NewValetClient() ValetClient {
	return ValetClient{
		apiRoot: apiRoot,
	}
}

func (c *ValetClient) urlFor(suffix ...string) (string, error) {
	return url.JoinPath(c.apiRoot, suffix...)
}

func (c *ValetClient) getUrlFor(r interface{}, suffix ...string) error {
	u, err := c.urlFor(suffix...)
	if err != nil {
		return fmt.Errorf("joining URL: %w", err)
	}

	resp, err := http.Get(u)
	if err != nil {
		return fmt.Errorf("GETting: %w", err)
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return fmt.Errorf("Unmarshalling JSON: %w", err)
	}

	return nil
}

func (c *ValetClient) SeriesList() (map[SeriesName]*SeriesInfo, error) {
	r := ListSeriesResponse{}

	if err := c.getUrlFor(&r, "lists", "series"); err != nil {
		return r.Series, err
	}

	// the response doesn't inject the name into the JSON dict
	// so we do it instead to keep ourselves a bit more sane
	for name, info := range r.Series {
		info.Name = name
	}
	return r.Series, nil
}

func (c *ValetClient) GroupList() (map[GroupName]*GroupInfo, error) {
	r := ListGroupsResponse{}

	if err := c.getUrlFor(&r, "lists", "groups"); err != nil {
		return r.Groups, err
	}

	// the response doesn't inject the name into the JSON dict
	// so we do it instead to keep ourselves a bit more sane
	for name, info := range r.Groups {
		info.Name = name
	}
	return r.Groups, nil
}

func (c *ValetClient) Series(name string) (*SeriesInfo, error) {
	r := SeriesResponse{}
	return r.Info, c.getUrlFor(&r, "series", name)
}

func (c *ValetClient) Group(name string) (*GroupInfo, error) {
	r := GroupResponse{}
	return r.Info, c.getUrlFor(&r, "groups", name)
}
