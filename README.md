# discocat

discocat is a simple commandline utility to post snippets to Discord.

## Quick Start

Make sure your `PATH` includes the `$GOPATH/bin` directory.

```
$ go get -u github.com/wan-nyan-wan/discocat
$ mkdir -p ~/.config/discocat
$ cp config.yml.sample ~/.config/discocat/config.yml
$ vim ~/.config/discocat/config.yml
$ echo "hello" | discocat --bot bot01
```

## Configuration

1. Create Discord bot and get bot TokenID in [Discord Develper Portal](https://discord.com/developers/applications)
2. Get ChannelID that you want to post snippets
3. write configuration in `~/.config/discocat/config.yml`(refereing to config.yml.sample`)

The below is a sample configuration. In default, discocat use `default` as a bot name. If you want to use others, use `--bot` option.

```
default:
    BotToken: "default bot token"
    ChannelID: "default channel id"
bot01:
    BotToken: "bot01 token"
    ChannelID: "channel01 id"
bot02:
    BotToken: "bot02 token"
    ChannelID: "channel02 id"
```

## Building

```
$ git clone https://github.com/wan-nyan-wan/discocat.git
$ go build
```

## Installation

```
$ go get -u github.com/wan-nyan-wan/discocat
```

## Usage

```
NAME:
   discocat - redirect a file or string to Discord

USAGE:
   discocat [global options] command [command options] [arguments...]

VERSION:
   v.1.0

AUTHOR:
   wan-nyan-wan <kazuki.hanai@wan-nyan-wan.net>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --list, -l             list bot and channel names (default: false)
   --bot value, -b value  bot name to post (default: "default")
   --tee, -t              print stdin to screen before posting (default: false)
   --help, -h             show help (default: false)
   --version, -v          print the version (default: false)
```

### Examples

````
$ discocat # post stdin(If you want to stop to input, please input Ctrl-D) to the channel via default bot
$ discocat -h # show help
$ discocat -l # show config
$ echo "aiueo" | discocat # post text to the channel via default bot
$ cat test.png | discocat # post image to the channel via default bot
$ echo "hello" | discocat --bot testbot # post testbot

```

## References

discocat is greatly inspired by [slackcat](https://github.com/bcicen/slackcat).
```
````
