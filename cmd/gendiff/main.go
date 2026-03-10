package main

import (
	"context"
	"fmt"
	"log"
	"os"

	cli "github.com/urfave/cli/v3"
	gendiff "code"
	"code/internal/parser"
)

func main() {
	cmd := &cli.Command{
		Name:      "gendiff",
		Usage:     "Compares two configuration files and shows a difference.",
		ArgsUsage: "[filepath1] [filepath2]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Value:   "stylish",
				Usage:   "output format",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			args := cmd.Args()
			if args.Len() < 2 {
				return fmt.Errorf("two file paths are required")
			}

			data1, err := parser.Parse(args.Get(0))
			if err != nil {
				return err
			}

			data2, err := parser.Parse(args.Get(1))
			if err != nil {
				return err
			}

			result, err := gendiff.GenDiff(data1, data2, cmd.String("format"))
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
