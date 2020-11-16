package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"

	_ "github.com/joho/godotenv/autoload"
)

var sendMessageURL string

func init() {
	token := os.Getenv("TOKEN")

	if len(token) > 0 {
		sendMessageURL = fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	}
}

func main() {
	server := gin.Default()

	server.Any("*action", informerHandler)

	server.Run(":80")
}

func informerHandler(c *gin.Context) {
	buf := new(bytes.Buffer)
	c.Request.Write(buf)

	sendMessage(buf.String())

	c.String(http.StatusOK, buf.String())
}

func sendMessage(text string) {
	chatID := os.Getenv("CHAT_ID")

	if sendMessageURL == "" || len(chatID) < 1 {
		return
	}

	gorequest.New().Get(sendMessageURL).
		Param("chat_id", chatID).
		Param("parse_mode", "MarkdownV2").
		Param("text", "```\n"+text+"\n```").
		End()
}
