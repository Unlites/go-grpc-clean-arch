package usecase

import (
	"context"
	"fmt"

	"github.com/Unlites/airport_grpc_backend/internal/domain"
)

type ticketUsecase struct {
	prRepo domain.PlaneRouteRepository
	tRepo  domain.TicketRepository
}

func NewTicketUsecase(prRepo domain.PlaneRouteRepository, tRepo domain.TicketRepository) *ticketUsecase {
	return &ticketUsecase{prRepo: prRepo, tRepo: tRepo}
}

func (uc *ticketUsecase) AddTicket(ctx context.Context, ticket *domain.Ticket) error {
	route, err := uc.prRepo.GetById(ctx, ticket.PlaneRouteId)
	if err != nil {
		return fmt.Errorf("failed to get route matching this ticket: %w", err)
	}

	if route == nil {
		return fmt.Errorf("no route matching this ticket")
	}

	if route.TicketsLeft == 0 {
		return fmt.Errorf("no tickets left")
	}

	route.TicketsLeft--

	if err := uc.prRepo.UpdateRoute(ctx, route); err != nil {
		return fmt.Errorf("failed to update route: %w", err)
	}

	if err := uc.tRepo.AddTicket(ctx, ticket); err != nil {
		return fmt.Errorf("failed to add ticket: %w", err)
	}

	return nil
}
