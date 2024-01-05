package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"

	pb "example.com/cmrGrpc/cmrGrpc"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type CmrGrpcServer struct {
	pb.UnimplementedCmrGrpcServer
}

// reciever pointer of type CmrGrpcServer
// we are overriding the function that we declared in the proto buffer so that
// we know what kind of
func (s *CmrGrpcServer) CreateNewCollection(ctx context.Context, in *pb.NewCollection) (*pb.Collection, error) {
	log.Printf("Received: %v", in.GetShortname())

	var concept_id_num int32 = int32(rand.Intn(90000) + 10000)

	concept_id_str := fmt.Sprintf("%d", concept_id_num)

	// Create the base string
	collectionHeader := "C"

	// Append the random string to the base string
	concept_id := collectionHeader + concept_id_str

	return &pb.Collection{Shortname: in.GetShortname(), Version: in.GetVersion(), ConceptId: concept_id}, nil
}

func main() {
	log.Printf("Staring up the main method of the server")
	// use net to listen over this port for incoming messages
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start up the listener %v", err)
	}
	server := grpc.NewServer()
	// second argument is the address to the type
	// TODO: note the last thing has a map as part of the struct
	pb.RegisterCmrGrpcServer(server, &CmrGrpcServer{})
	log.Printf("Server listening on: %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

// implementing the gRPC reciver method from the proto file
