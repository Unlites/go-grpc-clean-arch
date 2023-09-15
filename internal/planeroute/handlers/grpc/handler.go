package grpc

import (
	"context"
	"time"

	"github.com/Unlites/airport_grpc_backend/internal/domain"
	"github.com/Unlites/airport_grpc_backend/internal/planeroute/handlers/grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PlaneRouteHandler struct {
	proto.UnimplementedPlaneRouteServer
	uc domain.PlaneRouteUsecase
}

func NewPlaneRouteHandler(uc domain.PlaneRouteUsecase) *PlaneRouteHandler {
	return &PlaneRouteHandler{uc: uc}
}

func (h *PlaneRouteHandler) AddPlaneRoute(
	ctx context.Context,
	req *proto.PlaneRouteRequest,
) (*proto.ResultResponse, error) {
	planeRoute := &domain.PlaneRoute{
		PlaneId:       int(req.PlaneId),
		From:          req.From,
		To:            req.To,
		Status:        req.Status,
		TicketsLeft:   int(req.TicketsLeft),
		ArrivalTime:   req.ArrivalTime.AsTime(),
		DepartureTime: req.DepartureTime.AsTime(),
	}

	if err := h.uc.AddPlaneRoute(ctx, planeRoute); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.ResultResponse{Message: "Created"}, nil
}

func (h *PlaneRouteHandler) StreamCurrentRoutes(
	req *proto.EmptyRequest,
	stream proto.PlaneRoute_StreamCurrentRoutesServer,
) error {
	ticker := time.NewTicker(time.Second * 1)
	offset, limit := 0, 10

	for {
		select {
		case <-ticker.C:
			routes, err := h.uc.GetCurrentPlaneRoutes(stream.Context(), limit, offset)
			if err != nil {
				return status.Error(codes.Internal, err.Error())
			}

			if routes == nil {
				continue
			}

			for _, route := range routes {
				if err := stream.Send(&proto.PlaneRouteResponse{
					Id:            int32(route.Id),
					PlaneId:       int32(route.PlaneId),
					From:          route.From,
					To:            route.To,
					Status:        route.Status,
					TicketsLeft:   int32(route.TicketsLeft),
					ArrivalTime:   timestamppb.New(route.ArrivalTime),
					DepartureTime: timestamppb.New(route.DepartureTime),
				}); err != nil {
					return status.Error(codes.Internal, err.Error())
				}
			}

			limit += len(routes)
			offset += len(routes)
		}
	}
}
