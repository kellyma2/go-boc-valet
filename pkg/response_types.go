package pkg

// HTTP response structures

type ListSeriesResponse struct {
	Series map[SeriesName]*SeriesInfo `json:"series"`
}

type ListGroupsResponse struct {
	Groups map[GroupName]*GroupInfo `json:"groups"`
}
