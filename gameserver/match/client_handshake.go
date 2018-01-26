package match

import (
	"fmt"
	"net"

	"github.com/datashit/CaseGameServer-Go/gameserver/game"
	"github.com/datashit/CaseGameServer-Go/gameserver/security"
)

// IClientHandshake : Client Handshake Arayüzü
type IClientHandshake interface {
	HandShake(*net.Conn)
}

// NoClientHandshake : direk kabul eden hand shake yığını
type NoClientHandshake struct {
	CryptoType security.Imessagecrypto
}

// HandShake :  Gelen isteği değerlendirir
func (shake *NoClientHandshake) HandShake(conn *net.Conn) {
	fmt.Println("Handshake Complate!")

	c := game.CreateClient(conn) // Kullanıcı oluşturuluyor.
	go clientReadHandle(c)       // Kullanıcı match handle'ı başlatılıyor.
}
