package main

import (
	"log"
	"os"
	"scrap/botFuncs"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	// loading the .env bot token
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Access the loaded environment variables
	botToken := os.Getenv("BOT_TOKEN")
	// creating the bot
	
	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		panic("Error creating Discord session: " + err.Error())
	}

	//Handlers.
	dg.AddHandler(botFuncs.MessageCreate)
	dg.AddHandler(botFuncs.ChangeStatus)

	err = dg.Open()
	if err != nil {
		panic("Error opening connection: " + err.Error())
	}

	//intents
	dg.Identify.Intents = discordgo.IntentsGuildMessages
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildMembers | discordgo.IntentsGuildPresences

	// Keep the bot running.
	select {}

}
