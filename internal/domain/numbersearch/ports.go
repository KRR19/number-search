package numbersearch

import "context"

//go:generate mockgen -source=ports.go -destination=./mocks/mock.go -package=mocks
type Store interface {
	SortedNumbers() ([]int, error)
}

type Logger interface {
	InfoContext(ctx context.Context, msg string, args ...any)
	WarnContext(ctx context.Context, msg string, args ...any)
	ErrorContext(ctx context.Context, msg string, args ...any)
}

type Config interface {
	Variation() float64
}
