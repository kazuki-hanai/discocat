package main

import (
	"os"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"

	"github.com/wan-nyan-wan/discocat/config"
)

var (
	configuration config.Configuration
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file, ", err)
		os.Exit(1)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Println("Unable to decode into struct, ", err)
		os.Exit(1)
	}
}

func main() {
	discord, err := discordgo.New("Bot " + configuration.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	msg, err := discord.ChannelMessageSend(configuration.ChannelId, "hello")
	if err != nil {
		fmt.Println("cannot sent message, ", err)
		return
	}
	fmt.Println("send message: ", msg)
}
