# Alien Invasion challenge
Mad aliens are about to invade the earth and you are tasked with simulating the
invasion.
You are given a map containing the names of cities in the non-existent world of
X. The map is in a file, with one city per line. The city name is first,
followed by 1-4 directions (north, south, east, or west). Each one represents a
road to another city that lies in that direction.
For example:
Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee
The city and each of the pairs are separated by a single space, and the
directions are separated from their respective cities with an equals (=) sign.
You should create N aliens, where N is specified as a command-line argument.
These aliens start out at random places on the map, and wander around randomly,
following links. Each iteration, the aliens can travel in any of the directions
leading out of a city. In our example above, an alien that starts at Foo can go
north to Bar, west to Baz, or south to Qu-ux.
When two aliens end up in the same place, they fight, and in the process kill
each other and destroy the city. When a city is destroyed, it is removed from
the map, and so are any roads that lead into or out of it.
In our example above, if Bar were destroyed the map would now be something
like:
Foo west=Baz south=Qu-ux
Once a city is destroyed, aliens can no longer travel to or through it. This
may lead to aliens getting "trapped".
You should create a program that reads in the world map, creates N aliens, and
unleashes them. The program should run until all the aliens have been
destroyed, or each alien has moved at least 10,000 times. When two aliens
fight, print out a message like:

Bar has been destroyed by alien 10 and alien 34!
(If you want to give them names, you may, but it is not required.) Once the
program has finished, it should print out whatever is left of the world in the
same format as the input file.
Feel free to make assumptions (for example, that the city names will never
contain numeric characters), but please add comments or assertions describing
the assumptions you are making.

## Build & Run
```
$cd alien-inv
$make
$cd bin
$ ./alien-invasion 
SW version: 1.0.0 buildtime: 2023-03-08_20:27:43
Usage of ./alien-invasion:
  -h, --help                print help and exit
  -i, --maxAlienMoves int   max alien moves (default 10000)
  -n, --numaliens int       number of aliens (default 10)
  -m, --worldmap string     world map file (Mandatory)
```
Example Usage:
```
$ ./alien-invasion -m default_map.txt -n 20 -i 4000
```

## Assumptions:
    1. One entry per city.
    2. Space is used as delimiter to parse cities and links.So assuming no space in the names of cities.
    3. No leading or trailing spaces around "=" sign.
    4. City and Direction names are case sensitive.
    5. Consider the below example,
        Foo north=Bar west=Baz south=Qu-ux
        Bar south=Foo west=Bee
       In the above example "Baz" and "Bee" are present in the links.But there is no city entry for "Baz" an "Bee". In such cases,
       We will explicitly add city entries in the "World". For instance, "Baz" will be added with a link having direction "west" with city "Foo"
    6. Aliens are numbered from 0 to n

## Implementation Details
  1. At every iteration, aliens are picked randomly for attack.
  2. The city for initial attack and subsequent attacks are chosen randomly.
  3. Remaining cities are printed after every attack.
  4. Simulations are stopped if no aliens made a move(dead, trapped or world is destryed) in an 
     iteration or if all the aliens have moved "maxAlienMoves" times  
  5. Unit tests are written and they are run while issuing "make" command.
  6. To run the tests separately use ``` go test ./... -v ``` command