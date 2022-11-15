package vec2d

import (
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	v1, v2 := New(10, 10), New(90, 90)
	v3 := v1.Add(v2)

	if v3.X != 100 || v3.Y != 100 {
		t.Fail()
	}
}

func TestSub(t *testing.T) {
	t.Parallel()

	v1, v2 := New(100, 100), New(90, 90)
	v3 := v1.Sub(v2)

	if v3.X != 10 || v3.Y != 10 {
		t.Fail()
	}
}

func TestMul(t *testing.T) {
	t.Parallel()

	v1, v2 := New(2, 50), New(50, 2)
	v3 := v1.Mul(v2)

	if v3.X != 100 || v3.Y != 100 {
		t.Fail()
	}
}

func TestMulScalar(t *testing.T) {
	t.Parallel()

	v1 := New(50, 50).MulScalar(2)

	if v1.X != 100 || v1.Y != 100 {
		t.Fail()
	}
}

func TestDiv(t *testing.T) {
	t.Parallel()

	v1, v2 := New(100, 100), New(50, 2)
	v3 := v1.Div(v2)

	if v3.X != 2 || v3.Y != 50 {
		t.Fail()
	}
}

func TestDivScalar(t *testing.T) {
	t.Parallel()

	v1 := New(100, 100).DivScalar(2)

	if v1.X != 50 || v1.Y != 50 {
		t.Fail()
	}
}

func TestMin(t *testing.T) {
	t.Parallel()

	v1, v2 := New(100, 1), New(1, 100)
	v3 := v1.Min(v2)

	if v3.X != 1 || v3.Y != 1 {
		t.Fail()
	}
}

func TestMax(t *testing.T) {
	t.Parallel()

	v1, v2 := New(100, 1), New(1, 100)
	v3 := v1.Max(v2)

	if v3.X != 100 || v3.Y != 100 {
		t.Fail()
	}
}

func TestNeg(t *testing.T) {
	t.Parallel()

	v1 := New(1, -1).Neg()

	if v1.X != -1 || v1.Y != 1 {
		t.Fail()
	}
}

func TestAbs(t *testing.T) {
	t.Parallel()

	v1 := New(-1, 1).Abs()

	if v1.X != 1 || v1.Y != 1 {
		t.Fail()
	}
}

func TestFloor(t *testing.T) {
	t.Parallel()

	v1 := New(2.3, 4.6).Floor()

	if v1.X > 2.0 || v1.Y > 4.0 {
		t.Fail()
	}
}

func TestCeil(t *testing.T) {
	t.Parallel()

	v1 := New(2.3, 4.6).Ceil()

	if v1.X < 3.0 || v1.Y < 5.0 {
		t.Fail()
	}
}

func TestCross(t *testing.T) {
	t.Parallel()

	v1, v2 := New(10, 5), New(5, 5)

	if v1.Cross(v2) != 25 {
		t.Fail()
	}
}

func TestDot(t *testing.T) {
	t.Parallel()

	v1, v2 := New(10, 5), New(5, 5)

	if v1.Dot(v2) != 75 {
		t.Fail()
	}
}

func TestLen(t *testing.T) {
	t.Parallel()

	v1 := New(10, 5)

	if v1.Len() != 11 {
		t.Fail()
	}
}

func TestNorm(t *testing.T) {
	t.Parallel()

	v1 := New(0.0, 10.0).Norm()

	if v1.X > 0.0 || v1.Y < 1.0 {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	t.Parallel()

	v1 := New(10, 5)
	v := v1.String()

	if !strings.ContainsAny(v, "015") {
		t.Fail()
	}
}
