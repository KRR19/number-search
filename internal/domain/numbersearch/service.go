package numbersearch

import (
	"context"
	"log/slog"

	"github.com/pkg/errors"
)

type Service struct {
	log   *slog.Logger
	store Store
}

func NewService(log *slog.Logger, store Store) *Service {
	return &Service{
		log:   log,
		store: store,
	}
}

func (s *Service) SearchNumber(ctx context.Context, target int) (int, error) {
	list, err := s.store.SortedNumbers()
	if err != nil {
		s.log.ErrorContext(ctx, "error getting sorted numbers", "error", err)

		return -1, errors.Wrap(err, "error getting sorted numbers")
	}

	r := s.numberPosition(list, target)
	if r == -1 {
		s.log.WarnContext(ctx, "number not found", "number", target)

		return -1, ErrNumberNotFound
	}

	return r, nil
}

func (s *Service) numberPosition(list []int, target int) int {
	for l, r := 0, len(list)-1; l <= r; {
		m := l + (r-l)/2

		if list[m] == target {
			return m
		}

		if list[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
}
