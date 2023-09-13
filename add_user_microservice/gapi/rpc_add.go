package gapi

import (
	"context"
	"fmt"

	//"fmt"
	"log"

	"github.com/add_user/pb"
	"github.com/add_user/utils"

	//"google.golang.org/grpc"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)


func (server *Server) Adduser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	// Create a new Neo4j driver
	driver, err := neo4j.NewDriver("bolt://neo4j_db:7687", neo4j.BasicAuth("neo4j", "12345678", ""))
	// docker : bolt://neo4j_db:7687
	if err != nil {
		log.Fatalf("Failed to create Neo4j driver: %v", err)
	}
	defer driver.Close()

	// Create a new node
	res , err := utils.CreateNode(driver, req.Name)
if err != nil {
	return nil , err
}
	transferReq := utils.Transfer_node{
		To:      req.To,
		From:    res.ID,
		Methode: "new",
	}
	fmt.Println(transferReq)
		_ = utils.TransferNode(driver, transferReq )
	if err != nil {
		response := &pb.AddUserResponse{
			NewId: 0,
			Err:   err.Error(),
		}
		return response, nil
	}
response := &pb.AddUserResponse{
			NewId: res.ID ,
			Err:   "" ,
		}

//utils.Add_json(int(req.To) , req.Name , int(response.NewId))

	return response, nil
}