package mo

func NewState[S any, A any](f func(state S) (A, S)) State[S, A] {
	_ = "STUB: not implemented"
	return nil
}

func ReturnState[S any, A any](x A) State[S, A] { _ = "STUB: not implemented"; return nil }

// State represents a function `(S) -> (A, S)`, where `S` is state, `A` is result.
type State[S any, A any] struct {
	run func(state S) (A, S)
}

// Run executes a computation in the State monad.
func (s State[S, A]) Run(state S) (A, S) {
	_ = "STUB: not implemented"
	return *

	// Get returns the current state.
	new(A), *new(S)
}

func (s State[S, A]) Get() State[S, S] { _ = "STUB: not implemented"; return nil }

// Modify the state by applying a function to the current state.
func (s State[S, A]) Modify(f func(state S) S) State[S, A] { _ = "STUB: not implemented"; return nil }

// Put set the state.
func (s State[S, A]) Put(state S) State[S, A] { _ = "STUB: not implemented"; return nil }
