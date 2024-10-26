package numbersearch

import "context"

//go:generate mockgen -source=ports.go -destination=./mocks/store_mock.go -package=mocks -mock_names=Store=StoreMock
//go:generate mockgen -source=ports.go -destination=./mocks/log_mock.go -package=mocks -mock_names=Logger=LoggerMock
type Store interface {
	SortedNumbers() ([]int, error)
}

type Logger interface {
	InfoContext(ctx context.Context, msg string, args ...any)
	WarnContext(ctx context.Context, msg string, args ...any)
	ErrorContext(ctx context.Context, msg string, args ...any)
}
