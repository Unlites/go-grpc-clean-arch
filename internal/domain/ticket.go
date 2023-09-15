package domain

import "context"

type Ticket struct {
	Id                 int
	PassengerFirstName string
	PassenderLastName  string
	PlaneRouteId       int
}

type TicketUsecase interface {
	AddTicket(ctx context.Context, ticket *Ticket) error
}

type TicketRepository interface {
	AddTicket(ctx context.Context, ticket *Ticket) error
}
