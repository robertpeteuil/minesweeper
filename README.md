# Minesweeper

Go(lang) implementation of Minesweeper game API

[Minesweeper History](https://en.wikipedia.org/wiki/Minesweeper_(video_game))

## Build Binary

``` sh
go build -o build/minesweeper .
```

## Run Binary

``` sh
./build/minesweeper
```

## Build Container

``` sh
./build-docker.sh
```

## Create a New Game

``` sh
curl -i -X POST '127.0.0.1:3000/game' -d '{"name": "teste", "rows": 10, "cols": 8, "mines": 20}'
```

## Start the game

``` sh
curl -i -X POST '127.0.0.1:3000/game/teste/start'
```

## Play

``` sh
curl -i -X POST '127.0.0.1:3000/game/teste/click' -d '{"row": 1,"col":1}'
```

## Run tests

``` sh
go clean  $(go list ./... | grep -v /vendor/)
go test  $(go list ./... | grep -v /vendor/) -v
```
