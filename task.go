package mo

// NewTask instanciates a new Task.
func NewTask[R any](f ff0[R]) Task[R] { _ = "STUB: not implemented"; return nil }

// NewTaskFromIO instanciates a new Task from an existing IO.
func NewTaskFromIO[R any](io IO[R]) Task[R] { _ = "STUB: not implemented"; return nil }

// Task represents a non-deterministic asynchronous computation that
// can cause side effects, yields a value of type `R` and never fails.
type Task[R any] struct {
	unsafePerform ff0[R]
}

// Run execute the non-deterministic asynchronous computation, with side effect.
func (t Task[R]) Run() *Future[R] { _ = "STUB: not implemented"; return nil }

// NewTask1 instanciates a new Task1.
func NewTask1[R any, A any](f ff1[R, A]) Task1[R, A] { _ = "STUB: not implemented"; return nil }

// NewTaskFromIO1 instanciates a new Task1 from an existing IO1.
func NewTaskFromIO1[R any, A any](io IO1[R, A]) Task1[R, A] { _ = "STUB: not implemented"; return nil }

// Task1 represents a non-deterministic asynchronous computation that
// can cause side effects, yields a value of type `R` and never fails.
type Task1[R any, A any] struct {
	unsafePerform ff1[R, A]
}

// Run execute the non-deterministic asynchronous computation, with side effect.
func (t Task1[R, A]) Run(a A) *Future[R] { _ = "STUB: not implemented"; return nil }

// NewTask2 instanciates a new Task2.
func NewTask2[R any, A any, B any](f ff2[R, A, B]) Task2[R, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// NewTaskFromIO2 instanciates a new Task2 from an existing IO2.
func NewTaskFromIO2[R any, A any, B any](io IO2[R, A, B]) Task2[R, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// Task2 represents a non-deterministic asynchronous computation that
// can cause side effects, yields a value of type `R` and never fails.
type Task2[R any, A any, B any] struct {
	unsafePerform ff2[R, A, B]
}

// Run execute the non-deterministic asynchronous computation, with side effect.
func (t Task2[R, A, B]) Run(a A, b B) *Future[R] { _ = "STUB: not implemented"; return nil }

// NewTask3 instanciates a new Task3.
func NewTask3[R any, A any, B any, C any](f ff3[R, A, B, C]) Task3[R, A, B, C] {
	_ = "STUB: not implemented"
	return nil
}

// NewTaskFromIO3 instanciates a new Task3 from an existing IO3.
func NewTaskFromIO3[R any, A any, B any, C any](io IO3[R, A, B, C]) Task3[R, A, B, C] {
	_ = "STUB: not implemented"
	return nil
}

// Task3 represents a non-deterministic asynchronous computation that
// can cause side effects, yields a value of type `R` and never fails.
type Task3[R any, A any, B any, C any] struct {
	unsafePerform ff3[R, A, B, C]
}

// Run execute the non-deterministic asynchronous computation, with side effect.
func (t Task3[R, A, B, C]) Run(a A, b B, c C) *Future[R] { _ = "STUB: not implemented"; return nil }

// NewTask4 instanciates a new Task4.
func NewTask4[R any, A any, B any, C any, D any](f ff4[R, A, B, C, D]) Task4[R, A, B, C, D] {
	_ = "STUB: not implemented"
	return nil
}

// NewTaskFromIO4 instanciates a new Task4 from an existing IO4.
func NewTaskFromIO4[R any, A any, B any, C any, D any](io IO4[R, A, B, C, D]) Task4[R, A, B, C, D] {
	_ = "STUB: not implemented"
	return nil
}

// Task4 represents a non-deterministic asynchronous computation that
// can cause side effects, yields a value of type `R` and never fails.
type Task4[R any, A any, B any, C any, D any] struct {
	unsafePerform ff4[R, A, B, C, D]
}

// Run execute the non-deterministic asynchronous computation, with side effect.
func (t Task4[R, A, B, C, D]) Run(a A, b B, c C, d D) *Future[R] {
	_ = "STUB: not implemented"
	return nil
}

// NewTask5 instanciates a new Task5.
func NewTask5[R any, A any, B any, C any, D any, E any](f ff5[R, A, B, C, D, E]) Task5[R, A, B, C, D, E] {
	_ = "STUB: not implemented"
	return nil
}

// NewTaskFromIO5 instanciates a new Task5 from an existing IO5.
func NewTaskFromIO5[R any, A any, B any, C any, D any, E any](io IO5[R, A, B, C, D, E]) Task5[R, A, B, C, D, E] {
	_ = "STUB: not implemented"
	return nil
}

// Task5 represents a non-deterministic asynchronous computation that
// can cause side effects, yields a value of type `R` and never fails.
type Task5[R any, A any, B any, C any, D any, E any] struct {
	unsafePerform ff5[R, A, B, C, D, E]
}

// Run execute the non-deterministic asynchronous computation, with side effect.
func (t Task5[R, A, B, C, D, E]) Run(a A, b B, c C, d D, e E) *Future[R] {
	_ = "STUB: not implemented"
	return nil
}
