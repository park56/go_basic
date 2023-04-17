package mysocket

import (
	"github.com/gorilla/websocket"
)

/*var upgrader = websocket.Upgrader{ // 버퍼를 얼마나 유지할 것 인가
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}*/

type Message struct {
	Type string      `json:"type"`
	User string      `json:"user"`
	Data interface{} `json:"data"`
}

var upgrader = websocket.Upgrader{
	//check origin will check the cross region source (note : please not using in production)
	/*CheckOrigin: func(r *http.Request) bool {
		//Here we just allow the chrome extension client accessable (you should check this verify accourding your client source)
		return origin == "chrome-extension://cbcbkhdmedgianpaifchdaddpnmgnknn"
	},*/

	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
