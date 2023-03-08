// Package cmdargs parses the command line args.
// Uses pflag go package for command line parsing.
package cmdargs

import (
	"os"

	"github.com/spf13/pflag"
)

// CmdLineArgs interface for retrieving command line args
type CmdLineArgs interface {
	// GetWorldMap returns the world map containing cities and paths
	GetWorldMap() string
	// GetNumberOfAliens returns the number of aliens
	GetNumberOfAliens() int
	// GetMaxAlienMoves returns the max number of alien moves
	GetMaxAlienMoves() int
}

// cmdArguments
type cmdArguments struct {
	numAliens     int    // Number of aliens
	maxAlienMoves int    // Max alien moves
	worldMap      string // Text file containing cities and directions
	help          bool   // Help
}

// New parses command line arguments and initializes
// cmdArguments. Prints usage if parsing fails
func New() CmdLineArgs {
	c := &cmdArguments{}
	pflag.CommandLine.AddFlagSet(c.FlagSet())
	pflag.Parse()
	if c.worldMap == "" || c.help {
		pflag.Usage()
		os.Exit(2)
	}
	return c
}

// FlagSet return a *pflag.FlagSet that can be used to register CLI options to configure logging
func (c *cmdArguments) FlagSet() *pflag.FlagSet {
	var set pflag.FlagSet
	set.StringVarP(&c.worldMap, "worldmap", "m", "", "world map file (Mandatory)")
	set.IntVarP(&c.numAliens, "numaliens", "n", 10, "number of aliens")
	set.IntVarP(&c.maxAlienMoves, "maxAlienMoves", "i", 10000, "max alien moves")
	set.BoolVarP(&c.help, "help", "h", false, "print help and exit")
	return &set
}

// GetNumberOfAliens returns the number of aliens
func (c *cmdArguments) GetNumberOfAliens() int { return c.numAliens }

// GetURL returns the world map containing cities and paths
func (c *cmdArguments) GetWorldMap() string { return c.worldMap }

// GetNumberOfAliens returns the max number of alien moves
func (c *cmdArguments) GetMaxAlienMoves() int { return c.numAliens }
