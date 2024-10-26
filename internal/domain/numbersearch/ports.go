package numbersearch

//go:generate mockgen -source=ports.go -destination=./mocks/store_mock.go -package=mocks
type Store interface {
	SortedNumbers() ([]int, error)
}
