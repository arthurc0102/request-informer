package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"

	_ "github.com/joho/godotenv/autoload"
)

var (
	sendMessageURL string
	chatID         string
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.WarnLevel)
	logrus.SetReportCaller(true)

	token := os.Getenv("TOKEN")
	if len(token) > 0 {
		sendMessageURL = fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	} else {
		logrus.Warning("No token set.")
	}

	chatID = os.Getenv("CHAT_ID")
	if len(chatID) < 1 {
		logrus.Warning("No chat id set.")
		return
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
	gorequest.New().Get(sendMessageURL).
		Param("chat_id", chatID).
		Param("parse_mode", "MarkdownV2").
		Param("text", "```\n"+text+"\n```").
		End()
}
