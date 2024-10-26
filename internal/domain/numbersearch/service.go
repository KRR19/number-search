package numbersearch

import (
	"context"

	"github.com/pkg/errors"
)

type Service struct {
	log   Logger
	store Store
	cfg   Config
}

func NewService(log Logger, store Store, cfg Config) *Service {
	return &Service{
		log:   log,
		store: store,
		cfg:   cfg,
	}
}

func (s *Service) SearchNumber(ctx context.Context, target int) (int, error) {
	list, err := s.store.SortedNumbers()
	if err != nil {
		s.log.ErrorContext(ctx, "error getting sorted numbers", "error", err)

		return -1, errors.Wrap(err, "error getting sorted numbers")
	}

	r := s.numberPosition(ctx, list, target)
	if r == -1 {
		s.log.WarnContext(ctx, "number not found", "number", target)

		return -1, ErrNumberNotFound
	}

	s.log.InfoContext(ctx, "number found", "number", target, "position", r)

	return r, nil
}

func (s *Service) numberPosition(ctx context.Context, list []int, target int) int {
	precision := s.cfg.Precision()
	l, r := 0, len(list)-1
	for l <= r {
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

	if l != 0 && s.targetInRange(list[l-1], target*(100-precision)/100, target) {
		s.log.WarnContext(ctx, "number not found, but found the closest one", "number", target, "closest", list[l-1])

		return l - 1
	}

	if r < len(list)-1 && s.targetInRange(list[r+1], target, target*(100+precision)/100) {
		s.log.WarnContext(ctx, "number not found, but found the closest one", "number", target, "closest", list[r+1])

		return r + 1
	}

	return -1
}

func (s *Service) targetInRange(target, min, max int) bool {
	return target >= min && target <= max
}
