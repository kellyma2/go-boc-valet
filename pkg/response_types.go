package pkg

// HTTP response structures

type ListSeriesResponse struct {
	Series map[string]*SeriesInfo `json:"series"`
}

type SeriesResponse struct {
	Info *SeriesInfo `json:"seriesDetails"`
}

type ListGroupsResponse struct {
	Groups map[string]*GroupInfo `json:"groups"`
}

type GroupResponse struct {
	Info *GroupInfo `json:"groupDetails"`
}
