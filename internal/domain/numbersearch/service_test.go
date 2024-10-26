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
	}
}

func TestSearchNumber(t *testing.T) {
	test := NewTest(t)
	ctx := context.TODO()
	t.Run("NumberFoundEven", func(t *testing.T) {
		t.Parallel()
		test.store.EXPECT().SortedNumbers().Return([]int{1, 2, 3, 4, 5, 7}, nil)
		test.cfg.EXPECT().Precision().Return(10)
		i, err := test.service.SearchNumber(ctx, 3)
		assert.NoError(t, err)
		assert.Equal(t, 2, i)
	})

	t.Run("NumberFoundOdd", func(t *testing.T) {
		t.Parallel()
		test.store.EXPECT().SortedNumbers().Return([]int{1, 2, 3, 4, 5, 7, 8}, nil)
		test.cfg.EXPECT().Precision().Return(10)
		i, err := test.service.SearchNumber(ctx, 3)
		assert.NoError(t, err)
		assert.Equal(t, 2, i)
	})

	t.Run("NumberNotFound", func(t *testing.T) {
		t.Parallel()
		test.store.EXPECT().SortedNumbers().Return([]int{1, 20, 40, 55, 100}, nil)
		test.cfg.EXPECT().Precision().Return(10)
		_, err := test.service.SearchNumber(ctx, 3)
		assert.Error(t, err)
		assert.Equal(t, numbersearch.ErrNumberNotFound, err)
	})

	t.Run("ErrorGettingSortedNumbers", func(t *testing.T) {
		t.Parallel()
		test.store.EXPECT().SortedNumbers().Return(nil, assert.AnError)
		_, err := test.service.SearchNumber(ctx, 3)
		assert.Error(t, err)
	})

	t.Run("NumberCloseToTargetL", func(t *testing.T) {
		t.Parallel()
		test.store.EXPECT().SortedNumbers().Return([]int{1000, 1100, 1200, 1300, 1400}, nil)
		test.cfg.EXPECT().Precision().Return(10)
		i, err := test.service.SearchNumber(ctx, 1150)
		assert.NoError(t, err)
		assert.Equal(t, 1, i)
	})

	t.Run("NumberCloseToTargetR", func(t *testing.T) {
		t.Parallel()
		test.store.EXPECT().SortedNumbers().Return([]int{900, 1000, 1200, 1300, 1400}, nil)
		test.cfg.EXPECT().Precision().Return(10)
		i, err := test.service.SearchNumber(ctx, 1150)
		assert.NoError(t, err)
		assert.Equal(t, 2, i)
	})
}
