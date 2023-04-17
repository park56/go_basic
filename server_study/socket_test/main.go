package main

import (
	"socket_test/mysocket"

	"github.com/gin-gonic/gin"
)

func main() {

	//mux := pat.New()
	//mux.Get("/ws", mysocket.Wshandler) // /ws에 get으로 wshandler에 따라 커넥션을 맷게 됨
	//n := negroni.Classic()
	//n.UseHandler(mux)
	//http.ListenAndServe(":8080", n)

	g := gin.Default()

	g.GET("/ws", mysocket.Wshandler)
	g.POST("/ws", mysocket.Wshandler)

	g.Run()
}
