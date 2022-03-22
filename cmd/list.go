package cmd

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/bingoohuang/channelzcli/channelz"
	"github.com/spf13/cobra"
)

type ListCommand struct {
	cmd  *cobra.Command
	opts *channelz.Options
	addr string
	long bool
	full bool
}

func NewListCommand(opts *channelz.Options) *ListCommand {
	c := &ListCommand{
		cmd: &cobra.Command{
			Use:          "list (channel|server|serversocket)",
			Short:        "list (channel|server|serversocket)",
			Args:         cobra.ExactArgs(1),
			Aliases:      []string{"ls"},
			SilenceUsage: true,
		},
		opts: opts,
	}
	c.cmd.RunE = c.Run
	return c
}

func (c *ListCommand) Command() *cobra.Command {
	return c.cmd
}

func closeX(conn io.Closer) {
	if err := conn.Close(); err != nil {
		log.Printf("close failed: %v", err)
	}
}

func (c *ListCommand) Run(_ *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	typ := args[0]

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
		return cc.ListTopChannels(c.opts, ctx)
	case "server", "s":
		return cc.ListServers(c.opts, ctx)
	case "serversocket", "so", "ss":
		cc.ListServerSockets(ctx)
	default:
		_ = c.cmd.Usage()
		os.Exit(1)
	}

	return nil
}
