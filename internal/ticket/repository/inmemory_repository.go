package repository

import (
	"context"
	"time"

	"github.com/Unlites/airport_grpc_backend/internal/domain"
)

type ticketInMemoryRepository struct {
	ticketsStore []*domain.Ticket
}

func NewTicketInMemoryRepository() *ticketInMemoryRepository {
	return &ticketInMemoryRepository{ticketsStore: make([]*domain.Ticket, 0)}
}

func (pr *ticketInMemoryRepository) AddTicket(ctx context.Context, ticket *domain.Ticket) error {
	ticket.Id = int(time.Now().Unix())
	pr.ticketsStore = append(pr.ticketsStore, ticket)
	return nil
}
