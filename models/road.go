package models

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// Road models a road
type Road struct {
	Name       string   // Name of the road. e.g: Bar-Foo
	ConnCities []string // Connecting cities
}

// String representation
func (r Road) String() string {
	return fmt.Sprintf("Name=%s ConnCities=%s\n", r.Name, r.ConnCities)
}

// NewRoad creates a new road between 2 cities
func NewRoad(cities ...string) (*Road, error) {
	if len(cities) < 2 || len(cities) > 2 {
		return nil, errors.New("only-2-cities-supported")
	}
	sort.Strings(cities)
	name := strings.Join(cities, "-")
	return &Road{
		Name:       name,
		ConnCities: cities,
	}, nil
}
