package discord

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	*discordgo.Session
}

func CreateNewBot() *Bot {
	bot, err := discordgo.New("Bot " + os.Getenv("DISCORD_API_SECRET"))
	if err != nil {
		slog.Error("error in creating a new bot ->", err)
	}
	return &Bot{
		bot,
	}
}

func (botSession *Bot) StartDiscordBot() error {
	err := botSession.Open()
	if err != nil {
		slog.Error("error in connecting with discord", err)
		return err
	}
	slog.Info("Creating application commands . . .. . . ")
	for _, v := range commands {
		_, err := botSession.ApplicationCommandCreate(botSession.State.User.ID, os.Getenv("GUILD_ID"), v)
		if err != nil {
			slog.Error("error creating application command ", err)
			return err
		}
	}
	slog.Info("Discord bot is listening ")
	return nil
}

func (botSession *Bot) CloseDiscordBot() {
	err := botSession.Close()
	if err != nil {
		slog.Error("error closing the bot", err)
	}
}

func TerminateOnSignal() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
