package pkg

import "fmt"

// data types returned by the API

type SeriesName string

type SeriesInfo struct {
	// Name of the series (can be used for lookup)
	Name SeriesName `json:"name"`

	// Series label
	Label string `json:"label"`

	// Description of the underlying series
	Description string `json:"description"`
}

func (s *SeriesInfo) PrettyPrint() {
	fmt.Printf("Name: %s\n\tLabel: %s\n\tDescription: %s\n\n",
		s.Name, s.Label, s.Description)
}

type GroupName string

type GroupInfo struct {
	// Name of the group (can be used for lookup)
	Name GroupName `json:"name"`

	// Group Label
	Label string `json:"label"`

	// Description of the underlying series
	Description string `json:"description"`
}

func (g *GroupInfo) PrettyPrint() {
	fmt.Printf("Name: %s\n\tLabel: %s\n\tDescription: %s\n\n",
		g.Name, g.Label, g.Description)
}
