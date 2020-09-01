package main

import (
	"os"
	"fmt"
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
	version = "dev-build"
	defaultConfigPaths = os.Getenv("HOME") + "/.config/discocat/"
	cmdConfig CommandConfig
	discoConfig DiscordConfig
)

func handleUsageError(c *cli.Context, err error, _ bool) error {
	printErr(fmt.Errorf("%s %s", "Incorrect Usage.", err.Error()))
	cli.ShowAppHelp(c)
	return cli.NewExitError("", 1)
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
	} else {
		if format == "png" {
			return Png, nil
		} else if format == "Gif" {
			return Gif, nil
		} else if format == "Jpeg" {
			return Jpeg, nil
		} else {
			return -1, cli.NewExitError("Could not detect filetype", 1)
		}
	}
}

func post(raw []byte) error {
	discord, err := discordgo.New("Bot " + discoConfig.Token)
	if err != nil {
		return err
	}

	mtype, err := detectMessageType(raw)
	if err != nil {
		return err
	}

	if mtype == Text {
		_, err := discord.ChannelMessageSend(discoConfig.ChannelID, string(raw))
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
		discord.ChannelFileSend(discoConfig.ChannelID, filename, bytes.NewReader(raw))
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
			Name: "configigure",
			Aliases: []string{"c"},
			Usage: "[NOT IMPREMENTED] Configure discocat",
		},
		&cli.BoolFlag {
			Name: "list",
			Aliases: []string{"l"},
			Usage: "[NOT IMPREMENTED] List bot and channel names",
		},
		&cli.BoolFlag {
			Name: "tee",
			Aliases: []string{"t"},
			Usage: "[NOT IMPREMENTED] Print stdin to screen before posting",
		},
	}
	app.Action = func(c *cli.Context) error {
		cmdConfig.filepath = c.String("")
		cmdConfig.comment = c.String("comment")

		viper.SetConfigName("config")
		viper.AddConfigPath(defaultConfigPaths)

		if err := viper.ReadInConfig(); err != nil {
			return err
		}
		err := viper.Unmarshal(&discoConfig)
		if err != nil {
			return err
		}


		raw, err := ioutil.ReadAll(os.Stdin)

		if post(raw) != nil {
			return err
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		exitErr(err)
	}
}
