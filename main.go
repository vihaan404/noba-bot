package main

import (
	"log/slog"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/vihaan404/noba-bot/discord"
)

func main() {
	bootStrapLocalDev()
	bot := discord.CreateNewBot()

	bot.AddHandler(simpleHandler)
	bot.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := discord.CommandHandler[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
	registeredCommands, err := bot.StartDiscordBot()
	if err != nil {
		slog.Error("discord", err)
	}

	defer bot.CloseDiscordBot()
	discord.TerminateOnSignal()
	bot.DeleteCommands(registeredCommands)

	slog.Info("Gracefully shutting down.")
}

func bootStrapLocalDev() {
	if os.Getenv("ENV") == "local" {
		err := godotenv.Load()
		if err != nil {
			slog.Error("Error loading .env file", err)
			return
		}
	}
	return
}

func simpleHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "hey there this is your first discord slah command",
		},
	})
}
