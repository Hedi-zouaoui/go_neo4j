package main

import (
	"context"
	"log"
	"net"
	"net/http"
	

	"github.com/add_user/gapi"
	"github.com/add_user/pb"
	"github.com/add_user/utils"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
)

func main() {

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot log config ", err)
	}
	driver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "12345678", ""))
	if err != nil {
		log.Fatalf("Failed to create Neo4j driver: %v", err)
	}
	defer driver.Close()
	go runGatewayServer(config)
	runGrpcserver(config)
		
	
		

}

func runGatewayServer(config utils.Config ){


	
	server , err := gapi.NewServer(config )
	
	if  err!=nil{
		log.Fatal("cannot create new server ", err)
	}
	grpcMux:= runtime.NewServeMux()
	ctx , cancel := context.WithCancel(context.Background())
	defer cancel()
	err = pb.RegisterAddHandlerServer(ctx , grpcMux , server )
	if err != nil {
	log.Fatal("cannot register handler server ")
}
//////////// 
    corsMiddleware := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:8080" , "https://admin.myapp.com" , "http://localhost:8012"  }, // Change this to your frontend URL
        AllowedMethods: []string{"POST"},
    })
	//////////
mux := http.NewServeMux()
mux.Handle("/" , grpcMux)
///////////////////
 handler := corsMiddleware.Handler(mux)


listener , err := net.Listen("tcp", config.HTTPServerAddress)
if err != nil {
	log.Fatal(err)
}
log.Printf("start http gateway server at %s" , listener.Addr().String())
err = http.Serve(listener , handler)
if err != nil {
	log.Fatal("cannot start http gateway server  " , err )
}
}







func runGrpcserver(config utils.Config ){


	
	server , err := gapi.NewServer(config )
	
	if  err!=nil{
		log.Fatal("cannot create new server ", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAddServer(grpcServer , server )
reflection.Register(grpcServer) //what rpc available in the server 
listener , err := net.Listen("tcp", config.GRPCServerAddress)
if err != nil {
	log.Fatal(err)
}
log.Printf("start gRPC server at %s" , listener.Addr().String())
err = grpcServer.Serve(listener)
if err != nil {
	log.Fatal("cannot start grpc server  " , err )
}
}