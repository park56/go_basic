package mysocket

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

/*func Wshandler(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		m := &Message{}
		err := conn.ReadJSON(m)
		if err != nil {
			log.Println(err)
			return
		}

		err = conn.WriteJSON(m)
		if err != nil {
			log.Println(err)
			return
		}
	}
	// ... Use conn to send and receive messages.
}*/

func Wshandler(c *gin.Context) {

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()
	for {
		m := &Message{}
		err := ws.ReadJSON(m)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = ws.WriteJSON(m)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
