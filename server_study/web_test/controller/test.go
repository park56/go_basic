package controller

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {

	io.WriteString(c.Writer, "Tesst string")

	fmt.Println("URl := " + c.Request.URL.Path)
}
