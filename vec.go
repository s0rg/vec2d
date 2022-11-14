package vec2d

import (
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
)

type numeric interface {
	constraints.Float | constraints.Integer
}

// V2D represents 2D vector.
type V2D[T numeric] struct {
	X, Y T
}

// New creates new vector with given coordinates.
func New[T numeric](x, y T) (r V2D[T]) {
	return V2D[T]{
		X: x,
		Y: y,
	}
}

// Add performs vector addition.
func (v V2D[T]) Add(n V2D[T]) (r V2D[T]) {
	return V2D[T]{
		X: v.X + n.X,
		Y: v.Y + n.Y,
	}
}

// Sub performs vector substraction.
func (v V2D[T]) Sub(n V2D[T]) (r V2D[T]) {
	return V2D[T]{
		X: v.X - n.X,
		Y: v.Y - n.Y,
	}
}

// Mul performs vector multiplication.
func (v V2D[T]) Mul(n V2D[T]) (r V2D[T]) {
	return V2D[T]{
		X: v.X * n.X,
		Y: v.Y * n.Y,
	}
}

// MulScalar performs vector * scalar multiplication.
func (v V2D[T]) MulScalar(n T) (r V2D[T]) {
	return V2D[T]{
		X: v.X * n,
		Y: v.Y * n,
	}
}

// Div performs vector division.
func (v V2D[T]) Div(n V2D[T]) (r V2D[T]) {
	return V2D[T]{
		X: v.X / n.X,
		Y: v.Y / n.Y,
	}
}

// DivScalar performs vector * scalar division.
func (v V2D[T]) DivScalar(n T) (r V2D[T]) {
	return V2D[T]{
		X: v.X / n,
		Y: v.Y / n,
	}
}

// Len calculates length of vector.
func (v V2D[T]) Len() (r T) {
	return T(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

// Norm calculates normal vector.
func (v V2D[T]) Norm() (r V2D[T]) {
	return v.DivScalar(T(v.Len()))
}

// Dot calculates dot-product.
func (v V2D[T]) Dot(n V2D[T]) (r T) {
	return v.X*n.X + v.Y*n.Y
}

// Cross calculates cross-product.
func (v V2D[T]) Cross(n V2D[T]) (r T) {
	return v.X*n.X - v.Y*n.Y
}

// Min returns minimum of two vectors.
func (v V2D[T]) Min(n V2D[T]) (r V2D[T]) {
	return V2D[T]{
		X: T(math.Min(float64(v.X), float64(n.X))),
		Y: T(math.Min(float64(v.Y), float64(n.Y))),
	}
}

// Max returns maximum of two vectors.
func (v V2D[T]) Max(n V2D[T]) (r V2D[T]) {
	return V2D[T]{
		X: T(math.Max(float64(v.X), float64(n.X))),
		Y: T(math.Max(float64(v.Y), float64(n.Y))),
	}
}

// Floor returns rounded-down vector.
func (v V2D[T]) Floor() (r V2D[T]) {
	return V2D[T]{
		X: T(math.Floor(float64(v.X))),
		Y: T(math.Floor(float64(v.Y))),
	}
}

// Ceil returns rounded-up vector.
func (v V2D[T]) Ceil() (r V2D[T]) {
	return V2D[T]{
		X: T(math.Ceil(float64(v.X))),
		Y: T(math.Ceil(float64(v.Y))),
	}
}

// Neg returns negative vector.
func (v V2D[T]) Neg() (r V2D[T]) {
	return V2D[T]{
		X: -v.X,
		Y: -v.Y,
	}
}

// Abs returns vector of absolute values.
func (v V2D[T]) Abs() (r V2D[T]) {
	return V2D[T]{
		X: T(math.Abs(float64(v.X))),
		Y: T(math.Abs(float64(v.Y))),
	}
}

// String returns string representation of vector.
func (v V2D[T]) String() (rv string) {
	return fmt.Sprintf("vec(%v, %v)", v.X, v.Y)
}
