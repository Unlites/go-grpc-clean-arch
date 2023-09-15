package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Unlites/airport_grpc_backend/internal/domain"
)

type planeRouteInMemoryRepository struct {
	planeRoutesStore []*domain.PlaneRoute
}

func NewPlaneRouteInMemoryRepository() *planeRouteInMemoryRepository {
	return &planeRouteInMemoryRepository{planeRoutesStore: make([]*domain.PlaneRoute, 0)}
}

func (pr *planeRouteInMemoryRepository) GetPlaneRoutes(ctx context.Context, limit, offset int) ([]*domain.PlaneRoute, error) {
	if len(pr.planeRoutesStore) < offset {
		return nil, nil
	}

	if len(pr.planeRoutesStore) < offset+limit {
		limit = len(pr.planeRoutesStore) - offset
	}

	return pr.planeRoutesStore[offset : offset+limit], nil
}

func (pr *planeRouteInMemoryRepository) AddRoute(ctx context.Context, planeRoute *domain.PlaneRoute) error {
	planeRoute.Id = int(time.Now().Unix())
	pr.planeRoutesStore = append(pr.planeRoutesStore, planeRoute)
	return nil
}

func (pr *planeRouteInMemoryRepository) GetById(ctx context.Context, id int) (*domain.PlaneRoute, error) {
	for _, route := range pr.planeRoutesStore {
		if route.Id == id {
			return route, nil
		}
	}

	return nil, nil
}

func (pr *planeRouteInMemoryRepository) UpdateRoute(ctx context.Context, planeRoute *domain.PlaneRoute) error {
	for i, storedRoute := range pr.planeRoutesStore {
		if storedRoute.Id == planeRoute.Id {
			pr.planeRoutesStore[i] = planeRoute
			return nil
		}
	}

	return fmt.Errorf("route not found")
}
