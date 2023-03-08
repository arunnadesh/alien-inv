package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkRoad(t *testing.T) {
	// Create city
	foo := NewCity("foo")
	assert.NotNil(t, foo)
	assert.Equal(t, "foo", foo.Name)
	bar := NewCity("bar")
	assert.NotNil(t, bar)
	assert.Equal(t, "bar", bar.Name)
	// Create road
	road, _ := NewRoad("foo", "bar")
	assert.NotNil(t, road)
	assert.NotEqual(t, "foo-bar", road.Name)
	assert.Equal(t, "bar-foo", road.Name)
	// Link roads between the cities
	foo.LinkRoad(road, bar)
	assert.NotNil(t, foo.Roads)
	assert.Equal(t, road, foo.Roads[0])
	assert.Equal(t, bar, foo.RoadToCity[road.Name])
	bar.LinkRoad(road, foo)
	assert.NotNil(t, bar.Roads)
	assert.Equal(t, road, bar.Roads[0])
	assert.Equal(t, foo, bar.RoadToCity[road.Name])
}
