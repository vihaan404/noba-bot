package discord

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var CommandHandler = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"basic": simpleHandler,
}

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

func (botSession *Bot) DeleteCommands(registeredCommands []*discordgo.ApplicationCommand) {
	if true {
		slog.Info("Removing commands...")
		// // We need to fetch the commands, since deleting requires the command ID.
		// // We are doing this from the returned commands on line 375, because using
		// // this will delete all the commands, which might not be desirable, so we
		// // are deleting only the commands that we added.
		// registeredCommands, err := s.ApplicationCommands(s.State.User.ID, *GuildID)
		// if err != nil {
		// 	log.Fatalf("Could not fetch registered commands: %v", err)
		// }

		for _, v := range registeredCommands {
			err := botSession.ApplicationCommandDelete(botSession.State.User.ID, os.Getenv("GUILD_ID"), v.ID)
			if err != nil {
				slog.Error("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}
}

func (botSession *Bot) StartDiscordBot() ([]*discordgo.ApplicationCommand, error) {
	err := botSession.Open()
	if err != nil {
		slog.Error("error in connecting with discord", err)
		return nil, err
	}

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	slog.Info("Creating application commands . . .. . . ")
	for i, v := range commands {
		cmd, err := botSession.ApplicationCommandCreate(botSession.State.User.ID, os.Getenv("GUILD_ID"), v)
		if err != nil {
			slog.Error("error creating application command ", err)
			return nil, err
		}
		registeredCommands[i] = cmd
	}
	slog.Info("Discord bot is listening ")
	return registeredCommands, nil
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
