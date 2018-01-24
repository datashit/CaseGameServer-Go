package gameserver

import (
	"net"
	"testing"
)

func TestSocketCheckServerStartandShutdown(t *testing.T) {
	addr := "127.0.0.1:3333"
	server := NewInstance(addr)
	go server.Start() // Server baslatiliyor

	for !server.isRun() {
		// Serverın online olması bekleniyor.
	}

	conn, err := net.Dial("tcp", addr) // Server listener check ediliyor.
	if err != nil {
		t.Error("Server not starting!")
	}
	conn.Close()

	server.Shutdown()

	for server.isRun() {
		// Serverın kapatılması bekleniyor.
	}

	conn2, err := net.Dial("tcp", addr) // Serverın kapandığı check ediliyor.
	if err == nil {
		defer conn2.Close()
		t.Error("Server not shutdown!")
	}

}
