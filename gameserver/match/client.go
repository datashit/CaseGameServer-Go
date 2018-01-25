package match

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"sync/atomic"
)

// Client : kullanıcı yığını
type Client struct {
	playerID uint64   // Oyuncu numarası
	conn     net.Conn // Oyuncu tcp bağlantısı
}

var playerIDCounter uint64 //Benzersiz player id sayacı

func createClient(c net.Conn) *Client {

	client := &Client{
		playerID: atomic.AddUint64(&playerIDCounter, 1), // Benzersiz player id kullanılmalı
		conn:     c,
	}

	return client
}

func (client *Client) handle() {

	defer client.conn.Close()

	//	decoder := json.NewDecoder(client.conn)
	encoder := json.NewEncoder(client.conn)
	var wel = message{
		PlayerID: client.playerID,
		Command:  "WELCOME",
		Data:     "",
	}
	reader := bufio.NewReader(client.conn)

	encoder.Encode(wel)

	for {
		lineData, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			fmt.Println("Server disconnected: " + client.conn.RemoteAddr().String())
			return
		}

		if err != nil {
			fmt.Println("Reader error: ", err)
			return
		}

		if isPrefix {

			continue
		}

		var msg message
		err = json.Unmarshal(lineData, &msg)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if msg.PlayerID != client.playerID {
			fmt.Println("Player id does not match")
			return
		}

		// Message handle'a gönderilmeli
		fmt.Println(msg)

	}

}
