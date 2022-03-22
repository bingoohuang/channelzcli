package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/bingoohuang/channelzcli/channelz"
	"github.com/spf13/cobra"
)

type DescribeCommand struct {
	cmd  *cobra.Command
	opts *channelz.Options
	addr string
	long bool
	full bool
}

func NewDescribeCommand(opts *channelz.Options) *DescribeCommand {
	c := &DescribeCommand{
		cmd: &cobra.Command{
			Use:          "describe (channel|server|serversocket) (NAME|ID)",
			Short:        "describe (channel|server|serversocket) (NAME|ID)",
			Aliases:      []string{"desc", "d"},
			Args:         cobra.ExactArgs(2),
			SilenceUsage: true,
		},
		opts: opts,
	}
	c.cmd.RunE = c.Run
	return c
}

func (c *DescribeCommand) Command() *cobra.Command {
	return c.cmd
}

func (c *DescribeCommand) Run(_ *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	typ := args[0]
	name := args[1]

	dialCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	conn, err := newGRPCConnection(dialCtx, c.opts.Address, c.opts.Insecure)
	if err != nil {
		return fmt.Errorf("failed to connect %v: %v", c.opts.Address, err)
	}
	defer closeX(conn)

	cc := channelz.NewClient(conn, c.opts.Output)

	switch typ {
	case "channel", "c":
		return cc.DescribeChannel(c.opts, ctx, name)
	case "server", "s":
		return cc.DescribeServer(c.opts, ctx, name)
	case "serversocket", "so", "ss":
		return cc.DescribeServerSocket(c.opts, ctx, name)
	default:
		_ = c.cmd.Usage()
		os.Exit(1)
	}

	return nil
}
