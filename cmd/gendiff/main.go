package main

import (
	"os"
	"context"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:		"gendiff",
		Usage:		"Compares two configuration files and shows a difference",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:		"format",
				Aliases:	[]string{"f"},
				DefaultText:	"stylish",
                		Value:		"stylish",
                		Usage:		"output format",
			},
		},
		Action: func(_ context.Context, cmd *cli.Command) error {
			return nil
		},
	}

	_ = cmd.Run(context.Background(), os.Args)
}
