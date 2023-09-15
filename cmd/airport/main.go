package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	ph "github.com/Unlites/airport_grpc_backend/internal/planeroute/handlers/grpc"
	phProto "github.com/Unlites/airport_grpc_backend/internal/planeroute/handlers/grpc/proto"
	pr "github.com/Unlites/airport_grpc_backend/internal/planeroute/repository"
	pu "github.com/Unlites/airport_grpc_backend/internal/planeroute/usecase"

	th "github.com/Unlites/airport_grpc_backend/internal/ticket/handlers/grpc"
	thProto "github.com/Unlites/airport_grpc_backend/internal/ticket/handlers/grpc/proto"
	tr "github.com/Unlites/airport_grpc_backend/internal/ticket/repository"
	tu "github.com/Unlites/airport_grpc_backend/internal/ticket/usecase"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	s := grpc.NewServer()
	reflection.Register(s)

	planeRouteRepository := pr.NewPlaneRouteInMemoryRepository()
	planeRouteUsecase := pu.NewPlaneRouteUsecase(planeRouteRepository)
	planeRouteHandler := ph.NewPlaneRouteHandler(planeRouteUsecase)
	phProto.RegisterPlaneRouteServer(s, planeRouteHandler)

	ticketRepository := tr.NewTicketInMemoryRepository()
	ticketUsecase := tu.NewTicketUsecase(planeRouteRepository, ticketRepository)
	ticketHandler := th.NewTicketHandler(ticketUsecase)
	thProto.RegisterTicketServer(s, ticketHandler)

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		<-sigCh
		s.GracefulStop()
		wg.Done()
	}()

	err = s.Serve(l)
	if err != nil {
		log.Fatalf("could not serve: %v", err)
	}
	wg.Wait()
}
