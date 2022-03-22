package cmd

import (
	"io"

	"github.com/bingoohuang/channelzcli/channelz"

	"github.com/spf13/cobra"
)

type RootCommand struct {
	cmd  *cobra.Command
	opts *channelz.Options
}

func NewRootCommand(r io.Reader, w io.Writer) *RootCommand {
	c := &RootCommand{
		cmd: &cobra.Command{
			Use:   "channelzcli",
			Short: "cli for gRPC channelz",
			RunE: func(cmd *cobra.Command, args []string) error {
				return cmd.Help()
			},
		},
		opts: &channelz.Options{
			Input:  r,
			Output: w,
		},
	}
	c.cmd.PersistentFlags().BoolVarP(&c.opts.Json, "json", "j", false, "JSON output")
	c.cmd.PersistentFlags().BoolVarP(&c.opts.Verbose, "verbose", "v", false, "verbose output")
	c.cmd.PersistentFlags().BoolVarP(&c.opts.Insecure, "insecure", "k", true, "with insecure")
	c.cmd.PersistentFlags().StringVarP(&c.opts.Address, "addr", "a", "", "address to gRPC server")
	c.cmd.AddCommand(NewListCommand(c.opts).Command())
	c.cmd.AddCommand(NewTreeCommand(c.opts).Command())
	c.cmd.AddCommand(NewDescribeCommand(c.opts).Command())
	return c
}

func (c *RootCommand) Execute() error { return c.cmd.Execute() }
