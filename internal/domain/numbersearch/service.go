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

func (s *Service) SearchNumberV2(ctx context.Context, target int) (int, error) {
	const step = 100

	if target%step == 0 {
		p := target / step
		s.log.InfoContext(ctx, "number found", "number", target, "position", p)

		return p, nil
	}

	diff := (target/step+1)*step - target
	persentege := s.cfg.Variation() / 100

	if diff <= 1000000 && float64(diff) <= float64(target)*persentege {
		p := target/step+1
		s.log.InfoContext(ctx, "number found", "number", target, "position", p)

		return p, nil
	}

	diff = target - (target/step)*step

	if diff >= 0 && float64(diff) <= float64(target)*persentege {
		p := target / step
		s.log.InfoContext(ctx, "number found", "number", target, "position", p)

		return p, nil
	}

	s.log.WarnContext(ctx, "number not found", "number", target)

	return -1, ErrNumberNotFound
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
	variation := s.cfg.Variation()

	min := float64(target) * (100 - variation) / 100
	max := float64(target) * (100 + variation) / 100

	if l < len(list) && s.targetInRange(float64(list[l]), min, max) {
		s.log.WarnContext(ctx, "number not found, but found the closest one", "number", target, "closest", list[l])

		return l
	}

	if r >= 0 && s.targetInRange(float64(list[r]), min, max) {
		s.log.WarnContext(ctx, "number not found, but found the closest one", "number", target, "closest", list[r])

		return r
	}

	return -1
}

func (s *Service) targetInRange(curr, min, max float64) bool {
	return curr >= min && curr <= max
}
