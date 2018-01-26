# CI
[![Build Status](https://travis-ci.org/datashit/CaseGameServer-Go.svg?branch=master)](https://travis-ci.org/datashit/CaseGameServer-Go)

# CaseGameServer-Go
TCP ve UDP üzerinden haberleşen 2 kişilik game loby kuran bir game server case yazılımı.

# Kullanımı
main.go altında tanımlanan Flag üzerindeki Portu hem TCP hemde UDP olarak dinler.<br />
var addrServer = flag.String("addrServer", "localhost:3333", "Game server address")

Server online olduğunda aşağıdakine benzer bir console çıktısı verir:<br />
Socket is listening: 127.0.0.1:3333

Client TCP bağlantısı yaptığında server player ID ile birlikte welcome  aşağıdakine benzer mesaj döner.<br />
{"PlayerID":1,"Command":"WELCOME","Data":""}

Client Server' a aşağıdaki mesajı yollarsa oyun arama kuyruğuna girer:<br />
{"PlayerID":1,"Command":"FIND_GAME","Data":""}<br />
Oyun arama kuyruk cevabı:<br />
{"PlayerID":1,"Command":"SEARCH_GAME","Data":""}

Oyuncu bulunduğunda serverdan aşağıdaki örnekte olduğu gibi game ID ile birlikte mesaj gelir:<br />
{"PlayerID":1,"Command":"GAME_CREATED","Data":"{"GameID":1}"}


Gelen game ID ile oyuna aşağıdaki mesaj ile bağlanılır:<br />
{"GameID":1,"PlayerID":1,"Command":"CONNECT_GAME","Data":""}

Her iki oyuncuda oyuna bağlandığında oyun başlar.<br />
