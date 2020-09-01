package main

import (
	"fmt"
)

// DiscordConfig is a configuration for Discord
type DiscordConfig map[string] struct {
	BotToken string
	ChannelIDs map[string] string
}

// CommandConfig is a configuration for Command
type CommandConfig struct {
	comment string
	filepath string
	isPipe bool
}

func (discoConfig DiscordConfig) printConfig() {
	for k, v := range discoConfig {
		fmt.Println(k, ":")
		fmt.Println("\tBotToken:", v.BotToken)
		fmt.Println("\tChannelIDs:")
		for k2, v2 := range v.ChannelIDs {
			fmt.Println("\t\t", k2, ":", v2)
		}
	}
}
