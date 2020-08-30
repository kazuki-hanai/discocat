# discocat

discocat is a simple commandline utility to post snippets to Discord

## Configuration

1. Create Discord bot and get Token ID
2. Get ChannelID for posting snippets
3. write configuration in `config.yml`(refereing to `config.yml.sample`)

## Building

```
go get github.com/wan-nyan-wan/discocat && \
cd $GOPATH/src/github.com/wan-nyan-wan/discocat && \
go build
```

## Installation

```
go install github.com/wan-nyan-wan/discocat
```

## Usage

```
NAME:
   discocat - redirect a file or string to Discord

USAGE:
   discocat [global options] command [command options] [arguments...]

VERSION:
   dev-build

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --comment value, -c value   posting comment
   --filepath value, -f value  filepath for upload
   --help, -h                  show help
   --version, -v               print the version
```

### Examples

```
$ echo "aiueo" | xargs discocat -c
$ discocat -c hello
$ discocat -f test.png
$ discocat -c hello -f test.png
```

## References

discocat is greatly inspired by [slackcat](https://github.com/bcicen/slackcat).
