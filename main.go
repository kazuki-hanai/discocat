package main

import (
	"os"
	"fmt"
	"errors"
	"io/ioutil"
	"time"
	"bytes"
	"image"
	_ "image/gif"
	_ "image/png"
	_ "image/jpeg"

	"github.com/bwmarrin/discordgo"
	"github.com/urfave/cli/v2"
	"github.com/spf13/viper"
)

var (
	commandName = "discocat"
	build   = ""
	version = "v.1.0"
	defaultConfigPaths = os.Getenv("HOME") + "/.config/discocat/"
)

func handleUsageError(c *cli.Context, err error, _ bool) error {
	printErr(fmt.Errorf("%s %s", "Incorrect Usage.", err.Error()))
	cli.ShowAppHelp(c)
	return cli.Exit("", 1)
}

func printFullVersion(c *cli.Context) {
	fmt.Fprintf(c.App.Writer, "%v version %v, build %v\n", c.App.Name, c.App.Version, build)
}

// Type of stdin buffer
const (
	Text = iota
	Png
	Gif
	Jpeg
)

func detectMessageType(raw []byte) (int, error) {
	_, format, err := image.Decode(bytes.NewReader(raw))
	if err != nil {
		return Text, nil
	}

	if format == "png" {
		return Png, nil
	} else if format == "gif" {
		return Gif, nil
	} else if format == "jpeg" {
		return Jpeg, nil
	} else {
		return -1, errors.New("Could not detect filetype")
	}
}

func post(raw []byte, token string, channel string) error {
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		return err
	}

	mtype, err := detectMessageType(raw)
	if err != nil {
		return err
	}

	if mtype == Text {
		_, err := discord.ChannelMessageSend(channel, string(raw))
		if err != nil {
			return err
		}
	} else {
		times := time.Now().Format("20060102150405")
		var filename = ""
		if mtype == Png {
			filename = fmt.Sprintf("%s.%s", times, ".png")
		} else if mtype == Gif {
			filename = fmt.Sprintf("%s.%s", times, ".gif")
		} else if mtype == Jpeg {
			filename = fmt.Sprintf("%s.%s", times, ".jpg")
		}
		_, err := discord.ChannelFileSend(channel, filename, bytes.NewReader(raw))
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	cli.VersionPrinter = printFullVersion

	app := cli.NewApp()
	app.Name = "discocat"
	app.Usage = "redirect a file or string to Discord"
	app.Version = version
	app.OnUsageError = handleUsageError
	app.Authors = []*cli.Author {
		&cli.Author {
			Name: "hnkz",
			Email: "hanakazu8989@gmail.com",
		},
	}
	app.Flags = []cli.Flag {
		&cli.BoolFlag {
			Name: "list",
			Aliases: []string{"l"},
			Usage: "list bot and channel names",
		},
		&cli.StringFlag {
			Name: "bot",
			Aliases: []string{"b"},
			Value: "default",
			Usage: "bot name to post",
		},
		&cli.StringFlag {
			Name: "channel",
			Aliases: []string{"c"},
			Value: "default",
			Usage: "channel name to post",
		},
		&cli.BoolFlag {
			Name: "tee",
			Aliases: []string{"t"},
			Usage: "print stdin to screen before posting",
		},
	}
	app.Action = func(c *cli.Context) error {
		var (
			botTokenKey = c.String("bot")
			channelIDKey = c.String("channel")
		)

		viper.SetConfigName("config")
		viper.AddConfigPath(defaultConfigPaths)

		if err := viper.ReadInConfig(); err != nil {
			return exitErr(err)
		}
		var discoConfig DiscordConfig
		err := viper.Unmarshal(&discoConfig)
		if err != nil {
			return exitErr(err)
		}

		if c.Bool("list") {
			discoConfig.printConfig()
			return cli.Exit("", 0)
		}

		raw, err := ioutil.ReadAll(os.Stdin)

		if c.Bool("tee") {
			fmt.Fprint(os.Stderr, string(raw))
		}

		var (
			botToken = discoConfig[botTokenKey].BotToken
			channelID = discoConfig[botTokenKey].ChannelIDs[channelIDKey]
		)

		if botToken == "" {
			return exitErr(fmt.Errorf("botToken specified by '%s' is empty. Please specify valid key of BotToken", botTokenKey))
		}
		if channelID == "" {
			return exitErr(fmt.Errorf("channelID specified by '%s' is empty. Please specify valid key of ChannelID", channelIDKey))
		}

		if err := post(raw, botToken, channelID); err != nil {
			return exitErr(err)
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		printErr(err)
		os.Exit(1)
	}
}
