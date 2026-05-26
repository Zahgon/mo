package either

import (
	"testing"

	"github.com/samber/mo"
)

func Pipe1[A1 any, A2 any, B1 any, B2 any](
	source mo.Either[A1, A2],
	operator1 func(mo.Either[A1, A2]) mo.Either[B1, B2],
) mo.Either[B1, B2] {
	_ = "STUB: not implemented"
	return nil
}

func Pipe2[A1 any, A2 any, B1 any, B2 any, C1 any, C2 any](
	source mo.Either[A1, A2],
	operator1 func(mo.Either[A1, A2]) mo.Either[B1, B2],
	operator2 func(mo.Either[B1, B2]) mo.Either[C1, C2],
) mo.Either[C1, C2] {
	_ = "STUB: not implemented"
	return nil
}

func Pipe3[A1 any, A2 any, B1 any, B2 any, C1 any, C2 any, D1 any, D2 any](
	source mo.Either[A1, A2],
	operator1 func(mo.Either[A1, A2]) mo.Either[B1, B2],
	operator2 func(mo.Either[B1, B2]) mo.Either[C1, C2],
	operator3 func(mo.Either[C1, C2]) mo.Either[D1, D2],
) mo.Either[D1, D2] {
	_ = "STUB: not implemented"
	return nil
}

func Pipe4[A1 any, A2 any, B1 any, B2 any, C1 any, C2 any, D1 any, D2 any, E1 any, E2 any](
	source mo.Either[A1, A2],
	operator1 func(mo.Either[A1, A2]) mo.Either[B1, B2],
	operator2 func(mo.Either[B1, B2]) mo.Either[C1, C2],
	operator3 func(mo.Either[C1, C2]) mo.Either[D1, D2],
	operator4 func(mo.Either[D1, D2]) mo.Either[E1, E2],
) mo.Either[E1, E2] {
	_ = "STUB: not implemented"
	return nil
}

func Pipe5[A1 any, A2 any, B1 any, B2 any, C1 any, C2 any, D1 any, D2 any, E1 any, E2 any, F1 any, F2 any](
	source mo.Either[A1, A2],
	operator1 func(mo.Either[A1, A2]) mo.Either[B1, B2],
	operator2 func(mo.Either[B1, B2]) mo.Either[C1, C2],
	operator3 func(mo.Either[C1, C2]) mo.Either[D1, D2],
	operator4 func(mo.Either[D1, D2]) mo.Either[E1, E2],
	operator5 func(mo.Either[E1, E2]) mo.Either[F1, F2],
) mo.Either[F1, F2] {
	_ = "STUB: not implemented"
	return nil
}

func Pipe6[A1 any, A2 any, B1 any, B2 any, C1 any, C2 any, D1 any, D2 any, E1 any, E2 any, F1 any, F2 any, G1 any, G2 any](
	source mo.Either[A1, A2],
	operator1 func(mo.Either[A1, A2]) mo.Either[B1, B2],
	operator2 func(mo.Either[B1, B2]) mo.Either[C1, C2],
	operator3 func(mo.Either[C1, C2]) mo.Either[D1, D2],
	operator4 func(mo.Either[D1, D2]) mo.Either[E1, E2],
	operator5 func(mo.Either[E1, E2]) mo.Either[F1, F2],
	operator6 func(mo.Either[F1, F2]) mo.Either[G1, G2],
) mo.Either[G1, G2] {
	_ = "STUB: not implemented"
	return nil
}

func Pipe7[A1 any, A2 any, B1 any, B2 any, C1 any, C2 any, D1 any, D2 any, E1 any, E2 any, F1 any, F2 any, G1 any, G2 any, H1 any, H2 any](
	source mo.Either[A1, A2],
	operator1 func(mo.Either[A1, A2]) mo.Either[B1, B2],
	operator2 func(mo.Either[B1, B2]) mo.Either[C1, C2],
	operator3 func(mo.Either[C1, C2]) mo.Either[D1, D2],
	operator4 func(mo.Either[D1, D2]) mo.Either[E1, E2],
	operator5 func(mo.Either[E1, E2]) mo.Either[F1, F2],
	operator6 func(mo.Either[F1, F2]) mo.Either[G1, G2],
	operator7 func(mo.Either[G1, G2]) mo.Either[H1, H2],
) mo.Either[H1, H2] {
	_ = "STUB: not implemented"
	return nil
}

