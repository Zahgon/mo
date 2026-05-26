package mo

import (
	"database/sql/driver"
	"fmt"
)

var errOptionNoSuchElement = fmt.Errorf("no such element")

type zeroer interface {
	IsZero() bool
}

// Some builds an Option when value is present.
// Play: https://go.dev/play/p/iqz2n9n0tDM
func Some[T any](value T) Option[T] { _ = "STUB: not implemented"; return nil }

// None builds an Option when value is absent.
// Play: https://go.dev/play/p/yYQPsYCSYlD
func None[T any]() Option[T] { _ = "STUB: not implemented"; return nil }

// TupleToOption builds a Some Option when second argument is true, or None.
// Play: https://go.dev/play/p/gkrg2pZwOty
func TupleToOption[T any](value T, ok bool) Option[T] { _ = "STUB: not implemented"; return nil }

// EmptyableToOption builds a Some Option when value is not empty, or None.
// Play: https://go.dev/play/p/GSpQQ-q-UES
func EmptyableToOption[T any](value T) Option[T] {
	_ = "STUB: not implemented"
	// 🤮
	return nil
}

// PointerToOption builds a Some Option when value is not nil, or None.
// Play: https://go.dev/play/p/yPVMj4DUb-I
func PointerToOption[T any](value *T) Option[T] { _ = "STUB: not implemented"; return nil }

// Option is a container for an optional value of type T. If value exists, Option is
// of type Some. If the value is absent, Option is of type None.
type Option[T any] struct {
	isPresent bool
	value     T
}

// IsPresent returns false when value is absent.
// Play: https://go.dev/play/p/nDqIaiihyCA
func (o Option[T]) IsPresent() bool {
	_ = "STUB: not implemented"

	// IsSome is an alias to IsPresent.
	// Play: https://go.dev/play/p/DyvGRy7fP9m
	return false
}

func (o Option[T]) IsSome() bool { _ = "STUB: not implemented"; return false }

// IsAbsent returns false when value is present.
// Play: https://go.dev/play/p/23e2zqyVOQm
func (o Option[T]) IsAbsent() bool { _ = "STUB: not implemented"; return false }

// IsNone is an alias to IsAbsent.
// Play: https://go.dev/play/p/EdqxKhborIP
func (o Option[T]) IsNone() bool { _ = "STUB: not implemented"; return false }

// Size returns 1 when value is present or 0 instead.
// Play: https://go.dev/play/p/7ixCNG1E9l7
func (o Option[T]) Size() int { _ = "STUB: not implemented"; return 0 }

// Get returns value and presence.
// Play: https://go.dev/play/p/0-JBa1usZRT
func (o Option[T]) Get() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

// MustGet returns value if present or panics instead.
// Play: https://go.dev/play/p/RVBckjdi5WR
func (o Option[T]) MustGet() T { _ = "STUB: not implemented"; return *new(T) }

// OrElse returns value if present or default value.
// Play: https://go.dev/play/p/TrGByFWCzXS
func (o Option[T]) OrElse(fallback T) T { _ = "STUB: not implemented"; return *new(T) }

// OrEmpty returns value if present or empty value.
// Play: https://go.dev/play/p/SpSUJcE-tQm
func (o Option[T]) OrEmpty() T {
	_ = "STUB: not implemented"

	// ForEach executes the given side-effecting function of value is present.
	return *new(T)
}

func (o Option[T]) ForEach(onValue func(value T)) { _ = "STUB: not implemented"; return }

// Match executes the first function if value is present and second function if absent.
// It returns a new Option.
// Play: https://go.dev/play/p/1V6st3LDJsM
func (o Option[T]) Match(onValue func(value T) (T, bool), onNone func() (T, bool)) Option[T] {
	_ = "STUB: not implemented"
	return nil
}

// Map executes the mapper function if value is present or returns None if absent.
// Play: https://go.dev/play/p/mvfP3pcP_eJ
func (o Option[T]) Map(mapper func(value T) (T, bool)) Option[T] {
	_ = "STUB: not implemented"
	return nil
}

// MapNone executes the mapper function if value is absent or returns Option.
// Play: https://go.dev/play/p/_KaHWZ6Q17b
func (o Option[T]) MapNone(mapper func() (T, bool)) Option[T] {
	_ = "STUB: not implemented"
	return nil
}

// FlatMap executes the mapper function if value is present or returns None if absent.
// Play: https://go.dev/play/p/OXO-zJx6n5r
func (o Option[T]) FlatMap(mapper func(value T) Option[T]) Option[T] {
	_ = "STUB: not implemented"
	return nil
}

// MapValue executes the mapper function if value is present or returns None if absent.
func (o Option[T]) MapValue(mapper func(value T) T) Option[T] {
	_ = "STUB: not implemented"
	return nil
}

// ToPointer returns value if present or a nil pointer.
// Play: https://go.dev/play/p/KJc2Pv3KNPW
func (o Option[T]) ToPointer() *T { _ = "STUB: not implemented"; return nil }

// MarshalJSON encodes Option into json.
// Go 1.20+ relies on the IsZero method when the `omitempty` tag is used
// unless a custom MarshalJSON method is defined.  Then the IsZero method is ignored.
// current best workaround is to instead use `omitzero` tag with Go 1.24+
func (o Option[T]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON decodes Option from json.
func (o *Option[T]) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	// reset the value if not set later.
	return nil
}

// If user manually set the field to be `null`, then it either means the option is absent or present with a zero value.

// // If the type is a pointer, then it means the option is present with a zero value.
// o.isPresent = reflect.TypeOf(o.value).Kind() == reflect.Ptr
// return nil

// IsZero assists `omitzero` tag introduced in Go 1.24
func (o Option[T]) IsZero() bool { _ = "STUB: not implemented"; return false }

// MarshalText implements the encoding.TextMarshaler interface.
func (o Option[T]) MarshalText() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// UnmarshalText implements the encoding.TextUnmarshaler interface.
		nil
}

func (o *Option[T]) UnmarshalText(data []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalBinary is the interface implemented by an object that can marshal itself into a binary form.
func (o Option[T]) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalBinary is the interface implemented by an object that can unmarshal a binary representation of itself.
func (o *Option[T]) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

// GobEncode implements the gob.GobEncoder interface.
func (o Option[T]) GobEncode() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// GobDecode implements the gob.GobDecoder interface.
		nil
}

func (o *Option[T]) GobDecode(data []byte) error { _ = "STUB: not implemented"; return nil }

// Scan implements the SQL sql.Scanner interface.
func (o *Option[T]) Scan(src any) error { _ = "STUB: not implemented"; return nil }

// is is only possible to assert interfaces, so convert first
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#why-not-permit-type-assertions-on-values-whose-type-is-a-type-parameter

// Value implements the driver Valuer interface.
func (o Option[T]) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// Equal compares two Option[T] instances for equality
func (o Option[T]) Equal(other Option[T]) bool { _ = "STUB: not implemented"; return false }

// leftValue returns an error if the Option is None, otherwise nil
//
//nolint:unused
func (o Option[T]) leftValue() error { _ = "STUB: not implemented"; return nil }

// rightValue returns the value if the Option is Some, otherwise the zero value of T
//
//nolint:unused
func (o Option[T]) rightValue() T { _ = "STUB: not implemented"; return *new(T) }

// hasLeftValue returns true if the Option represents a None state
//
//nolint:unused
func (o Option[T]) hasLeftValue() bool { _ = "STUB: not implemented"; return false }
