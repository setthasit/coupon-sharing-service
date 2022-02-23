package repositories

import (
	"context"
	"coupon-service/domains/entities"
)

type BoardUserRepository interface {
	Create(ctx context.Context, newUser *entities.BoardUser) (*entities.BoardUser, error)
}
