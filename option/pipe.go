package option

import "github.com/samber/mo"

func Pipe1[A any, B any](
	source mo.Option[A],
	operator1 func(mo.Option[A]) mo.Option[B],
) mo.Option[B] {
	_ = "STUB: not implemented"
	return nil
}

func Pipe2[A any, B any, C any](
	source mo.Option[A],
	operator1 func(mo.Option[A]) mo.Option[B],
	operator2 func(mo.Option[B]) mo.Option[C],
) mo.Option[C] {
	_ = "STUB: not implemented"
	return nil
}

func Pipe3[A any, B any, C any, D any](
	source mo.Option[A],
	operator1 func(mo.Option[A]) mo.Option[B],
	operator2 func(mo.Option[B]) mo.Option[C],
	operator3 func(mo.Option[C]) mo.Option[D],
) mo.Option[D] {
	_ = "STUB: not implemented"
	return nil
}

func Pipe4[A any, B any, C any, D any, E any](
	source mo.Option[A],
	operator1 func(mo.Option[A]) mo.Option[B],
	operator2 func(mo.Option[B]) mo.Option[C],
	operator3 func(mo.Option[C]) mo.Option[D],
	operator4 func(mo.Option[D]) mo.Option[E],
) mo.Option[E] {
	_ = "STUB: not implemented"
	return nil
}

func Pipe5[A any, B any, C any, D any, E any, F any](
	source mo.Option[A],
	operator1 func(mo.Option[A]) mo.Option[B],
	operator2 func(mo.Option[B]) mo.Option[C],
	operator3 func(mo.Option[C]) mo.Option[D],
	operator4 func(mo.Option[D]) mo.Option[E],
	operator5 func(mo.Option[E]) mo.Option[F],
) mo.Option[F] {
	_ = "STUB: not implemented"
	return nil
}

func Pipe6[A any, B any, C any, D any, E any, F any, G any](
	source mo.Option[A],
	operator1 func(mo.Option[A]) mo.Option[B],
	operator2 func(mo.Option[B]) mo.Option[C],
	operator3 func(mo.Option[C]) mo.Option[D],
	operator4 func(mo.Option[D]) mo.Option[E],
	operator5 func(mo.Option[E]) mo.Option[F],
	operator6 func(mo.Option[F]) mo.Option[G],
) mo.Option[G] {
	_ = "STUB: not implemented"
	return nil
}

func Pipe7[A any, B any, C any, D any, E any, F any, G any, H any](
	source mo.Option[A],
	operator1 func(mo.Option[A]) mo.Option[B],
	operator2 func(mo.Option[B]) mo.Option[C],
	operator3 func(mo.Option[C]) mo.Option[D],
	operator4 func(mo.Option[D]) mo.Option[E],
	operator5 func(mo.Option[E]) mo.Option[F],
	operator6 func(mo.Option[F]) mo.Option[G],
	operator7 func(mo.Option[G]) mo.Option[H],
) mo.Option[H] {
	_ = "STUB: not implemented"
	return nil
}

func Pipe8[A any, B any, C any, D any, E any, F any, G any, H any, I any](
	source mo.Option[A],
	operator1 func(mo.Option[A]) mo.Option[B],
	operator2 func(mo.Option[B]) mo.Option[C],
	operator3 func(mo.Option[C]) mo.Option[D],
	operator4 func(mo.Option[D]) mo.Option[E],
	operator5 func(mo.Option[E]) mo.Option[F],
	operator6 func(mo.Option[F]) mo.Option[G],
	operator7 func(mo.Option[G]) mo.Option[H],
	operator8 func(mo.Option[H]) mo.Option[I],
) mo.Option[I] {
	_ = "STUB: not implemented"
	return nil
}

func Pipe9[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](
	source mo.Option[A],
	operator1 func(mo.Option[A]) mo.Option[B],
	operator2 func(mo.Option[B]) mo.Option[C],
	operator3 func(mo.Option[C]) mo.Option[D],
	operator4 func(mo.Option[D]) mo.Option[E],
	operator5 func(mo.Option[E]) mo.Option[F],
	operator6 func(mo.Option[F]) mo.Option[G],
	operator7 func(mo.Option[G]) mo.Option[H],
	operator8 func(mo.Option[H]) mo.Option[I],
	operator9 func(mo.Option[I]) mo.Option[J],
) mo.Option[J] {
	_ = "STUB: not implemented"
	return nil
}

func Pipe10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any](
	source mo.Option[A],
	operator1 func(mo.Option[A]) mo.Option[B],
	operator2 func(mo.Option[B]) mo.Option[C],
	operator3 func(mo.Option[C]) mo.Option[D],
	operator4 func(mo.Option[D]) mo.Option[E],
	operator5 func(mo.Option[E]) mo.Option[F],
	operator6 func(mo.Option[F]) mo.Option[G],
	operator7 func(mo.Option[G]) mo.Option[H],
	operator8 func(mo.Option[H]) mo.Option[I],
	operator9 func(mo.Option[I]) mo.Option[J],
	operator10 func(mo.Option[J]) mo.Option[K],
) mo.Option[K] {
	_ = "STUB: not implemented"
	return nil
}
