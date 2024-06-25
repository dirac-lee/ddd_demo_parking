package domain

import "context"

type IDGen interface {
	GetID(ctx context.Context) int64
}
