package grpc

import (
	"context"

	"github.com/Unlites/airport_grpc_backend/internal/domain"
	"github.com/Unlites/airport_grpc_backend/internal/ticket/handlers/grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TicketHandler struct {
	proto.UnimplementedTicketServer
	uc domain.TicketUsecase
}

func NewTicketHandler(uc domain.TicketUsecase) *TicketHandler {
	return &TicketHandler{uc: uc}
}

func (h *TicketHandler) AddTicket(
	ctx context.Context,
	req *proto.TicketRequest,
) (*proto.ResultResponse, error) {
	ticket := &domain.Ticket{
		PlaneRouteId:       int(req.PlaneRouteId),
		PassengerFirstName: req.PassengerFirstName,
		PassenderLastName:  req.PassengerLastName,
	}

	if err := h.uc.AddTicket(ctx, ticket); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.ResultResponse{Message: "Created"}, nil
}
