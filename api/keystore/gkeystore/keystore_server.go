// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gkeystore

import (
	"context"

	"github.com/shubham1dubay/metalgo/api/keystore"
	"github.com/shubham1dubay/metalgo/database"
	"github.com/shubham1dubay/metalgo/database/rpcdb"
	"github.com/shubham1dubay/metalgo/vms/rpcchainvm/grpcutils"

	keystorepb "github.com/shubham1dubay/metalgo/proto/pb/keystore"
	rpcdbpb "github.com/shubham1dubay/metalgo/proto/pb/rpcdb"
)

var _ keystorepb.KeystoreServer = (*Server)(nil)

// Server is a snow.Keystore that is managed over RPC.
type Server struct {
	keystorepb.UnsafeKeystoreServer
	ks keystore.BlockchainKeystore
}

// NewServer returns a keystore connected to a remote keystore
func NewServer(ks keystore.BlockchainKeystore) *Server {
	return &Server{
		ks: ks,
	}
}

func (s *Server) GetDatabase(
	_ context.Context,
	req *keystorepb.GetDatabaseRequest,
) (*keystorepb.GetDatabaseResponse, error) {
	db, err := s.ks.GetRawDatabase(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	closer := dbCloser{Database: db}

	serverListener, err := grpcutils.NewListener()
	if err != nil {
		return nil, err
	}

	server := grpcutils.NewServer()
	closer.closer.Add(server)
	rpcdbpb.RegisterDatabaseServer(server, rpcdb.NewServer(&closer))

	// start the db server
	go grpcutils.Serve(serverListener, server)

	return &keystorepb.GetDatabaseResponse{ServerAddr: serverListener.Addr().String()}, nil
}

type dbCloser struct {
	database.Database
	closer grpcutils.ServerCloser
}

func (db *dbCloser) Close() error {
	err := db.Database.Close()
	db.closer.Stop()
	return err
}
