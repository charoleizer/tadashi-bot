package bot

import (
	"github.com/charoleizer/tadashi-bot/bot/actions"

	"github.com/bwmarrin/discordgo"
)

const (
	ping = "ping"
)

func Router(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Content == ping {
		actions.DoPing(session, message)
	}

}
