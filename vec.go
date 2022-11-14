package vec2d

import (
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
)

type numeric interface {
	constraints.Float | constraints.Integer
}

// V represents 2D vector.
type V[T numeric] struct {
	X, Y T
}

// New creates new vector with given coordinates.
func New[T numeric](x, y T) (r V[T]) {
	return V[T]{
		X: x,
		Y: y,
	}
}

// Add performs vector addition.
func (v V[T]) Add(n V[T]) (r V[T]) {
	return V[T]{
		X: v.X + n.X,
		Y: v.Y + n.Y,
	}
}

// Sub performs vector substraction.
func (v V[T]) Sub(n V[T]) (r V[T]) {
	return V[T]{
		X: v.X - n.X,
		Y: v.Y - n.Y,
	}
}

// Mul performs vector multiplication.
func (v V[T]) Mul(n V[T]) (r V[T]) {
	return V[T]{
		X: v.X * n.X,
		Y: v.Y * n.Y,
	}
}

// MulScalar performs vector * scalar multiplication.
func (v V[T]) MulScalar(n T) (r V[T]) {
	return V[T]{
		X: v.X * n,
		Y: v.Y * n,
	}
}

// Div performs vector division.
func (v V[T]) Div(n V[T]) (r V[T]) {
	return V[T]{
		X: v.X / n.X,
		Y: v.Y / n.Y,
	}
}

// DivScalar performs vector * scalar division.
func (v V[T]) DivScalar(n T) (r V[T]) {
	return V[T]{
		X: v.X / n,
		Y: v.Y / n,
	}
}

// Len calculates length of vector.
func (v V[T]) Len() (r T) {
	return T(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

// Norm calculates normal vector.
func (v V[T]) Norm() (r V[T]) {
	return v.DivScalar(T(v.Len()))
}

// Dot calculates dot-product.
func (v V[T]) Dot(n V[T]) (r T) {
	return v.X*n.X + v.Y*n.Y
}

// Cross calculates cross-product.
func (v V[T]) Cross(n V[T]) (r T) {
	return v.X*n.X - v.Y*n.Y
}

// Min returns minimum of two vectors.
func (v V[T]) Min(n V[T]) (r V[T]) {
	return V[T]{
		X: T(math.Min(float64(v.X), float64(n.X))),
		Y: T(math.Min(float64(v.Y), float64(n.Y))),
	}
}

// Max returns maximum of two vectors.
func (v V[T]) Max(n V[T]) (r V[T]) {
	return V[T]{
		X: T(math.Max(float64(v.X), float64(n.X))),
		Y: T(math.Max(float64(v.Y), float64(n.Y))),
	}
}

// Floor returns rounded-down vector.
func (v V[T]) Floor() (r V[T]) {
	return V[T]{
		X: T(math.Floor(float64(v.X))),
		Y: T(math.Floor(float64(v.Y))),
	}
}

// Ceil returns rounded-up vector.
func (v V[T]) Ceil() (r V[T]) {
	return V[T]{
		X: T(math.Ceil(float64(v.X))),
		Y: T(math.Ceil(float64(v.Y))),
	}
}

// Neg returns negative vector.
func (v V[T]) Neg() (r V[T]) {
	return V[T]{
		X: -v.X,
		Y: -v.Y,
	}
}

// Abs returns vector of absolute values.
func (v V[T]) Abs() (r V[T]) {
	return V[T]{
		X: T(math.Abs(float64(v.X))),
		Y: T(math.Abs(float64(v.Y))),
	}
}

// String returns string representation of vector.
func (v V[T]) String() (rv string) {
	return fmt.Sprintf("vec(%v, %v)", v.X, v.Y)
}
