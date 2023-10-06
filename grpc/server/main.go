package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "go-grpc-api/grpc/flights"
	"go-grpc-api/internal/repositories"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedFlightServiceServer
	flightRepository repositories.FlightsRepository
}

func newServer() *server {
	s := &server{
		flightRepository: *repositories.NewFlightsRepository(),
	}
	return s
}

func (s *server) List(input *pb.ListRequest, stream pb.FlightService_ListServer) error {
	log.Printf("Received: %v", input.GetFilter())

	skip := 0
	for {
		flights, total, err := s.flightRepository.List(0, 50000)
		if err != nil {
			log.Fatalf("failed to get flights: %v", err)
		}

		for _, flight := range flights {
			_f := pb.ListReply{
				Year:          flight.Year,
				Month:         flight.Month,
				DayofMonth:    flight.DayofMonth,
				DayOfWeek:     flight.DayOfWeek,
				DepTime:       flight.DepTime,
				CRSDepTime:    flight.CRSDepTime,
				ArrTime:       flight.ArrTime,
				CRSArrTime:    flight.CRSArrTime,
				UniqueCarrier: flight.UniqueCarrier,
			}

			if err := stream.Send(&_f); err != nil {
				return err
			}
		}

		skip = skip + len(flights)
		if skip >= total {
			break
		}
		log.Printf("get next %v", skip)
	}
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFlightServiceServer(s, newServer())
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
