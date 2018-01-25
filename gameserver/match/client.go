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
	encoder  *json.Encoder
	reader   *bufio.Reader
}

var playerIDCounter uint64 //Benzersiz player id sayacı

func createClient(c net.Conn) *Client {

	client := &Client{
		playerID: atomic.AddUint64(&playerIDCounter, 1), // Benzersiz player id kullanılmalı
		conn:     c,
		encoder:  json.NewEncoder(c),
		reader:   bufio.NewReader(c),
	}

	return client
}

func (client *Client) welcome() {
	var wel = message{
		PlayerID: client.playerID,
		Command:  "WELCOME",
		Data:     "",
	}

	client.encoder.Encode(wel)
}

func (client *Client) handle() {

	defer fmt.Printf("Simultaneous : %v \r\n", atomic.AddUint32(&simultaneous, ^uint32(0)))
	defer client.conn.Close()

	client.welcome()

	for {
		lineData, isPrefix, err := client.reader.ReadLine()
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

		fmt.Println(msg)
		// Mesaj handle'a gönderilmeli
		ijob := incomeJobs{inClient: *client, inMsg: msg}
		incomeMessageJobs <- ijob
	}

}
