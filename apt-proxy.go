package main

import (
	"github.com/soulteary/apt-proxy/cli"
)

func main() {
	flags := cli.ParseFlags()
	cli.Daemon(&flags)
}
