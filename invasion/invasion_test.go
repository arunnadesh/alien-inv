package invasion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAliens(t *testing.T) {
	i, _ := NewInvasion("../tests/default_map.txt", 10000)
	assert.NotNil(t, i)
	i.CreateAliens(10)
	inv, ok := i.(*invasion)
	assert.True(t, ok)
	assert.Equal(t, 10, len(inv.aliens))
	assert.Equal(t, 10000, inv.maxIterations)
}

func TestReadWorldMap(t *testing.T) {
	// Foo north=Bar west=Baz south=Qu-ux
	// Bar south=Foo west=Bee
	i, _ := NewInvasion("../tests/default_map.txt", 10000)
	assert.NotNil(t, i)
	i.CreateAliens(10)
	inv, ok := i.(*invasion)
	assert.True(t, ok)
	assert.Equal(t, 10, len(inv.aliens))
	assert.Equal(t, 10000, inv.maxIterations)

	err := i.ReadWorldMap()
	assert.Nil(t, err)
	// Foo, Bar, Baz, Qu-ux, Bee
	assert.Equal(t, 5, len(inv.World.Cities))
	// check road-name and direction mappings
	assert.Equal(t, "north", inv.World.Cities["Foo"].RoadDir["Bar-Foo"])
	assert.Equal(t, "south", inv.World.Cities["Bar"].RoadDir["Bar-Foo"])
	assert.Equal(t, "west", inv.World.Cities["Foo"].RoadDir["Baz-Foo"])
	assert.Equal(t, "east", inv.World.Cities["Baz"].RoadDir["Baz-Foo"])
	assert.Equal(t, "south", inv.World.Cities["Foo"].RoadDir["Foo-Qu-ux"])
	assert.Equal(t, "north", inv.World.Cities["Qu-ux"].RoadDir["Foo-Qu-ux"])
	// check road-name and direction city mapping
	assert.Equal(t, "Bar", inv.World.Cities["Foo"].RoadToCity["Bar-Foo"].Name)
	assert.Equal(t, "Foo", inv.World.Cities["Bar"].RoadToCity["Bar-Foo"].Name)
	assert.Equal(t, "Baz", inv.World.Cities["Foo"].RoadToCity["Baz-Foo"].Name)
	assert.Equal(t, "Foo", inv.World.Cities["Baz"].RoadToCity["Baz-Foo"].Name)
	assert.Equal(t, "Qu-ux", inv.World.Cities["Foo"].RoadToCity["Foo-Qu-ux"].Name)
	assert.Equal(t, "Foo", inv.World.Cities["Qu-ux"].RoadToCity["Foo-Qu-ux"].Name)
}

func TestOppositeDirections(t *testing.T) {
	assert.Equal(t, "north", oppositeDir("south"))
	assert.Equal(t, "south", oppositeDir("north"))
	assert.Equal(t, "east", oppositeDir("west"))
	assert.Equal(t, "west", oppositeDir("east"))
}
