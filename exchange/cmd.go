package exchange

import (
	cli "github.com/jawher/mow.cli"
)

func Command(cmd *cli.Cmd) {
	cmd.Command("prices", "fetch prices from dex's", pricesCommand)
}
