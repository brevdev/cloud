package collections

import (
	"fmt"

	"github.com/cenkalti/backoff/v4"
)

func Flatten[T any](listOfLists [][]T) []T {
	result := []T{}
	for _, list := range listOfLists {
		result = append(result, list...)
	}
	return result
}

func GroupBy[K comparable, A any](list []A, keyGetter func(A) K) map[K][]A {
	result := map[K][]A{}
	for _, item := range list {
		key := keyGetter(item)
		result[key] = append(result[key], item)
	}
	return result
}

func MapE[T, R any](items []T, mapper func(T) (R, error)) ([]R, error) {
	results := []R{}
	for _, item := range items {
		res, err := mapper(item)
		if err != nil {
			return results, err
		}
		results = append(results, res)
	}
	return results, nil
}

func GetMapValues[K comparable, V any](m map[K]V) []V {
	values := []V{}
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// loops over list and returns when has returns true
func ListHas[K any](list []K, has func(l K) bool) bool {
	k := Find(list, has)
	return k != nil
}

func MapHasKey[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

func ListContains[K comparable](list []K, item K) bool {
	return ListHas(list, func(l K) bool { return l == item })
}

func Find[T any](list []T, f func(T) bool) *T {
	for _, item := range list {
		if f(item) {
			return &item
		}
	}
	return nil
}

// returns those that are true
func Filter[T any](list []T, f func(T) bool) []T {
	result := []T{}
	for _, item := range list {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}

func Ptr[T any](x T) *T {
	return &x
}

func RetryWithDataAndAttemptCount[T any](o backoff.OperationWithData[T], b backoff.BackOff) (T, error) {
	attemptCount := 0
	t, err := backoff.RetryWithData(func() (T, error) {
		attemptCount++
		return o()
	}, b)
	if err != nil {
		return t, fmt.Errorf("attemptCount %d: %w", attemptCount, err)
	}
	return t, nil
}
