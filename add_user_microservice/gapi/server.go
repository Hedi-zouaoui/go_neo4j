package gapi

import (
	

"github.com/add_user/utils"
	"github.com/add_user/pb"
	
	
)

type Server struct {
	pb.UnimplementedAddServer
	config utils.Config
}
func NewServer(config utils.Config )(*Server , error ) {
	server := &Server{
		config:     config,
		
	}

	return server, nil
}