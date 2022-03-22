package channelz

import "io"

type Options struct {
	Address  string
	Verbose  bool
	Insecure bool
	Json     bool
	Input    io.Reader
	Output   io.Writer
}
