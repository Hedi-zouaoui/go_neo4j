package gapi

import (
	"github.com/neo4j_golang/pb"
	"github.com/neo4j_golang/utils"
	
)

type Server struct {
	pb.UnimplementedTransferServer 
	config utils.Config
}
func NewServer(config utils.Config )(*Server , error ) {
	server := &Server{
		config:     config,
		
	}

	return server, nil
}