package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRoad(t *testing.T) {
	// Create road tests
	actuals, _ := NewRoad("foo", "bar")
	assert.NotNil(t, actuals)
	assert.NotEqual(t, "foo-bar", actuals.Name)
	assert.Equal(t, "bar-foo", actuals.Name)

	actuals1, err := NewRoad("foo")
	assert.Nil(t, actuals1)
	assert.Equal(t, "only-2-cities-supported", err.Error())

	actuals2, err := NewRoad("foo", "bar", "zoo")
	assert.Nil(t, actuals2)
	assert.Equal(t, "only-2-cities-supported", err.Error())
}
