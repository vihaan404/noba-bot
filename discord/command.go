package discord

import "github.com/bwmarrin/discordgo"

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "basic",
		Description: "its a basic command what else do you need ",
	},
}
