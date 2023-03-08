// Package models model the alien, city, road and the world
package models

import (
	"fmt"
	"strconv"
)

// Alien contains alien info
type Alien struct {
	ID       int   // ID of the alien
	Dead     bool  // Dead or alive
	CurrCity *City // Current city
}

// NewAlien creates a new alien with the given id
func NewAlien(id int) *Alien {
	return &Alien{
		ID: id,
	}
}

// IsOkay checks if the alien is alive or not-trapped
func (a *Alien) IsOkay() bool {
	// check if the alien is dead
	if a.Dead {
		return false
	}
	// check if the alien started invasion
	if a.CurrCity == nil {
		return true
	}
	// if in a city check if all connected cities are destroyed
	for _, city := range a.CurrCity.RoadToCity {
		if !city.Destroyed {
			return true
		}
	}
	return false
}

// String representation of an alien
func (a *Alien) String() string {
	return fmt.Sprintf("id=%s city={%s}\n", strconv.Itoa(a.ID), a.CurrCity)
}
