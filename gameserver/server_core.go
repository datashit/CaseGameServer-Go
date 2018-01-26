package gameserver

import (
	"fmt"
	"net"
)

//ServerCore : Server yığını
type ServerCore struct {
	ListenerSocket net.Listener   // Dinlenen soket
	shutdownSignal chan bool      // Kapatma sinyali
	runing         bool           // Server calıştı bilgisi
	settings       ServerSettings // Server ayarları
}

//NewInstance : bir server nesnesi oluşturup geri döner.
func NewInstance(address string) *ServerCore {
	server := &ServerCore{
		shutdownSignal: make(chan bool, 1),
	}
	server.settings.HostAddr = address
	server.settings.Load() // Server ayarları yükleniyor

	return server
}

// IsRun : serverın çalışıp çalışmadığı bilgisini döner.
func (server *ServerCore) IsRun() bool {
	return server.runing
}

//Start : metodu serverı başlatır
func (server *ServerCore) Start() {
	fmt.Println("Server Starting...")
	server.listen() // Server ana döngüsü başlatılıyor
	server.runing = false
}

// Shutdown : metodu serverı kapatır
func (server *ServerCore) Shutdown() {
	server.ListenerSocket.Close()
	server.shutdownSignal <- true
	fmt.Println("Server Shutdown!")

}

func (server *ServerCore) listen() {
	l, err := net.Listen("tcp", server.settings.HostAddr)
	if err != nil {
		fmt.Println("Error: ")
		fmt.Println(err)
		return
	}
	defer l.Close()
	server.ListenerSocket = l

	fmt.Printf("Socket is listening: %s \r\n", server.ListenerSocket.Addr())

	server.runing = true
	for {
		select {
		case <-server.shutdownSignal:
			return
		default:
			conn, err := server.ListenerSocket.Accept()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Incoming connection request <--")

			// Baglantı handle' a gönderiliyor.
			go server.settings.handsaheker.HandShake(&conn)
		}
	}

}
