package models

import (
	"fmt"
)

// World models the worls
type World struct {
	Cities map[string]*City // Cities of the world
}

// CreateWorld creates an empty world
func CreateWorld() *World {
	return &World{Cities: make(map[string]*City)}
}

// CreateCity creates a city with a given name in the world
func (w *World) CreateCity(name string) *City {
	city := NewCity(name)
	w.Cities[city.Name] = city
	return w.Cities[city.Name]
}

// String  string representation of a city
func (w *World) String() string {
	var cities string
	for _, city := range w.Cities {
		cities += fmt.Sprintf("%s\n", city)
	}
	return cities
}

// GetUndestroyedCities gets the undestroyed cities
func (w *World) GetUndestroyedCities() string {
	out := ""
	for _, city := range w.Cities {
		if city.Destroyed {
			continue
		}
		out += fmt.Sprintf("%s\n", city)
	}
	return out
}