func Pipe8[A1 any, A2 any, B1 any, B2 any, C1 any, C2 any, D1 any, D2 any, E1 any, E2 any, F1 any, F2 any, G1 any, G2 any, H1 any, H2 any, I1 any, I2 any](
	source mo.Either[A1, A2],
	operator1 func(mo.Either[A1, A2]) mo.Either[B1, B2],
	operator2 func(mo.Either[B1, B2]) mo.Either[C1, C2],
	operator3 func(mo.Either[C1, C2]) mo.Either[D1, D2],
	operator4 func(mo.Either[D1, D2]) mo.Either[E1, E2],
	operator5 func(mo.Either[E1, E2]) mo.Either[F1, F2],
	operator6 func(mo.Either[F1, F2]) mo.Either[G1, G2],
	operator7 func(mo.Either[G1, G2]) mo.Either[H1, H2],
	operator8 func(mo.Either[H1, H2]) mo.Either[I1, I2],
) mo.Either[I1, I2] {
	_ = "STUB: not implemented"
	return nil
}

func Pipe9[A1 any, A2 any, B1 any, B2 any, C1 any, C2 any, D1 any, D2 any, E1 any, E2 any, F1 any, F2 any, G1 any, G2 any, H1 any, H2 any, I1 any, I2 any, J1 any, J2 any](
	source mo.Either[A1, A2],
	operator1 func(mo.Either[A1, A2]) mo.Either[B1, B2],
	operator2 func(mo.Either[B1, B2]) mo.Either[C1, C2],
	operator3 func(mo.Either[C1, C2]) mo.Either[D1, D2],
	operator4 func(mo.Either[D1, D2]) mo.Either[E1, E2],
	operator5 func(mo.Either[E1, E2]) mo.Either[F1, F2],
	operator6 func(mo.Either[F1, F2]) mo.Either[G1, G2],
	operator7 func(mo.Either[G1, G2]) mo.Either[H1, H2],
	operator8 func(mo.Either[H1, H2]) mo.Either[I1, I2],
	operator9 func(mo.Either[I1, I2]) mo.Either[J1, J2],
) mo.Either[J1, J2] {
	_ = "STUB: not implemented"
	return nil
}

func Pipe10[A1 any, A2 any, B1 any, B2 any, C1 any, C2 any, D1 any, D2 any, E1 any, E2 any, F1 any, F2 any, G1 any, G2 any, H1 any, H2 any, I1 any, I2 any, J1 any, J2 any, K1 any, K2 any](
	source mo.Either[A1, A2],
	operator1 func(mo.Either[A1, A2]) mo.Either[B1, B2],
	operator2 func(mo.Either[B1, B2]) mo.Either[C1, C2],
	operator3 func(mo.Either[C1, C2]) mo.Either[D1, D2],
	operator4 func(mo.Either[D1, D2]) mo.Either[E1, E2],
	operator5 func(mo.Either[E1, E2]) mo.Either[F1, F2],
	operator6 func(mo.Either[F1, F2]) mo.Either[G1, G2],
	operator7 func(mo.Either[G1, G2]) mo.Either[H1, H2],
	operator8 func(mo.Either[H1, H2]) mo.Either[I1, I2],
	operator9 func(mo.Either[I1, I2]) mo.Either[J1, J2],
	operator10 func(mo.Either[J1, J2]) mo.Either[K1, K2],
) mo.Either[K1, K2] {
	_ = "STUB: not implemented"
	return nil
}

func TestPipeTypeTransformations(t *testing.T) { _ = "STUB: not implemented"; return }
