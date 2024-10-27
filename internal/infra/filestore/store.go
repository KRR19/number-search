package filestore

import (
	"bufio"
	"os"
	"strconv"
)

type Store struct {
	numbers []int
}

func NewStore() *Store {

	return &Store{
		numbers: make([]int, 0),
	}
}

func (s *Store) ReadFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		v, err := strconv.Atoi(t)
		if err != nil {
			return err
		}

		s.numbers = append(s.numbers, v)
	}

	return nil
}

func (s *Store) SortedNumbers() ([]int, error) {
	return s.numbers, nil
}
