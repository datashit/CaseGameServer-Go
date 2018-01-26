package match

import (
	"encoding/json"

	"github.com/datashit/CaseGameServer-Go/gameserver/game"
)

// message : TCP mesaj komut yığını
type message struct {
	PlayerID uint64 // Oyuncu ID
	Command  string // Komut
	Data     string // Komut datası
}

// incomeJobs : Gelen komut işleme yığını
type incomeJobs struct {
	inClient *game.Client  // Client
	inMsg    message       // Gelen mesaj
	encoder  *json.Encoder // Client json encoder
}
