package mo

import "fmt"

const (
	either5ArgId1 = iota
	either5ArgId2
	either5ArgId3
	either5ArgId4
	either5ArgId5
)

var (
	errEither5InvalidArgumentId = fmt.Errorf("either5 argument should be between 1 and 5")
	errEither5MissingArg1       = fmt.Errorf("either5 doesn't contain expected argument 1")
	errEither5MissingArg2       = fmt.Errorf("either5 doesn't contain expected argument 2")
	errEither5MissingArg3       = fmt.Errorf("either5 doesn't contain expected argument 3")
	errEither5MissingArg4       = fmt.Errorf("either5 doesn't contain expected argument 4")
	errEither5MissingArg5       = fmt.Errorf("either5 doesn't contain expected argument 5")
)

// NewEither5Arg1 builds the first argument of the Either5 struct.
func NewEither5Arg1[T1 any, T2 any, T3 any, T4 any, T5 any](value T1) Either5[T1, T2, T3, T4, T5] {
	_ = "STUB: not implemented"
	return nil
}

// NewEither5Arg2 builds the second argument of the Either5 struct.
func NewEither5Arg2[T1 any, T2 any, T3 any, T4 any, T5 any](value T2) Either5[T1, T2, T3, T4, T5] {
	_ = "STUB: not implemented"
	return nil
}

// NewEither5Arg3 builds the third argument of the Either5 struct.
func NewEither5Arg3[T1 any, T2 any, T3 any, T4 any, T5 any](value T3) Either5[T1, T2, T3, T4, T5] {
	_ = "STUB: not implemented"
	return nil
}

// NewEither5Arg4 builds the fourth argument of the Either5 struct.
func NewEither5Arg4[T1 any, T2 any, T3 any, T4 any, T5 any](value T4) Either5[T1, T2, T3, T4, T5] {
	_ = "STUB: not implemented"
	return nil
}

// NewEither5Arg5 builds the fith argument of the Either5 struct.
func NewEither5Arg5[T1 any, T2 any, T3 any, T4 any, T5 any](value T5) Either5[T1, T2, T3, T4, T5] {
	_ = "STUB: not implemented"
	return nil
}

// Either5 respresents a value of 5 possible types.
// An instance of Either5 is an instance of either T1, T2, T3, T4, or T5.
type Either5[T1 any, T2 any, T3 any, T4 any, T5 any] struct {
	argId int8

	arg1 T1
	arg2 T2
	arg3 T3
	arg4 T4
	arg5 T5
}

// IsArg1 returns true if Either5 uses the first argument.
func (e Either5[T1, T2, T3, T4, T5]) IsArg1() bool { _ = "STUB: not implemented"; return false }

// IsArg2 returns true if Either5 uses the second argument.
func (e Either5[T1, T2, T3, T4, T5]) IsArg2() bool { _ = "STUB: not implemented"; return false }

// IsArg3 returns true if Either5 uses the third argument.
func (e Either5[T1, T2, T3, T4, T5]) IsArg3() bool { _ = "STUB: not implemented"; return false }

// IsArg4 returns true if Either5 uses the fourth argument.
func (e Either5[T1, T2, T3, T4, T5]) IsArg4() bool { _ = "STUB: not implemented"; return false }

// IsArg5 returns true if Either5 uses the fith argument.
func (e Either5[T1, T2, T3, T4, T5]) IsArg5() bool { _ = "STUB: not implemented"; return false }

// Arg1 returns the first argument of a Either5 struct.
func (e Either5[T1, T2, T3, T4, T5]) Arg1() (T1, bool) {
	_ = "STUB: not implemented"
	return *new(T1), false
}

// Arg2 returns the second argument of a Either5 struct.
func (e Either5[T1, T2, T3, T4, T5]) Arg2() (T2, bool) {
	_ = "STUB: not implemented"
	return *new(T2), false
}

// Arg3 returns the third argument of a Either5 struct.
func (e Either5[T1, T2, T3, T4, T5]) Arg3() (T3, bool) {
	_ = "STUB: not implemented"
	return *new(T3), false
}

// Arg4 returns the fourth argument of a Either5 struct.
func (e Either5[T1, T2, T3, T4, T5]) Arg4() (T4, bool) {
	_ = "STUB: not implemented"
	return *new(T4), false
}

// Arg5 returns the fith argument of a Either5 struct.
func (e Either5[T1, T2, T3, T4, T5]) Arg5() (T5, bool) {
	_ = "STUB: not implemented"
	return *new(T5), false
}

// MustArg1 returns the first argument of a Either5 struct or panics.
func (e Either5[T1, T2, T3, T4, T5]) MustArg1() T1 { _ = "STUB: not implemented"; return *new(T1) }

// MustArg2 returns the second argument of a Either5 struct or panics.
func (e Either5[T1, T2, T3, T4, T5]) MustArg2() T2 { _ = "STUB: not implemented"; return *new(T2) }

// MustArg3 returns the third argument of a Either5 struct or panics.
func (e Either5[T1, T2, T3, T4, T5]) MustArg3() T3 { _ = "STUB: not implemented"; return *new(T3) }

// MustArg4 returns the fourth argument of a Either5 struct or panics.
func (e Either5[T1, T2, T3, T4, T5]) MustArg4() T4 { _ = "STUB: not implemented"; return *new(T4) }

