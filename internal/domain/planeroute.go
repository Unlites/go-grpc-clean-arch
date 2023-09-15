package domain

import (
	"context"
	"time"
)

type PlaneRoute struct {
	Id            int
	PlaneId       int
	From          string
	To            string
	Status        string
	TicketsLeft   int
	ArrivalTime   time.Time
	DepartureTime time.Time
}

type PlaneRouteUsecase interface {
	AddPlaneRoute(ctx context.Context, planeRoute *PlaneRoute) error
	GetCurrentPlaneRoutes(ctx context.Context, limit, offset int) ([]*PlaneRoute, error)
}

type PlaneRouteRepository interface {
	GetPlaneRoutes(ctx context.Context, limit, offset int) ([]*PlaneRoute, error)
	AddRoute(ctx context.Context, planeRoute *PlaneRoute) error
	GetById(ctx context.Context, id int) (*PlaneRoute, error)
	UpdateRoute(ctx context.Context, planeRoute *PlaneRoute) error
}
