// Package grpc exist just to register the grpc extension
package grpc

import (
	"github.com/farzanhaq/xk6-grpc-xds/grpc"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/grpc", new(grpc.RootModule))
}
