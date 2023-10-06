package main

import (
	"go-grpc-api/internal/repositories"
	"log"
)

func main() {
	flightRepository := repositories.NewFlightsRepository()
	flights, total, err := flightRepository.List(0, 1000)
	if err != nil {
		log.Printf("Failed to get Flights: %s\n", err.Error())
	}
	log.Printf("total %v", total)
	log.Printf("fligts %v", len(flights))
}
