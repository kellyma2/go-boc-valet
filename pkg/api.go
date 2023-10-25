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

func (c *ValetClient) SeriesList() (map[SeriesName]*SeriesInfo, error) {
	r := ListSeriesResponse{}

	u, err := c.urlFor("/lists/series")
	if err != nil {
		return r.Series, fmt.Errorf("joining URL: %w", err)
	}

	resp, err := http.Get(u)
	if err != nil {
		return r.Series, fmt.Errorf("GETting: %w", err)
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return r.Series, fmt.Errorf("Unmarshalling JSON: %w", err)
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

	u, err := c.urlFor("/lists/groups")
	if err != nil {
		return r.Groups, fmt.Errorf("joining URL: %w", err)
	}

	resp, err := http.Get(u)
	if err != nil {
		return r.Groups, fmt.Errorf("GETting: %w", err)
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return r.Groups, fmt.Errorf("Unmarshalling JSON: %w", err)
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

	u, err := c.urlFor("series", name)
	if err != nil {
		return r.Info, fmt.Errorf("joining URL: %w", err)
	}

	resp, err := http.Get(u)
	if err != nil {
		return r.Info, fmt.Errorf("GETting: %w", err)
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return r.Info, fmt.Errorf("Unmarshalling JSON: %w", err)
	}

	return r.Info, nil
}
