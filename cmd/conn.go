package cmd

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func newGRPCConnection(ctx context.Context, addr string, insecureOption bool) (*grpc.ClientConn, error) {
	var dialOpts []grpc.DialOption
	if insecureOption {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	}

	dialOpts = append(dialOpts, grpc.WithBlock(), grpc.WithBackoffMaxDelay(time.Second))
	return grpc.DialContext(ctx, addr, dialOpts...)
}
