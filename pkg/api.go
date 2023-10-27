package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
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

func (c *ValetClient) doGet(r interface{}, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("GETting: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GET failed: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return fmt.Errorf("Unmarshalling JSON: %w", err)
	}

	return nil
}

func (c *ValetClient) getQueryUrlFor(
	r interface{},
	q url.Values,
	suffix ...string,
) error {
	u, err := c.urlFor(suffix...)
	if err != nil {
		return fmt.Errorf("joining URL: %w", err)
	}

	return c.doGet(r, u+"?"+q.Encode())
}

func (c *ValetClient) getUrlFor(r interface{}, suffix ...string) error {
	u, err := c.urlFor(suffix...)
	if err != nil {
		return fmt.Errorf("joining URL: %w", err)
	}

	return c.doGet(r, u)
}

func (c *ValetClient) SeriesList() (map[string]*SeriesInfo, error) {
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

func (c *ValetClient) GroupList() (map[string]*GroupInfo, error) {
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

type ObservationOption func(o *observationOptions)

type observationOptions struct {
	series             []string
	recent             int
	recentWeeks        int
	recentMonths       int
	recentYears        int
	startDate, endDate string
}

func WithSeries(series string) ObservationOption {
	return func(o *observationOptions) {
		o.series = append(o.series, series)
	}
}

func WithRecent(recent int) ObservationOption {
	return func(o *observationOptions) {
		o.recent = recent
	}
}

func WithStartDate(start string) ObservationOption {
	return func(o *observationOptions) {
		o.startDate = start
	}
}

func WithEndDate(end string) ObservationOption {
	return func(o *observationOptions) {
		o.endDate = end
	}
}

func (c *ValetClient) SeriesObservations(
	opts ...ObservationOption,
) (SeriesObservations, error) {
	configuredOpts := &observationOptions{}
	for _, opt := range opts {
		opt(configuredOpts)
	}

	if len(configuredOpts.series) == 0 {
		return SeriesObservations{}, nil
	}

	suffix := strings.Join(configuredOpts.series, ",")
	queryParams := url.Values{}

	if configuredOpts.recent != 0 {
		queryParams.Add("recent", fmt.Sprintf("%d", configuredOpts.recent))
	}

	if configuredOpts.recentWeeks != 0 {
		queryParams.Add(
			"recent_weeks",
			fmt.Sprintf("%d", configuredOpts.recentWeeks),
		)
	}

	if configuredOpts.recentMonths != 0 {
		queryParams.Add(
			"recent_months",
			fmt.Sprintf("%d", configuredOpts.recentMonths),
		)
	}

	if configuredOpts.recentYears != 0 {
		queryParams.Add(
			"recent_years",
			fmt.Sprintf("%d", configuredOpts.recentYears),
		)
	}

	if len(configuredOpts.startDate) > 0 {
		queryParams.Add("start_date", configuredOpts.startDate)
	}

	if len(configuredOpts.endDate) > 0 {
		queryParams.Add("end_date", configuredOpts.endDate)
	}

	r := SeriesObservations{}
	if len(queryParams) > 0 {
		return r, c.getQueryUrlFor(&r, queryParams, "observations", suffix)
	} else {
		return r, c.getUrlFor(&r, "observations", suffix)
	}
}
