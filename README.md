# Minesweeper

Go(lang) implementation of Minesweeper game API

More about the game: https://en.wikipedia.org/wiki/Minesweeper_(video_game)

## Build

``` sh
go build -o build/minesweeper ./cmd
```

## Run

``` sh
./build/minesweeper
```

## Build and Run with Docker

``` sh
sh run-docker.sh
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

