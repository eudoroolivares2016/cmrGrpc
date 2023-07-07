package main

import (
	"context"
	"log"
	"time"
	pb "example.com/cmrGrpc/cmrGrpc"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)
func main(){
	// connect to the gRPC server
	// withBlock means functino will not return until the conn is mate blocking step to Dial
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	// used to delay the execution of a function until the surrounding function completes 
	// basically an await
	defer conn.Close()
	client := pb.NewCmrGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// defining new collections

	var new_colls = make(map[string]int32)
	new_colls["USAP-1753101"] = 1
	new_colls["latent-reserves-in-the-swiss-nfi"] = 2
	for short_name, version := range new_colls {
		r, err := client.CreateNewCollection(ctx, &pb.NewCollection{Shortname: short_name, Version: version})
		if err != nil {
			log.Fatalf("could not create collection: %v", err)
		}
		log.Printf(`collection Details:
NAME: %s
Version: %d
CONCEPT-ID: %s`, r.GetShortname(), r.GetVersion(), r.GetConceptId())

	}

}