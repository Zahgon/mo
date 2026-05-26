package mo

import (
	"sync"
)

// NewFuture instanciate a new future.
func NewFuture[T any](cb func(resolve func(T), reject func(error))) *Future[T] {
	_ = "STUB: not implemented"
	return nil
}

// Future represents a value which may or may not currently be available, but will be
// available at some point, or an exception if that value could not be made available.
type Future[T any] struct {
	mu sync.Mutex

	cb       func(func(T), func(error))
	cancelCb func()
	next     *Future[T]
	done     chan struct{}
	doneOnce sync.Once
	result   Result[T]
}

func (f *Future[T]) active() { _ = "STUB: not implemented"; return }

func (f *Future[T]) activeSync() { _ = "STUB: not implemented"; return }

func (f *Future[T]) resolve(value T) { _ = "STUB: not implemented"; return }

func (f *Future[T]) reject(err error) { _ = "STUB: not implemented"; return }

// Then is called when Future is resolved. It returns a new Future.
func (f *Future[T]) Then(cb func(T) (T, error)) *Future[T] { _ = "STUB: not implemented"; return nil }

// Catch is called when Future is rejected. It returns a new Future.
func (f *Future[T]) Catch(cb func(error) (T, error)) *Future[T] {
	_ = "STUB: not implemented"
	return nil
}

// Finally is called when Future is processed either resolved or rejected. It returns a new Future.
func (f *Future[T]) Finally(cb func(T, error) (T, error)) *Future[T] {
	_ = "STUB: not implemented"
	return nil
}

// Cancel cancels the Future chain.
func (f *Future[T]) Cancel() { _ = "STUB: not implemented"; return }

// Collect awaits and return result of the Future.
func (f *Future[T]) Collect() (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

// Result wraps Collect and returns a Result.
func (f *Future[T]) Result() Result[T] { _ = "STUB: not implemented"; return nil }

// Either wraps Collect and returns a Either.
func (f *Future[T]) Either() Either[error, T] { _ = "STUB: not implemented"; return nil }
