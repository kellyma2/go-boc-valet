package pkg

// HTTP response structures

type ListSeriesResponse struct {
	Series map[SeriesName]*SeriesInfo `json:"series"`
}

type SeriesResponse struct {
	Info *SeriesInfo `json:"seriesDetails"`
}

type ListGroupsResponse struct {
	Groups map[GroupName]*GroupInfo `json:"groups"`
}
