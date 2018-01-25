package match

import (
	"fmt"

	"github.com/datashit/CaseGameServer-Go/gameserver/game"
)

var incomeMessageJobs = make(chan incomeJobs, 1000)

// CreateMatchHandleWorker : fonksiyonu asenkron matchHandle worker olusturur.
// workerSize : worker sayısı
func CreateMatchHandleWorker(workerSize int) {
	for w := 1; w <= workerSize; w++ {
		go matchHandle(incomeMessageJobs)
	}
}

func matchHandle(job <-chan incomeJobs) {
	for j := range job {
		switch j.inMsg.Command {
		case "FIND_GAME":
			fmt.Println("Finding game")
			j.inMsg.Command = "SEARCH_GAME"
			j.inClient.encoder.Encode(j.inMsg)
			match(j.inClient)
		default:
		}
	}
}

var partner = make(chan Client)

func match(client Client) {
	fmt.Printf("Player %v : Waiting for a player... \r\n", client.playerID)
	select {
	case partner <- client:
	case p := <-partner:
		g := game.CreateGame(p.playerID, client.playerID)
		fmt.Printf("Game Create : %v", g.Id)
		go g.StartListen()
	}
}
