// Package invasion simulates the invasion
package invasion

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"alien-invasion/models"
)

// Invasion provides interfaces for the invasion
type Invasion interface {
	// CreateAliens n number of aliens from ids 0....n
	CreateAliens(n int)
	// ReadWorldMap reads world map from the given input file
	ReadWorldMap() error
	// UnleashAliens unleashes the aliens into the world
	UnleashAliens() error
	// WriteRemains prints the remains of the world on the console
	WriteRemains()
}

// invasion structure to hold the cities and aliens info
type invasion struct {
	randGen       *rand.Rand               // random seed
	mapFile       string                   // world map file path
	World         *models.World            // world made of cities connected by roads
	aliens        []*models.Alien          // aliens
	alienPresence map[string]*models.Alien // keeps track of alien presence in a city
	maxIterations int                      // max iterations
}

// NewInvasion creates a new instance of the invasion
func NewInvasion(fname string, maxIter int) (Invasion, error) {
	return &invasion{
		randGen:       rand.New(rand.NewSource(time.Now().UnixNano())),
		mapFile:       fname,
		World:         models.CreateWorld(),
		maxIterations: maxIter,
		alienPresence: make(map[string]*models.Alien),
	}, nil
}

// ReadWorldMap reads world map from the give input file
func (i *invasion) ReadWorldMap() error {
	file, err := os.Open(i.mapFile)
	if err != nil {
		return err
	}
	defer file.Close()
	// Read world map line by line
	// Assumptions:
	//    Consider the example,
	// 		Foo north=Bar west=Baz south=Qu-ux
	//  	Bar south=Foo west=Bee
	// 	 1. One entry per city.
	// 	 2. Space is used as delimiter to parse cities and links.
	//      So assuming no space in the names of cities.
	// 	 3. No leading or trailing spaces around "=" sign.
	// 	 4. City and Direction names are case sensitive.
	//   5. In the above example "Baz" and "Bee" are present in the links.
	//      But there is no city entry for "Baz" an "Bee". In such cases,
	//      We will explicitly add city entries in the "World". For instance,
	//      "Baz" will be added with a link having direction "west" with city "Foo"
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		// split each line entry by space, the first entry in the slice
		// is the current city
		entry := strings.Split(scanner.Text(), " ")
		// Create city in the world if no already present.
		// Note: We process the linked cities as well. So
		// current city might have been added already while
		// processing the links of a previous city.
		fmt.Printf("Processing  city:%s\n", entry[0])
		var city *models.City
		if c, present := i.World.Cities[entry[0]]; !present {
			city = i.World.CreateCity(entry[0])
		} else {
			city = c
		}
		// Connect the linked cities. Create linked cities if not present
		for _, link := range entry[1:] {
			if err := i.connectLinkedCities(city, link); err != nil {
				return err
			}
		}
	}

	return nil
}

// connectLinkedCities connects the linked cities
func (i *invasion) connectLinkedCities(city *models.City, link string) error {
	// Parse the links
	direction, cityName, err := parseNeighbor(link)
	if err != nil {
		return err
	}
	// Check if the adjacent city is already created.
	adjCity, present := i.World.Cities[cityName]
	if !present {
		adjCity = i.World.CreateCity(cityName)
	}
	// Create road connecting 2 cities
	road, err := models.NewRoad(city.Name, adjCity.Name)
	if err != nil {
		return err
	}
	fmt.Printf("City: %s, adjCity: %s, RoadName: %s\n", city.Name, adjCity.Name, road.Name)
	// Link cities at both the end of a road
	city.LinkRoad(road, adjCity)
	adjCity.LinkRoad(road, city)
	// Populate road-name vs direction map
	city.RoadDir[road.Name] = direction
	adjCity.RoadDir[road.Name] = oppositeDir(direction)
	return nil
}

// CreateAliens creates n new Aliens with ids from 0....n
func (i *invasion) CreateAliens(n int) {
	i.aliens = make([]*models.Alien, n)
	for j := 0; j < n; j++ {
		alien := models.NewAlien(j)
		i.aliens[j] = alien
	}
}

// parseNeighbor gets the direction and adjacent city
func parseNeighbor(adj string) (string, string, error) {
	entries := strings.Split(adj, "=")
	if len(entries) != 2 {
		return "", "", errors.New("invalid-adjacent-format")
	}
	return entries[0], entries[1], nil
}

// UnleashAliens unleashes the aliens into the world
func (i *invasion) UnleashAliens() error {
	fmt.Printf("Unleashing the aliens!! Max-iterations: %d\n", i.maxIterations)
	fmt.Println("==================")
	fmt.Println("Original map before attack")
	fmt.Println("==================")
	fmt.Print(i.World.GetUndestroyedCities())
	fmt.Println("==================")
	for j := 0; j < i.maxIterations; j++ {
		fmt.Printf("Iteration: %d\n", j)
		fmt.Println("==================")
		if err := i.charge(); err != nil {
			if _, ok := err.(*models.InvasionError); ok {
				fmt.Printf("Game over! Reason: %s\n", err.Error())
				return nil
			}
			return err
		}
	}
	return nil
}

