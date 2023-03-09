# Alien Invasion

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