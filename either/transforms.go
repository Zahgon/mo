// Package either provides cross type transformations for `mo.Either`.
//
// The functions provided by this package are not methods of `mo.Either` due to the lack of method type parameters
// on methods. This is part of the design decision of the Go's generics as explained here:
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#No-parameterized-methods
package either

import (
	"github.com/samber/mo"
)

// MapLeft returns a new `mo.Either` wrapping the result of applying `f` to the left value of the either.
func MapLeft[Lin any, R any, Lout any](f func(Lin) Lout) func(either mo.Either[Lin, R]) mo.Either[Lout, R] {
	_ = "STUB: not implemented"
	return nil
}

// MapRight returns a new `mo.Either` wrapping the result of applying `f` to the right value of the either.
func MapRight[L any, Rin any, Rout any](f func(Rin) Rout) func(either mo.Either[L, Rin]) mo.Either[L, Rout] {
	_ = "STUB: not implemented"
	return nil
}

// FlatMapLeft returns a new `mo.Either` wrapping the result of applying `f` to the left value of the either.
func FlatMapLeft[Lin any, R any, Lout any](f func(Lin) mo.Either[Lout, R]) func(either mo.Either[Lin, R]) mo.Either[Lout, R] {
	_ = "STUB: not implemented"
	return nil
}

// FlatMapRight returns a new `mo.Either` wrapping the result of applying `f` to the right value of the either.
func FlatMapRight[L any, Rin any, Rout any](f func(Rin) mo.Either[L, Rout]) func(either mo.Either[L, Rin]) mo.Either[L, Rout] {
	_ = "STUB: not implemented"
	return nil
}

// Match returns the result of applying `onLeft` to the left value of the either or `onRight` to the right value of the either.
func Match[Lin any, Rin any, Lout any, Rout any](onLeft func(Lin) Lout, onRight func(Rin) Rout) func(either mo.Either[Lin, Rin]) mo.Either[Lout, Rout] {
	_ = "STUB: not implemented"
	return nil
}

// Swap returns the left value in Right and vice versa.
func Swap[I any, O any]() func(either mo.Either[I, O]) mo.Either[O, I] {
	_ = "STUB: not implemented"
	return nil
}
