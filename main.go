package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"

	"github.com/lwmacct/251021-study-redis/app/version"
)

// buildCommands 根据环境变量条件性构建命令列表
func buildCommands() []*cli.Command {
	commands := []*cli.Command{
		version.Command,
	}
	return commands
}

func main() {
	cmd := &cli.Command{
		Usage:    version.AppRawName,
		Flags:    []cli.Flag{},
		Commands: buildCommands(),
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
