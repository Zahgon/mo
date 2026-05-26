package mo

// Ok builds a Result when value is valid.
// Play: https://go.dev/play/p/PDwADdzNoyZ
func Ok[T any](value T) Result[T] { _ = "STUB: not implemented"; return nil }

// Err builds a Result when value is invalid.
// Play: https://go.dev/play/p/PDwADdzNoyZ
func Err[T any](err error) Result[T] { _ = "STUB: not implemented"; return nil }

// Errf builds a Result when value is invalid.
// Errf formats according to a format specifier and returns the error as a value that satisfies Result[T].
// Play: https://go.dev/play/p/N43w92SM-Bs
func Errf[T any](format string, a ...any) Result[T] { _ = "STUB: not implemented"; return nil }

// TupleToResult convert a pair of T and error into a Result.
// Play: https://go.dev/play/p/KWjfqQDHQwa
func TupleToResult[T any](value T, err error) Result[T] { _ = "STUB: not implemented"; return nil }

// Try returns either a Ok or Err object.
// Play: https://go.dev/play/p/ilOlQx-Mx42
func Try[T any](f func() (T, error)) Result[T] { _ = "STUB: not implemented"; return nil }

// Result represents a result of an action having one
// of the following output: success or failure.
// An instance of Result is an instance of either Ok or Err.
// It could be compared to `Either[error, T]`.
type Result[T any] struct {
	isErr bool
	value T
	err   error
}

// IsOk returns true when value is valid.
// Play: https://go.dev/play/p/sfNvBQyZfgU
func (r Result[T]) IsOk() bool {
	_ = "STUB: not implemented"

	// IsError returns true when value is invalid.
	// Play: https://go.dev/play/p/xkV9d464scV
	return false
}

func (r Result[T]) IsError() bool {
	_ = "STUB: not implemented"

	// Error returns error when value is invalid or nil.
	// Play: https://go.dev/play/p/CSkHGTyiXJ5
	return false
}

func (r Result[T]) Error() error {
	_ = "STUB: not implemented"

	// Get returns value and error.
	// Play: https://go.dev/play/p/8KyX3z6TuNo
	return nil
}

func (r Result[T]) Get() (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

// MustGet returns value when Result is valid or panics.
// Play: https://go.dev/play/p/8LSlndHoTAE
func (r Result[T]) MustGet() T { _ = "STUB: not implemented"; return *new(T) }

// OrElse returns value when Result is valid or default value.
// Play: https://go.dev/play/p/MN_ULx0soi6
func (r Result[T]) OrElse(fallback T) T { _ = "STUB: not implemented"; return *new(T) }

// OrEmpty returns value when Result is valid or empty value.
// Play: https://go.dev/play/p/rdKtBmOcMLh
func (r Result[T]) OrEmpty() T {
	_ = "STUB: not implemented"

	// ToEither transforms a Result into an Either type.
	// Play: https://go.dev/play/p/Uw1Zz6b952q
	return *new(T)
}

func (r Result[T]) ToEither() Either[error, T] { _ = "STUB: not implemented"; return nil }

// ForEach executes the given side-effecting function if Result is valid.
func (r Result[T]) ForEach(mapper func(value T)) { _ = "STUB: not implemented"; return }

// Match executes the first function if Result is valid and second function if invalid.
// It returns a new Result.
// Play: https://go.dev/play/p/-_eFaLJ31co
func (r Result[T]) Match(onSuccess func(value T) (T, error), onError func(err error) (T, error)) Result[T] {
	_ = "STUB: not implemented"
	return nil
}

// Map executes the mapper function if Result is valid. It returns a new Result.
// Play: https://go.dev/play/p/-ndpN_b_OSc
func (r Result[T]) Map(mapper func(value T) (T, error)) Result[T] {
	_ = "STUB: not implemented"
	return nil
}

// MapValue executes the mapper function if Result is valid. It returns a new Result.
func (r Result[T]) MapValue(mapper func(value T) T) Result[T] {
	_ = "STUB: not implemented"
	return nil
}

// MapErr executes the mapper function if Result is invalid. It returns a new Result.
// Play: https://go.dev/play/p/WraZixg9GGf
func (r Result[T]) MapErr(mapper func(error) (T, error)) Result[T] {
	_ = "STUB: not implemented"
	return nil
}

// FlatMap executes the mapper function if Result is valid. It returns a new Result.
// Play: https://go.dev/play/p/Ud5QjZOqg-7
func (r Result[T]) FlatMap(mapper func(value T) Result[T]) Result[T] {
	_ = "STUB: not implemented"
	return nil
}

// MarshalJSON encodes Result into json, following the JSON-RPC specification for results,
// with one exception: when the result is an error, the "code" field is not included.
// Reference: https://www.jsonrpc.org/specification
func (o Result[T]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON decodes json into Result. If "error" is set, the result is an
// Err containing the error message as a generic error object. Otherwise, the
// result is an Ok containing the result. If the JSON object contains netiher
// an error nor a result, the result is an Ok containing an empty value. If the
// JSON object contains both an error and a result, the result is an Err. Finally,
// if the JSON object contains an error but is not structured correctly (no message
// field), the unmarshaling fails.
func (o *Result[T]) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// leftValue returns the error if the Result is an error, otherwise nil
//
//nolint:unused
func (r Result[T]) leftValue() error {
	_ = "STUB: not implemented"

	// rightValue returns the value if the Result is a success, otherwise the zero value of T
	//
	//nolint:unused
	return nil
}

func (r Result[T]) rightValue() T {
	_ = "STUB: not implemented"

	// hasLeftValue returns true if the Result represents an error state.
	//
	//nolint:unused
	return *new(T)
}

func (r Result[T]) hasLeftValue() bool { _ = "STUB: not implemented"; return false }
