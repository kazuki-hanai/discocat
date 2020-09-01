package main

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
