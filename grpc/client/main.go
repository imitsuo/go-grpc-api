package main

import (
	"context"
	"flag"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	// pb "google.golang.org/grpc/examples/helloworld/helloworld"
	pb "go-grpc-api/grpc/flights"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFlightServiceClient(conn)

	// Contact the server and print out its response.
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()

	stream, err := c.List(context.Background(), &pb.ListRequest{Filter: "x"})
	if err != nil {
		log.Fatalf("could not flights: %v", err)
	}

	log.Printf("Start")
	for {
		flight, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Flight: UniqueCarrier: %v", flight.GetUniqueCarrier())
			break
		}
		if err != nil {
			log.Fatalf("client.List failed: %v", err)
		}

		// log.Printf("Flight: UniqueCarrier: %v", flight.GetUniqueCarrier())
	}
	log.Printf("End")
	// log.Printf("Greeting: %s", r.GetMessage())
}
