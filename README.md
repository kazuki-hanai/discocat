# discocat

discocat is a simple commandline utility to post snippets to Discord

## Quick Start

Make sure your `PATH` includes the `$GOPATH/bin` directory.

```
$ go get http://github.com/wan-nyan-wan/discocat
$ mkdir -p ~/.config/discocat
$ cat <<EOF > ~/.config/discocat/config.yml
token: "your discord bot token"
channelID: "your discord channel ID"
EOF
$ echo "hello" | discocat
```

## Configuration

1. Create Discord bot and get Token ID
2. Get ChannelID for posting snippets
3. write configuration in `config.yml`(refereing to `config.yml.sample`)

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
   dev-build

AUTHOR:
   hnkz <hanakazu8989@gmail.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --configigure, -c  [NOT IMPREMENTED] Configure discocat (default: false)
   --list, -l         [NOT IMPREMENTED] List bot and channel names (default: false)
   --tee, -t          [NOT IMPREMENTED] Print stdin to screen before posting (default: false)
   --help, -h         show help (default: false)
   --version, -v      print the version (default: false)
```

### Examples

```
$ echo "aiueo" | discocat
$ cat test.png | discocat
```

## References

discocat is greatly inspired by [slackcat](https://github.com/bcicen/slackcat).
