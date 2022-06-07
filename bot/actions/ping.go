package actions

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

const (
	red    = "🔴"
	green  = "🟢"
	blue   = "🔵"
	yellow = "🟡"
	purple = "🟣"
)

var RoleByEmoji = make(map[string]string)

type Role struct {
	ID      string
	Name    string
	GuildID string
}

type Guild struct {
	Roles map[string]Role
}

var Guilds = make(map[string]Guild)

func mapRolesByEmoji() {
	RoleByEmoji[red] = "red"
	RoleByEmoji[green] = "green"
	RoleByEmoji[blue] = "blue"
	RoleByEmoji[yellow] = "yellow"
	RoleByEmoji[purple] = "purple"
}

func getRoles(session *discordgo.Session) {
	var Roles = make(map[string]Role)

	for _, guild := range session.State.Guilds {

		Guild := Guild{}
		for _, role := range guild.Roles {
			Role := Role{}
			Role.ID = role.ID
			Role.Name = role.Name
			Role.GuildID = guild.ID

			Guild.Roles = Roles
			Guild.Roles[role.Name] = Role
		}

		Guilds[guild.ID] = Guild
	}
}

func DoPing(session *discordgo.Session, requestMessage *discordgo.MessageCreate) {

	go func() {
		responseMessage, err := session.ChannelMessageSend(requestMessage.ChannelID, "pong - Escolhe quais cores você deseja representar")

		if err != nil {
			log.Println(err)
		}

		// TODO - fazer um loop percorrendo RoleByEmoji
		session.MessageReactionAdd(responseMessage.ChannelID, responseMessage.ID, red)
		session.MessageReactionAdd(responseMessage.ChannelID, responseMessage.ID, green)
		session.MessageReactionAdd(responseMessage.ChannelID, responseMessage.ID, blue)
		session.MessageReactionAdd(responseMessage.ChannelID, responseMessage.ID, yellow)
		session.MessageReactionAdd(responseMessage.ChannelID, responseMessage.ID, purple)
	}()

	go mapRolesByEmoji()
	go getRoles(session)

	go session.AddHandler(reactionHandler)
	go session.AddHandler(unreactionHandler)

}

func reactionHandler(session *discordgo.Session, reaction *discordgo.MessageReactionAdd) {
	if reaction.UserID == session.State.User.ID {
		return
	}

	guildID := reaction.MessageReaction.GuildID
	userID := reaction.MessageReaction.UserID
	roleID := Guilds[guildID].Roles[RoleByEmoji[reaction.Emoji.Name]].ID

	if roleID != "" {

		go func() {
			err := session.GuildMemberRoleAdd(guildID, userID, roleID)
			if err != nil {
				log.Println(err)
			}
		}()

		return
	}

	session.MessageReactionRemove(reaction.ChannelID, reaction.MessageID, reaction.Emoji.Name, reaction.Member.User.ID)

}

func unreactionHandler(session *discordgo.Session, reaction *discordgo.MessageReactionRemove) {
	if reaction.UserID == session.State.User.ID {
		return
	}

	guildID := reaction.MessageReaction.GuildID
	userID := reaction.MessageReaction.UserID
	roleID := Guilds[guildID].Roles[RoleByEmoji[reaction.Emoji.Name]].ID

	if roleID != "" {
		go func() {
			err := session.GuildMemberRoleRemove(guildID, userID, roleID)
			if err != nil {
				log.Println(err)
			}
		}()

		return
	}

}
