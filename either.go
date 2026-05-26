package mo

import "fmt"

var errEitherShouldBeLeftOrRight = fmt.Errorf("either should be Left or Right")
var errEitherMissingLeftValue = fmt.Errorf("no such Left value")
var errEitherMissingRightValue = fmt.Errorf("no such Right value")

// Left builds the left side of the Either struct, as opposed to the Right side.
func Left[L any, R any](value L) Either[L, R] { _ = "STUB: not implemented"; return nil }

// Right builds the right side of the Either struct, as opposed to the Left side.
func Right[L any, R any](value R) Either[L, R] { _ = "STUB: not implemented"; return nil }

// Either respresents a value of 2 possible types.
// An instance of Either is an instance of either A or B.
type Either[L any, R any] struct {
	isLeft bool

	left  L
	right R
}

// IsLeft returns true if Either is an instance of Left.
func (e Either[L, R]) IsLeft() bool {
	_ = "STUB: not implemented"

	// IsRight returns true if Either is an instance of Right.
	return false
}

func (e Either[L, R]) IsRight() bool {
	_ = "STUB: not implemented"

	// Left returns left value of a Either struct.
	return false
}

func (e Either[L, R]) Left() (L, bool) { _ = "STUB: not implemented"; return *new(L), false }

// Right returns right value of a Either struct.
func (e Either[L, R]) Right() (R, bool) { _ = "STUB: not implemented"; return *new(R), false }

// MustLeft returns left value of a Either struct or panics.
func (e Either[L, R]) MustLeft() L { _ = "STUB: not implemented"; return *new(L) }

// MustRight returns right value of a Either struct or panics.
func (e Either[L, R]) MustRight() R { _ = "STUB: not implemented"; return *new(R) }

// Unpack returns all values
func (e Either[L, R]) Unpack() (L, R) {
	_ = "STUB: not implemented"
	return *

	// LeftOrElse returns left value of a Either struct or fallback.
	new(L), *new(R)
}

func (e Either[L, R]) LeftOrElse(fallback L) L { _ = "STUB: not implemented"; return *new(L) }

// RightOrElse returns right value of a Either struct or fallback.
func (e Either[L, R]) RightOrElse(fallback R) R { _ = "STUB: not implemented"; return *new(R) }

// LeftOrEmpty returns left value of a Either struct or empty value.
func (e Either[L, R]) LeftOrEmpty() L { _ = "STUB: not implemented"; return *new(L) }

// RightOrEmpty returns right value of a Either struct or empty value.
func (e Either[L, R]) RightOrEmpty() R { _ = "STUB: not implemented"; return *new(R) }

// Swap returns the left value in Right and vice versa.
func (e Either[L, R]) Swap() Either[R, L] { _ = "STUB: not implemented"; return nil }

// ForEach executes the given side-effecting function, depending of value is Left or Right.
func (e Either[L, R]) ForEach(leftCb func(L), rightCb func(R)) { _ = "STUB: not implemented"; return }

// Match executes the given function, depending of value is Left or Right, and returns result.
func (e Either[L, R]) Match(onLeft func(L) Either[L, R], onRight func(R) Either[L, R]) Either[L, R] {
	_ = "STUB: not implemented"
	return nil
}

// MapLeft executes the given function, if Either is of type Left, and returns result.
func (e Either[L, R]) MapLeft(mapper func(L) Either[L, R]) Either[L, R] {
	_ = "STUB: not implemented"
	return nil
}

// MapRight executes the given function, if Either is of type Right, and returns result.
func (e Either[L, R]) MapRight(mapper func(R) Either[L, R]) Either[L, R] {
	_ = "STUB: not implemented"
	return nil
}

// leftValue returns left value of a Either struct.(implementation of Foldable interface)
//
//nolint:unused
func (e Either[L, R]) leftValue() L {
	_ = "STUB: not implemented"

	// rightValue returns right value of a Either struct.(implementation of Foldable interface)
	//
	//nolint:unused
	return *new(L)
}

func (e Either[L, R]) rightValue() R {
	_ = "STUB: not implemented"

	// hasLeft returns true if the Result represents an error state.
	//
	//nolint:unused
	return *new(R)
}

func (e Either[L, R]) hasLeftValue() bool { _ = "STUB: not implemented"; return false }
