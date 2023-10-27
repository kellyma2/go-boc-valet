package pkg

import (
	"encoding/json"
	"fmt"
)

// data types returned by the API

type SeriesInfo struct {
	// Name of the series (can be used for lookup)
	Name string `json:"name"`

	// Series label
	Label string `json:"label"`

	// Description of the underlying series
	Description string `json:"description"`
}

func (s *SeriesInfo) PrettyPrint() {
	fmt.Printf("Name: %s\n\tLabel: %s\n\tDescription: %s\n\n",
		s.Name, s.Label, s.Description)
}

type GroupInfo struct {
	// Name of the group (can be used for lookup)
	Name string `json:"name"`

	// Group Label
	Label string `json:"label"`

	// Description of the underlying series
	Description string `json:"description"`

	// Group Series Info
	Series map[string]*SeriesInfo `json:"groupSeries"`
}

func (g *GroupInfo) PrettyPrint() {
	fmt.Printf("Name: %s\n\tLabel: %s\n\tDescription: %s\n",
		g.Name, g.Label, g.Description)
	if len(g.Series) > 0 {
		fmt.Println("Series Info:")
		for name, series := range g.Series {
			fmt.Printf("\t%s (label: %s)\n", name, series.Label)
		}
	}
	fmt.Printf("\n")
}

type DimensionInfo struct {
	// Key that the dimension is indexed on
	Key string `json:"key"`

	// Proper name of the dimension
	Name string `json:"name"`
}

func (s *DimensionInfo) PrettyPrint() {
	fmt.Printf("Dimension Info: key = %s, name = %s\n", s.Key, s.Name)
}

type SeriesObservationInfo struct {
	// Series label
	Label string `json:"label"`

	// Description of the underlying series
	Description string `json:"description"`

	// Dimension provided by the series
	Dimension DimensionInfo `json:"dimension"`
}

func (s *SeriesObservationInfo) PrettyPrint() {
	fmt.Printf("Label: %s\n", s.Label)
	fmt.Printf("Description: %s\n", s.Description)
	s.Dimension.PrettyPrint()
}

type Observation struct {
	Series string
	Key    string
	Value  string
}

func (o *Observation) PrettyPrint() {
	fmt.Printf("series = %s, key = %s, value = %s\n", o.Series, o.Key, o.Value)
}

func (o *Observation) UnmarshalJSON(b []byte) error {
	d := map[string]interface{}{}
	if err := json.Unmarshal(b, &d); err != nil {
		return err
	}

	if len(d) > 2 {
		return fmt.Errorf("Observation: dimension has unexpected keys")
	}

	for key, raw := range d {
		if key == "d" {
			if seriesKey, ok := raw.(string); ok {
				o.Key = seriesKey
			} else {
				return fmt.Errorf("Observation: key is not a string")
			}
		} else {
			o.Series = key

			if valueDict, ok := raw.(map[string]interface{}); ok {
				if value, ok := valueDict["v"]; ok {
					if valueStr, ok := value.(string); ok {
						o.Value = valueStr
					} else {
						return fmt.Errorf(
							"Observation: value not a string")
					}
				} else {
					return fmt.Errorf(
						"Observation: no series value")
				}
			} else {
				return fmt.Errorf(
					"Observation: can't convert series value")
			}
		}
	}

	return nil
}

type SeriesObservations struct {
	// Description of the series being observed
	Details map[string]*SeriesObservationInfo `json:"seriesDetail"`

	// List of data points observed
	Observations []Observation `json:"observations"`
}

func (s *SeriesObservations) PrettyPrint() {
	for name, info := range s.Details {
		fmt.Printf("Series %s\n", name)
		info.PrettyPrint()
	}

	fmt.Println("Observations:")
	for _, observation := range s.Observations {
		observation.PrettyPrint()
	}
}
