package main

import (
	"os"
	"fmt"
	"io/ioutil"

	"github.com/bwmarrin/discordgo"
	"github.com/urfave/cli/v2"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
	"github.com/fatih/color"
)

var (
	commandName = "discocat"
	build   = ""
	version = "dev-build"
	cmdConfig CommandConfig
	discoConfig DiscordConfig
)

func printErr(err error) {
	red  := color.New(color.FgRed).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Println(cyan(commandName), red(err.Error()))
}

func exitErr(err error) {
	printErr(err)
	os.Exit(1)
}

func handleUsageError(c *cli.Context, err error, _ bool) error {
	printErr(fmt.Errorf("%s %s", "Incorrect Usage.", err.Error()))
	cli.ShowAppHelp(c)
	return cli.NewExitError("", 1)
}

func printFullVersion(c *cli.Context) {
	fmt.Fprintf(c.App.Writer, "%v version %v, build %v\n", c.App.Name, c.App.Version, build)
}

func post() error {
	discord, err := discordgo.New("Bot " + discoConfig.Token)
	if err != nil {
		return err
	}
	if cmdConfig.comment != "" {
		_, err := discord.ChannelMessageSend(discoConfig.ChannelID, cmdConfig.comment)
		if err != nil {
			return err
		}
	}
	if cmdConfig.filepath != "" {
		file, err := os.Open(cmdConfig.filepath)
		if err != nil {
			return err
		}
		discord.ChannelFileSend(discoConfig.ChannelID, cmdConfig.filepath, file)
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
	app.Flags = []cli.Flag {
		&cli.StringFlag {
			Name: "comment, c",
			Usage: "posting comment",
		},
		&cli.StringFlag {
			Name: "filepath, f",
			Usage: "filepath for upload",
		},
	}
	app.Action = func(c *cli.Context) error {
		cmdConfig.filepath = c.String("filepath")
		cmdConfig.comment = c.String("comment")

		viper.SetConfigName("config")
		viper.AddConfigPath(".")

		if err := viper.ReadInConfig(); err != nil {
			return err
		}
		err := viper.Unmarshal(&discoConfig)
		if err != nil {
			return err
		}

		if cmdConfig.filepath == "" && cmdConfig.comment == "" {
			cli.ShowAppHelp(c)
			return fmt.Errorf("Specify at least one --comment or --filepath option")
		}


		if !terminal.IsTerminal(0) {
			// TODO: cooperate with pipe
			_, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
		}

		if post() != nil {
			return err
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		exitErr(err)
	}
}
