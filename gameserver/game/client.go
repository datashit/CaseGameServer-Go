package game

import (
	"net"
	"sync/atomic"
)

var playerIDCounter uint64 //Benzersiz player ID sayacı

// Client : kullanıcı yığını
type Client struct {
	PlayerID uint64       // Oyuncu numarası
	Conn     *net.Conn    // Oyuncu tcp bağlantısı
	UDPCon   *net.UDPConn // Oyuncu udp socketi
	UDPAddr  *net.UDPAddr // Oyuncu udp adresi
}

// CreateClient : Bir kullanıcı olusturup geri doner.
func CreateClient(c *net.Conn) *Client {

	client := &Client{
		PlayerID: atomic.AddUint64(&playerIDCounter, 1), // Benzersiz player id ataması yapılıyor.
		Conn:     c,
	}

	return client
}
