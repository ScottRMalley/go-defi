package tokens

import (
	cli "github.com/jawher/mow.cli"
)

func Command(cmd *cli.Cmd) {
	cmd.Command("metadata", "fetch token information from public token lists", metadataCommand)
}
