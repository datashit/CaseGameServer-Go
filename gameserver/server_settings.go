package gameserver

import (
	"github.com/datashit/CaseGameServer-Go/gameserver/match"
	"github.com/datashit/CaseGameServer-Go/gameserver/security"
)

// ServerSettings : yığını içinde server ayarları bulunmakta.
type ServerSettings struct {
	handsaheker match.IClientHandshake // Client doğrulama yöntemi
}

// Load : metodu server ayarlarını yükler
func (s *ServerSettings) Load() {

	s.handsaheker = &match.NoClientHandshake{CryptoType: &security.Nocrypto{}}
}
