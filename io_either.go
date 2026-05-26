package mo

// NewIOEither instanciates a new IO.
func NewIOEither[R any](f fe0[R]) IOEither[R] { _ = "STUB: not implemented"; return nil }

// IOEither represents a non-deterministic synchronous computation that
// can cause side effects, yields a value of type `R` and can fail.
type IOEither[R any] struct {
	unsafePerform fe0[R]
}

// Run execute the non-deterministic synchronous computation, with side effect.
func (io IOEither[R]) Run() Either[error, R] { _ = "STUB: not implemented"; return nil }

// NewIOEither1 instanciates a new IO1.
func NewIOEither1[R any, A any](f fe1[R, A]) IOEither1[R, A] { _ = "STUB: not implemented"; return nil }

// IOEither1 represents a non-deterministic synchronous computation that
// can cause side effects, yields a value of type `R` and can fail.
type IOEither1[R any, A any] struct {
	unsafePerform fe1[R, A]
}

// Run execute the non-deterministic synchronous computation, with side effect.
func (io IOEither1[R, A]) Run(a A) Either[error, R] { _ = "STUB: not implemented"; return nil }

// NewIOEither2 instanciates a new IO2.
func NewIOEither2[R any, A any, B any](f fe2[R, A, B]) IOEither2[R, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// IOEither2 represents a non-deterministic synchronous computation that
// can cause side effects, yields a value of type `R` and can fail.
type IOEither2[R any, A any, B any] struct {
	unsafePerform fe2[R, A, B]
}

// Run execute the non-deterministic synchronous computation, with side effect.
func (io IOEither2[R, A, B]) Run(a A, b B) Either[error, R] { _ = "STUB: not implemented"; return nil }

// NewIOEither3 instanciates a new IO3.
func NewIOEither3[R any, A any, B any, C any](f fe3[R, A, B, C]) IOEither3[R, A, B, C] {
	_ = "STUB: not implemented"
	return nil
}

// IOEither3 represents a non-deterministic synchronous computation that
// can cause side effects, yields a value of type `R` and can fail.
type IOEither3[R any, A any, B any, C any] struct {
	unsafePerform fe3[R, A, B, C]
}

// Run execute the non-deterministic synchronous computation, with side effect.
func (io IOEither3[R, A, B, C]) Run(a A, b B, c C) Either[error, R] {
	_ = "STUB: not implemented"
	return nil
}

// NewIOEither4 instanciates a new IO4.
func NewIOEither4[R any, A any, B any, C any, D any](f fe4[R, A, B, C, D]) IOEither4[R, A, B, C, D] {
	_ = "STUB: not implemented"
	return nil
}

// IOEither4 represents a non-deterministic synchronous computation that
// can cause side effects, yields a value of type `R` and can fail.
type IOEither4[R any, A any, B any, C any, D any] struct {
	unsafePerform fe4[R, A, B, C, D]
}

// Run execute the non-deterministic synchronous computation, with side effect.
func (io IOEither4[R, A, B, C, D]) Run(a A, b B, c C, d D) Either[error, R] {
	_ = "STUB: not implemented"
	return nil
}

// NewIOEither5 instanciates a new IO5.
func NewIOEither5[R any, A any, B any, C any, D any, E any](f fe5[R, A, B, C, D, E]) IOEither5[R, A, B, C, D, E] {
	_ = "STUB: not implemented"
	return nil
}

// IOEither5 represents a non-deterministic synchronous computation that
// can cause side effects, yields a value of type `R` and can fail.
type IOEither5[R any, A any, B any, C any, D any, E any] struct {
	unsafePerform fe5[R, A, B, C, D, E]
}

// Run execute the non-deterministic synchronous computation, with side effect.
func (io IOEither5[R, A, B, C, D, E]) Run(a A, b B, c C, d D, e E) Either[error, R] {
	_ = "STUB: not implemented"
	return nil
}
