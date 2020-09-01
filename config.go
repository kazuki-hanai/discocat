package main

// DiscordConfig is a configuration for Discord
type DiscordConfig struct {
	Token string
	ChannelID string
}

// CommandConfig is a configuration for Command
type CommandConfig struct {
	comment string
	filepath string
	isPipe bool
}