// MustArg5 returns the fith argument of a Either5 struct or panics.
func (e Either5[T1, T2, T3, T4, T5]) MustArg5() T5 { _ = "STUB: not implemented"; return *new(T5) }

// Unpack returns all values
func (e Either5[T1, T2, T3, T4, T5]) Unpack() (T1, T2, T3, T4, T5) {
	_ = "STUB: not implemented"
	return *new(T1), *new(T2), *new(T3), *new(T4), *new(T5)
}

// Arg1OrElse returns the first argument of a Either5 struct or fallback.
func (e Either5[T1, T2, T3, T4, T5]) Arg1OrElse(fallback T1) T1 {
	_ = "STUB: not implemented"
	return *new(T1)
}

// Arg2OrElse returns the second argument of a Either5 struct or fallback.
func (e Either5[T1, T2, T3, T4, T5]) Arg2OrElse(fallback T2) T2 {
	_ = "STUB: not implemented"
	return *new(T2)
}

// Arg3OrElse returns the third argument of a Either5 struct or fallback.
func (e Either5[T1, T2, T3, T4, T5]) Arg3OrElse(fallback T3) T3 {
	_ = "STUB: not implemented"
	return *new(T3)
}

// Arg4OrElse returns the fourth argument of a Either5 struct or fallback.
func (e Either5[T1, T2, T3, T4, T5]) Arg4OrElse(fallback T4) T4 {
	_ = "STUB: not implemented"
	return *new(T4)
}

// Arg5OrElse returns the fith argument of a Either5 struct or fallback.
func (e Either5[T1, T2, T3, T4, T5]) Arg5OrElse(fallback T5) T5 {
	_ = "STUB: not implemented"
	return *new(T5)
}

// Arg1OrEmpty returns the first argument of a Either5 struct or empty value.
func (e Either5[T1, T2, T3, T4, T5]) Arg1OrEmpty() T1 { _ = "STUB: not implemented"; return *new(T1) }

// Arg2OrEmpty returns the second argument of a Either5 struct or empty value.
func (e Either5[T1, T2, T3, T4, T5]) Arg2OrEmpty() T2 { _ = "STUB: not implemented"; return *new(T2) }

// Arg3OrEmpty returns the third argument of a Either5 struct or empty value.
func (e Either5[T1, T2, T3, T4, T5]) Arg3OrEmpty() T3 { _ = "STUB: not implemented"; return *new(T3) }

// Arg4OrEmpty returns the fourth argument of a Either5 struct or empty value.
func (e Either5[T1, T2, T3, T4, T5]) Arg4OrEmpty() T4 { _ = "STUB: not implemented"; return *new(T4) }

// Arg5OrEmpty returns the fifth argument of a Either5 struct or empty value.
func (e Either5[T1, T2, T3, T4, T5]) Arg5OrEmpty() T5 { _ = "STUB: not implemented"; return *new(T5) }

// ForEach executes the given side-effecting function, depending of the argument set.
func (e Either5[T1, T2, T3, T4, T5]) ForEach(arg1Cb func(T1), arg2Cb func(T2), arg3Cb func(T3), arg4Cb func(T4), arg5Cb func(T5)) {
	_ = "STUB: not implemented"
	return
}

// Match executes the given function, depending of the argument set, and returns result.
func (e Either5[T1, T2, T3, T4, T5]) Match(
	onArg1 func(T1) Either5[T1, T2, T3, T4, T5],
	onArg2 func(T2) Either5[T1, T2, T3, T4, T5],
	onArg3 func(T3) Either5[T1, T2, T3, T4, T5],
	onArg4 func(T4) Either5[T1, T2, T3, T4, T5],
	onArg5 func(T5) Either5[T1, T2, T3, T4, T5]) Either5[T1, T2, T3, T4, T5] {
	_ = "STUB: not implemented"
	return nil
}

// MapArg1 executes the given function, if Either5 use the first argument, and returns result.
func (e Either5[T1, T2, T3, T4, T5]) MapArg1(mapper func(T1) Either5[T1, T2, T3, T4, T5]) Either5[T1, T2, T3, T4, T5] {
	_ = "STUB: not implemented"
	return nil
}

// MapArg2 executes the given function, if Either5 use the second argument, and returns result.
func (e Either5[T1, T2, T3, T4, T5]) MapArg2(mapper func(T2) Either5[T1, T2, T3, T4, T5]) Either5[T1, T2, T3, T4, T5] {
	_ = "STUB: not implemented"
	return nil
}

// MapArg3 executes the given function, if Either5 use the third argument, and returns result.
func (e Either5[T1, T2, T3, T4, T5]) MapArg3(mapper func(T3) Either5[T1, T2, T3, T4, T5]) Either5[T1, T2, T3, T4, T5] {
	_ = "STUB: not implemented"
	return nil
}

// MapArg4 executes the given function, if Either5 use the fourth argument, and returns result.
func (e Either5[T1, T2, T3, T4, T5]) MapArg4(mapper func(T4) Either5[T1, T2, T3, T4, T5]) Either5[T1, T2, T3, T4, T5] {
	_ = "STUB: not implemented"
	return nil
}

// MapArg5 executes the given function, if Either5 use the fith argument, and returns result.
func (e Either5[T1, T2, T3, T4, T5]) MapArg5(mapper func(T5) Either5[T1, T2, T3, T4, T5]) Either5[T1, T2, T3, T4, T5] {
	_ = "STUB: not implemented"
	return nil
}
