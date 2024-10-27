package numbersearch_test

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/KRR19/number-search/internal/domain/numbersearch"
	"github.com/KRR19/number-search/internal/domain/numbersearch/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type test struct {
	service *numbersearch.Service
	store   *mocks.MockStore
	cfg     *mocks.MockConfig
	ctrl    *gomock.Controller
}

func NewTest(t *testing.T) *test {
	ctrl := gomock.NewController(t)
	store := mocks.NewMockStore(ctrl)
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	cfg := mocks.NewMockConfig(ctrl)
	svc := numbersearch.NewService(log, store, cfg)
	return &test{
		service: svc,
		store:   store,
		cfg:     cfg,
		ctrl:    ctrl,
	}
}

func TestSearchNumber(t *testing.T) {
	t.Run("NumberFoundEven", func(t *testing.T) {
		t.Parallel()
		test := NewTest(t)
		defer test.ctrl.Finish()
		ctx := context.TODO()

		test.store.EXPECT().SortedNumbers().Return([]int{1, 2, 3, 4, 5, 7}, nil)

		i, err := test.service.SearchNumber(ctx, 3)

		assert.NoError(t, err)
		assert.Equal(t, 2, i)
	})

	t.Run("NumberFoundOdd", func(t *testing.T) {
		t.Parallel()
		test := NewTest(t)
		defer test.ctrl.Finish()
		ctx := context.TODO()

		test.store.EXPECT().SortedNumbers().Return([]int{1, 2, 3, 4, 5, 7, 8}, nil)

		i, err := test.service.SearchNumber(ctx, 3)

		assert.NoError(t, err)
		assert.Equal(t, 2, i)
	})

	t.Run("NumberNotFound", func(t *testing.T) {
		t.Parallel()
		test := NewTest(t)
		defer test.ctrl.Finish()
		ctx := context.TODO()

		test.store.EXPECT().SortedNumbers().Return([]int{1, 20, 40, 55, 100}, nil)
		test.cfg.EXPECT().Variation().Return(10.0)

		_, err := test.service.SearchNumber(ctx, 3)

		assert.Error(t, err)
		assert.Equal(t, numbersearch.ErrNumberNotFound, err)
	})

	t.Run("ErrorGettingSortedNumbers", func(t *testing.T) {
		t.Parallel()
		test := NewTest(t)
		defer test.ctrl.Finish()
		ctx := context.TODO()

		test.store.EXPECT().SortedNumbers().Return(nil, assert.AnError)

		_, err := test.service.SearchNumber(ctx, 3)
		assert.Error(t, err)
	})

	t.Run("NumberCloseToTargetL", func(t *testing.T) {
		t.Parallel()
		test := NewTest(t)
		defer test.ctrl.Finish()
		ctx := context.TODO()

		test.store.EXPECT().SortedNumbers().Return([]int{1000, 1100, 1270, 1300, 1400}, nil)
		test.cfg.EXPECT().Variation().Return(10.0)

		i, err := test.service.SearchNumber(ctx, 1150)

		assert.NoError(t, err)
		assert.Equal(t, 1, i)
	})

	t.Run("NumberCloseToTargetR", func(t *testing.T) {
		t.Parallel()
		test := NewTest(t)
		defer test.ctrl.Finish()
		ctx := context.TODO()

		test.store.EXPECT().SortedNumbers().Return([]int{900, 1000, 1200, 1300, 1400}, nil)
		test.cfg.EXPECT().Variation().Return(10.0)

		i, err := test.service.SearchNumber(ctx, 1150)

		assert.NoError(t, err)
		assert.Equal(t, 2, i)
	})

	t.Run("EmptyList", func(t *testing.T) {
		t.Parallel()
		test := NewTest(t)
		defer test.ctrl.Finish()
		ctx := context.TODO()

		test.store.EXPECT().SortedNumbers().Return([]int{}, nil)

		_, err := test.service.SearchNumber(ctx, 3)

		assert.Error(t, err)
		assert.Equal(t, numbersearch.ErrEmptyList, err)
	})

	t.Run("SingleElementListFound", func(t *testing.T) {
		t.Parallel()
		test := NewTest(t)
		defer test.ctrl.Finish()
		ctx := context.TODO()

		test.store.EXPECT().SortedNumbers().Return([]int{3}, nil)

		i, err := test.service.SearchNumber(ctx, 3)

		assert.NoError(t, err)
		assert.Equal(t, 0, i)
	})

	t.Run("SingleElementListNotFound", func(t *testing.T) {
		t.Parallel()
		test := NewTest(t)
		defer test.ctrl.Finish()
		ctx := context.TODO()

		test.store.EXPECT().SortedNumbers().Return([]int{5}, nil)
		test.cfg.EXPECT().Variation().Return(10.0)

		_, err := test.service.SearchNumber(ctx, 3)

		assert.Error(t, err)
		assert.Equal(t, numbersearch.ErrNumberNotFound, err)
	})

	t.Run("SingleElementListCloseToTarget", func(t *testing.T) {
		t.Parallel()
		test := NewTest(t)
		defer test.ctrl.Finish()
		ctx := context.TODO()

		test.store.EXPECT().SortedNumbers().Return([]int{10}, nil)
		test.cfg.EXPECT().Variation().Return(10.0)

		i, err := test.service.SearchNumber(ctx, 11)

		assert.NoError(t, err)
		assert.Equal(t, 0, i)
	})

	t.Run("SingleElementListCloseToTarget", func(t *testing.T) {
		t.Parallel()
		test := NewTest(t)
		defer test.ctrl.Finish()
		ctx := context.TODO()

		test.store.EXPECT().SortedNumbers().Return([]int{10}, nil)
		test.cfg.EXPECT().Variation().Return(10.0)

		i, err := test.service.SearchNumber(ctx, 9)

		assert.NoError(t, err)
		assert.Equal(t, 0, i)
	})
}
