package usecase

import (
	"context"
	"fmt"

	"github.com/Unlites/airport_grpc_backend/internal/domain"
)

type planeRouteUsecase struct {
	prRepo domain.PlaneRouteRepository
}

func NewPlaneRouteUsecase(prRepo domain.PlaneRouteRepository) *planeRouteUsecase {
	return &planeRouteUsecase{prRepo: prRepo}
}

func (uc *planeRouteUsecase) AddPlaneRoute(ctx context.Context, planeRoute *domain.PlaneRoute) error {
	if err := uc.prRepo.AddRoute(ctx, planeRoute); err != nil {
		return fmt.Errorf("failed to add route: %w", err)
	}

	return nil
}

func (uc *planeRouteUsecase) GetCurrentPlaneRoutes(ctx context.Context, limit, offset int) ([]*domain.PlaneRoute, error) {
	routes, err := uc.prRepo.GetPlaneRoutes(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get routes: %w", err)
	}

	return routes, nil
}