// charge aliens started charging
func (i *invasion) charge() error {
	fmt.Println("Charging...")
	// Shuffling the aliens for random selection
	i.randGen.Shuffle(len(i.aliens), func(a, b int) {
		i.aliens[a], i.aliens[b] = i.aliens[b], i.aliens[a]
	})
	noAttack := true
	for _, alien := range i.aliens {
		if err := i.beginAttack(alien); err != nil {
			// Attack didn't succeed
			//  1. Alien dead or trapped
			//  2. World is destroyed
			if _, ok := err.(*models.InvasionError); ok {
				fmt.Printf("Alien: %d attack failed with err: %s\n", alien.ID, err.Error())
				continue
			}
			// other errors
			return err
		}
		noAttack = false
		fmt.Printf("Remaning cities after attack by Alien: %d\n", alien.ID)
		fmt.Println("====================================")
		fmt.Print(i.World.GetUndestroyedCities())
		fmt.Println("====================================")
	}
	if noAttack {
		return &models.InvasionError{
			Reason: models.ErrNoAttack,
		}
	}
	return nil
}

// beginAttack alien begins attack
func (i *invasion) beginAttack(alien *models.Alien) error {
	fmt.Printf("\n\nAlien: %d is beginning attack\n", alien.ID)
	// Make sure the alien is not already dead or trapped
	if !alien.IsOkay() {
		return &models.InvasionError{
			Reason: models.ErrAlienDeadOrTrapped,
		}
	}
	// Choose a target city
	target, err := i.chooseTarget(alien)
	if err != nil {
		return err
	}
	if alien.CurrCity != nil {
		// update presence as this alien is making a move from current city
		i.alienPresence[alien.CurrCity.Name] = nil
	}
	fmt.Printf("Alien: %d currCity: %s, targetCity: %s\n", alien.ID, alien.CurrCity, target)
	// Invading target city
	alien.CurrCity = target
	// Check if the target city already has an alien presence
	if i.alienPresence[target.Name] != nil {
		tAlien := i.alienPresence[target.Name]
		// Destroy target city and mark both aliens dead
		fmt.Printf("%s has been destroyed by alien %d and alien %d\n", target.Name, tAlien.ID, alien.ID)
		target.Destroyed = true
		tAlien.Dead = true
		alien.Dead = true
		// Remove city from presence map
		delete(i.alienPresence, target.Name)
	} else {
		// If no alien presence in the target city, update presence
		i.alienPresence[target.Name] = alien
	}
	return nil
}

// chooseTarget  chooses a target city
func (i *invasion) chooseTarget(alien *models.Alien) (*models.City, error) {
	// Choose any random undestroyed city for initial attack
	if alien.CurrCity == nil {
		return i.chooseRandomCity()
	}
	// Choose any random connected city if its not initial attack.
	return i.chooseRandomConnectedCity(alien.CurrCity)
}

// chooseRandomCity chooses a random city for initial attack
func (i *invasion) chooseRandomCity() (*models.City, error) {
	// Get undestroyed cities
	var cities []string
	for k := range i.World.Cities {
		if !i.World.Cities[k].Destroyed {
			cities = append(cities, k)
		}
	}
	// All cities destroyed
	if len(cities) == 0 {
		return nil, &models.InvasionError{
			Reason: models.ErrWorldDestroyed,
		}
	}
	// Choose random index and return the city
	randIdx := i.randGen.Intn(len(cities))
	return i.World.Cities[cities[randIdx]], nil
}

// chooseRandomConnectedCity chooses any random connected city for subsequent attacks.
func (i *invasion) chooseRandomConnectedCity(currCity *models.City) (*models.City, error) {
	// Create a slice of length len(currCity.Roads) to shuffle the indices for random selection
	indices := make([]int, len(currCity.Roads))
	for i := 0; i < len(currCity.Roads); i++ {
		indices[i] = i
	}
	// Shuffle the temp slice
	i.randGen.Shuffle(len(currCity.Roads), func(a, b int) {
		indices[a], indices[b] = indices[b], indices[a]
	})
	// Any undestroyed connected city
	for _, val := range indices {
		roadName := currCity.Roads[val].Name
		toCity := currCity.RoadToCity[roadName]
		if c := i.World.Cities[toCity.Name]; !c.Destroyed {
			return c, nil
		}
	}
	// No valid cities found
	return nil, &models.InvasionError{
		Reason: models.ErrAlienDeadOrTrapped,
	}
}

// WriteRemains writes the remains of the world on the console
func (i *invasion) WriteRemains() {
	fmt.Println("Writing the remaning cities")
	fmt.Println(i.World.GetUndestroyedCities())
}

// oppositeDir gives the opposite of a given direction
func oppositeDir(dir string) string {
	switch dir {
	case "north":
		return "south"
	case "east":
		return "west"
	case "south":
		return "north"
	case "west":
		return "east"
	}
	return ""
}
