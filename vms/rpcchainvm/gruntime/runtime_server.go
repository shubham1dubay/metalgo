// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gruntime

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/shubham1dubay/metalgo/vms/rpcchainvm/runtime"

	pb "github.com/shubham1dubay/metalgo/proto/pb/vm/runtime"
)

var _ pb.RuntimeServer = (*Server)(nil)

// Server is a VM runtime initializer controlled by RPC.
type Server struct {
	pb.UnsafeRuntimeServer
	runtime runtime.Initializer
}

func NewServer(runtime runtime.Initializer) *Server {
	return &Server{
		runtime: runtime,
	}
}

func (s *Server) Initialize(ctx context.Context, req *pb.InitializeRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.runtime.Initialize(ctx, uint(req.ProtocolVersion), req.Addr)
}
