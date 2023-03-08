// Package main ...
package main

import (
	"alien-invasion/cmdargs"
	"alien-invasion/invasion"
	"fmt"
	"os"
)

// VersionString and BuildTime are used to inject the version and buildtime
// into the binary through LDFLAGS -X option during build time
var VersionString = "development"
var BuildTime = ""

func main() {
	fmt.Printf("SW version: %s buildtime: %s\n", VersionString, BuildTime)
	// Parse command line arguments
	c := cmdargs.New()
	// Create invasion instance
	i, err := invasion.NewInvasion(c.GetWorldMap(), c.GetMaxAlienMoves())
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		os.Exit(1)
	}
	// Create aliens
	i.CreateAliens(c.GetNumberOfAliens())
	// Read the world map
	if err := i.ReadWorldMap(); err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		os.Exit(1)
	}
	// Unleash the aliens into the world
	if err := i.UnleashAliens(); err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("Invasion Success!!")
	// Write the remains
	i.WriteRemains()
}
