package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsOkay(t *testing.T) {
	// Alien yet to start invasion
	alien1 := NewAlien(10)
	assert.NotNil(t, alien1)
	assert.Equal(t, 10, alien1.ID)
	assert.True(t, alien1.IsOkay())
	// Alien dead
	alien2 := NewAlien(20)
	assert.NotNil(t, alien2)
	assert.Equal(t, 20, alien2.ID)
	alien2.Dead = true
	assert.False(t, alien2.IsOkay())
	// Alien trapped
	foo := NewCity("foo")
	assert.NotNil(t, foo)
	assert.Equal(t, "foo", foo.Name)
	alien3 := NewAlien(30)
	assert.NotNil(t, alien3)
	assert.Equal(t, 30, alien3.ID)
	alien3.CurrCity = foo
	assert.False(t, alien2.IsOkay())

}
