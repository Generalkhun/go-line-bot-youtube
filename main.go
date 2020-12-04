package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

func getEnv() {
	err := godotenv.Load("./secret/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	getEnv()
	// Create router

	router := gin.Default()

	//Create linebot connection
	linebotConnection := connectLineBot()

	// Create bot response  services
	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"Message": "index"})
	// })

	// router.POST("/callback", func(c *gin.Context) {
	// 	lineBotReq(linebotConnection, c.Request)
	// })
	router.POST("/", func(c *gin.Context) {
		lineBotReq(linebotConnection, c.Request)
	})

	// Start service
	port := os.Getenv("PORT")
	router.Run(":" + port)
}

func connectLineBot() *linebot.Client {
	getEnv()
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
	return bot
}

// parse linebot request
func lineBotReq(bot *linebot.Client, req *http.Request) {
	events, err := bot.ParseRequest(req)
	if err != nil {
		// Do something when something bad happened.
		panic(err)
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			fmt.Println("Hello server from line bot!!!")

			// Send message back to linebot
			//userID := event.Source.UserID
			replyToken := event.ReplyToken

			var messages []linebot.SendingMessage

			//Get Youtube URL from list
			url := service.youtubeapi.msgYTCLipURL()

			// append some message to messages
			messages = append(messages, linebot.NewTextMessage("Hi linebot -> client from server"))
			_, err := bot.ReplyMessage(replyToken, messages...).Do()
			if err != nil {
				// Do something when some bad happened
				panic(err)
			}
		}
	}
}
