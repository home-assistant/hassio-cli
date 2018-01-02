# hassiocli


## Description

Commandline interface to facilitate interaction with hass.io server

## Usage
- `hassiocli help`
- `hassiocli <subcommand> <action> [<options>]`

E.g.

- `hassiocli homeassistant info --rawjson`   

## Supported endpoints
### homeassistant
- info
- logs
- restart
- start
- stop
- update

### supervisor
- info
- logs
- reload
- update

### host
- hardware
- reboot
- shutdown
- update

### network
- options


## Supported modifiers
- --rawjson,-j -> Will return the data in JSON format on a 
                    single line (useful for passing to other 
                    programs to parse / utilise)
- --options,-o -> Used to send commands to hass.io `hassiocli homeassistant update --options "version=0.60"`
- --filter,-f  -> Used to filter the data returned from hass.io so only the specified properties are output

## Install

To install, use `go get`:

```bash
$ go get -d github.com/home-assistant/hassio-cli
```

## Contribution

1. Fork ([https://github.com/home-assistant/hassio-cli/fork](https://github.com/home-assistant/hassio-cli/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request


## Building
```bash
go test ./...
gox -osarch="linux/arm" -ldflags="-s -w" -output="hassiocli"
upx --brute hassiocli
```

## Author

[home-assistant](https://github.com/home-assistant)
