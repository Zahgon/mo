// Package result provides cross type transformations for `mo.Result`.
//
// The functions provided by this package are not methods of `mo.Result` due to the lack of method type parameters
// on methods. This is part of the design decision of the Go's generics as explained here:
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#No-parameterized-methods
package result

import (
	"github.com/samber/mo"
)

// Map returns a new `mo.Result` wrapping the result of applying `f` to the value of result, if present, and None otherwise.
func Map[I any, O any](f func(I) O) func(result mo.Result[I]) mo.Result[O] {
	_ = "STUB: not implemented"
	return nil
}

// FlatMap returns the result of applying `f` to the value of result, if present, and None otherwise.
func FlatMap[I any, O any](f func(I) mo.Result[O]) func(result mo.Result[I]) mo.Result[O] {
	_ = "STUB: not implemented"
	return nil
}

// Match returns a new `mo.Result` from the result of applying `onValue` to the value of result, if present,
// or from the result of calling `onError` if absent.
func Match[I any, O any](onValue func(I) (O, error), onError func() (O, error)) func(result mo.Result[I]) mo.Result[O] {
	_ = "STUB: not implemented"
	return nil
}

// FlatMatch returns the result of applying `onValue` to the value of result, if present,
// or the result of `onError` if absent.
func FlatMatch[I any, O any](onValue func(I) mo.Result[O], onError func() mo.Result[O]) func(result mo.Result[I]) mo.Result[O] {
	_ = "STUB: not implemented"
	return nil
}
