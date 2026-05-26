// Package either5 provides cross type transformations for `mo.Either`.
//
// The functions provided by this package are not methods of `mo.Either` due to the lack of method type parameters
// on methods. This is part of the design decision of the Go's generics as explained here:
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#No-parameterized-methods
package either5

import (
	"github.com/samber/mo"
)

// Match returns the result of applying `onLeft` to the left value of the either or `onRight` to the right value of the either.
func Match[In1 any, In2 any, In3 any, In4 any, In5 any, Out1 any, Out2 any, Out3 any, Out4 any, Out5 any](onArg1 func(In1) Out1, onArg2 func(In2) Out2, onArg3 func(In3) Out3, onArg4 func(In4) Out4, onArg5 func(In5) Out5) func(either mo.Either5[In1, In2, In3, In4, In5]) mo.Either5[Out1, Out2, Out3, Out4, Out5] {
	_ = "STUB: not implemented"
	return nil
}

// MapArg1 executes the given function, if Either3 use the first argument, and returns result.
func MapArg1[In1 any, In2 any, In3 any, In4 any, In5 any, Out1 any](f func(In1) Out1) func(either mo.Either5[In1, In2, In3, In4, In5]) mo.Either5[Out1, In2, In3, In4, In5] {
	_ = "STUB: not implemented"
	return nil
}

// MapArg2 executes the given function, if Either3 use the second argument, and returns result.
func MapArg2[In1 any, In2 any, In3 any, In4 any, In5 any, Out2 any](f func(In2) Out2) func(either mo.Either5[In1, In2, In3, In4, In5]) mo.Either5[In1, Out2, In3, In4, In5] {
	_ = "STUB: not implemented"
	return nil
}

// MapArg3 executes the given function, if Either3 use the third argument, and returns result.
func MapArg3[In1 any, In2 any, In3 any, In4 any, In5 any, Out3 any](f func(In3) Out3) func(either mo.Either5[In1, In2, In3, In4, In5]) mo.Either5[In1, In2, Out3, In4, In5] {
	_ = "STUB: not implemented"
	return nil
}
