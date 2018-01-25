package match

import (
	"fmt"
	"net"
	"sync/atomic"

	"github.com/datashit/CaseGameServer-Go/gameserver/security"
)

// IClientHandshake : Client Handshake Arayüzü
type IClientHandshake interface {
	HandShake(net.Conn)
}

// NoClientHandshake : direk kabul eden hand shake yığını
type NoClientHandshake struct {
	CryptoType security.Imessagecrypto
}

var simultaneous uint64 // Bağlı tcp soket sayısı

// HandShake :  Gelen isteği değerlendirir
func (shake *NoClientHandshake) HandShake(conn net.Conn) {
	fmt.Println("Handshake Complate!")

	fmt.Printf("Simultaneous : %v \r\n", atomic.AddUint64(&simultaneous, 1))

	go createClient(conn).handle()
}
