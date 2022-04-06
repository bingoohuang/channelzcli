package cmd

import (
	"context"
	"time"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func newGRPCConnection(ctx context.Context, addr string, insecureOption bool) (*grpc.ClientConn, error) {
	var t grpc.DialOption
	if insecureOption {
		t = grpc.WithTransportCredentials(insecure.NewCredentials())
	} else {
		t = grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))
	}

	return grpc.DialContext(ctx, addr, t, grpc.WithBlock(), grpc.WithBackoffMaxDelay(time.Second))
}
