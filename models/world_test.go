package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateWorld(t *testing.T) {
	actuals := CreateWorld()
	assert.NotEqual(t, nil, actuals)
}

func TestCreateCity(t *testing.T) {
	world := CreateWorld()
	assert.NotEqual(t, nil, world)
	city := world.CreateCity("foo")
	assert.NotEqual(t, nil, city)
	assert.Equal(t, "foo", city.Name)
}

func TestGetUndestroyedCities(t *testing.T) {
	world := CreateWorld()
	assert.NotNil(t, world)
	city := world.CreateCity("foo")
	assert.NotNil(t, city)
	assert.Equal(t, "foo", city.Name)

	str1 := world.GetUndestroyedCities()

	assert.Equal(t, "foo\n", str1)

	city2 := world.CreateCity("bar")
	assert.NotNil(t, city2)
	assert.Equal(t, "bar", city2.Name)
	city2.Destroyed = true

	str2 := world.GetUndestroyedCities()
	assert.Equal(t, "foo\n", str2)
}
