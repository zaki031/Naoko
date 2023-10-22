package botFuncs

import (
	"fmt"
	"os"
	"scrap/ApiCalls"

	"github.com/bwmarrin/discordgo"
)





func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// check if the message is from a bot or not a user.
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	// check if the message content is !rand only commend.
	if m.Content == "!rand" {
		fmt.Println(m.Content[5:])
		ApiCalls.GetRandom()

		file, err := os.Open("image.png")
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error opening file.")
			return
		}
		defer file.Close()

		// Send the image through discord
		s.ChannelFileSend(m.ChannelID, "image.png", file)
	}

	// Check if the message content is !rand + lang command.

	if len(m.Content) > 6 {

		if m.Content[0:5] == "!rand" {

			ApiCalls.WaifuByLang(m.Content[6:])
			file, err := os.Open("image.png")
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Error opening file.")
				return
			}
			defer file.Close()

			// Send the image through discord
			s.ChannelFileSend(m.ChannelID, "image.png", file)
		}
	}

}

