package tests

import (
	"google.golang.org/grpc/test/bufconn"
)

func NewTestGrpcListener() *bufconn.Listener {
	buffsize := 1024 * 1024
	lis := bufconn.Listen(buffsize)
	return lis
}
