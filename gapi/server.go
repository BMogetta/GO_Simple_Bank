package gapi

import (
	"fmt"

	db "github.com/BMogetta/goBank_go-course/postgres/sqlc"
	pb "github.com/BMogetta/goBank_go-course/proto_go"
	"github.com/BMogetta/goBank_go-course/token"
	"github.com/BMogetta/goBank_go-course/util"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedGoBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
