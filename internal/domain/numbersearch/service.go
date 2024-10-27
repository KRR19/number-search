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

	if len(list) == 0 {
		s.log.ErrorContext(ctx, "empty list")

		return -1, ErrEmptyList
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
	precision := s.cfg.Precision()

	min := float64(target) * (100 - precision) / 100
	max := float64(target) * (100 + precision) / 100

	if s.targetInRange(float64(list[l]), min, max) {
		s.log.WarnContext(ctx, "number not found, but found the closest one", "number", target, "closest", list[l-1])

		return l
	}

	if s.targetInRange(float64(list[r]), min, max) {
		s.log.WarnContext(ctx, "number not found, but found the closest one", "number", target, "closest", list[r+1])

		return r
	}

	return -1
}

func (s *Service) targetInRange(curr, min, max float64) bool {
	return curr >= min && curr <= max
}
