package gameserver

import (
	"github.com/datashit/CaseGameServer-Go/gameserver/game"
	"github.com/datashit/CaseGameServer-Go/gameserver/match"
	"github.com/datashit/CaseGameServer-Go/gameserver/security"
)

// ServerSettings : yığını içinde server ayarları bulunmakta.
type ServerSettings struct {
	HostAddr    string                 // Host soket adresi
	handsaheker match.IClientHandshake // Client doğrulama yöntemi
}

// Load : metodu server ayarlarını yükler
func (s *ServerSettings) Load() {

	s.handsaheker = &match.NoClientHandshake{CryptoType: &security.Nocrypto{}}

	match.CreateMatchHandleWorker(5)

	game.InitGames()
	game.CreateGameRunWorker(10)
	game.CreateGameIncomeMessageWorker(10)
	go game.StartGameUDPListen(s.HostAddr)
}
