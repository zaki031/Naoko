package botFuncs


import(
	"github.com/bwmarrin/discordgo"
)



func ChangeStatus(s *discordgo.Session, event *discordgo.Event) {
	s.UpdateWatchStatus(5, "hmmmmmm")
}
