package models

import (
	"fmt"
)

// City models a city
type City struct {
	Name       string            // Name of the city
	Destroyed  bool              // Destroyed or not
	Roads      []*Road           // Roads originating from the city
	RoadToCity map[string]*City  // Road name to city map
	RoadDir    map[string]string // Road name to direction map
}

// NewCity creates a new city
func NewCity(name string) *City {
	return &City{
		Name:       name,
		Roads:      []*Road{},
		RoadToCity: make(map[string]*City),
		RoadDir:    make(map[string]string),
	}
}

// LinkRoad links 2 cities in one direction
func (c *City) LinkRoad(road *Road, toCity *City) {
	//fmt.Printf("Updating city:%s roads\n", c.Name)
	if c.RoadToCity[road.Name] == nil {
		//fmt.Printf("Updating road: %s\n", road.Name)
		c.Roads = append(c.Roads, road)
		c.RoadToCity[road.Name] = toCity
	}
}

// String string  representation of a city
func (c *City) String() string {
	// Ignores destroyed linked cities
	var roads string
	//fmt.Printf("City Stringer called for %s\n", c.Name)
	for _, road := range c.Roads {
		toCity := c.RoadToCity[road.Name]
		//fmt.Printf("ToCity:%s, Destroyed: %t\n", toCity.Name, toCity.Destroyed)
		if toCity.Destroyed {
			continue
		}
		roads += fmt.Sprintf("%s=%s ", c.RoadDir[road.Name], toCity.Name)
	}
	if len(roads) == 0 {
		return c.Name
	}
	return fmt.Sprintf("%s %s", c.Name, roads)
}
