package main

import (
	"flag"

	"github.com/datashit/CaseGameServer-Go/gameserver"
)

var addrServer = flag.String("addrServer", "localhost:3333", "Game server address")

func main() {
	server := gameserver.NewInstance(*addrServer)
	defer server.Shutdown()
	server.Start()
}
