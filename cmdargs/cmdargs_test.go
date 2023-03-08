package cmdargs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigParsingDefaults(t *testing.T) {
	var cfg cmdArguments
	assert.NoError(t, cfg.FlagSet().Parse(nil))
	assert.Equal(t, 10000, cfg.maxAlienMoves)
	assert.Equal(t, 10, cfg.GetNumberOfAliens())
}
