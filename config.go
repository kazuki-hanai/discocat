package main

import (
	"fmt"
)

// DiscordConfig is a configuration for Discord
type DiscordConfig map[string]struct {
	BotToken  string
	ChannelID string
}

// CommandConfig is a configuration for Command
type CommandConfig struct {
	comment  string
	filepath string
	isPipe   bool
}

func (discoConfig DiscordConfig) printConfig() {
	for k, v := range discoConfig {
		fmt.Println(k, ":")
		fmt.Println("\tBotToken:", v.BotToken)
		fmt.Println("\tChannelID", v.ChannelID)
	}
}
