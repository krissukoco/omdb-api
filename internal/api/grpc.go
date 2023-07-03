package api

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func grpcErrMsg(code int, msg string) string {
	return fmt.Sprintf("%d__%s", code, msg)
}

// GrpcError is a helper function to create a gRPC error
// that is used internally in the application.
// The error message is formatted as <code>__<msg>.
func GrpcError(grpcCode codes.Code, code int, msg string) error {
	return status.Error(grpcCode, grpcErrMsg(code, msg))
}
