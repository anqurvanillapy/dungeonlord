package main

import (
	"github.com/anqurvanillapy/dungeonlord/gameserver"
	"github.com/anqurvanillapy/dungeonlord/httpserver"
)

func main() {
	gameserver.GameServer()
	httpserver.HttpServer()
}
