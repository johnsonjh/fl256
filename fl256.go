// Package fl256 implements a fixed-point, '63' decimal precision type.
// The maximum integer value of the "fl256" type is '4,398,046,511,103'.
// It implements a serialisation codec for compatibility with big.Float
// that uses a "42.214" fixed-precision value. During calculations, it is
// possible to work with negative numbers for convenience, however, they
// cannot be imported or encoded in "fl256" representation. The functions
// return 'nil' to signify an error condition or when given invalid input.

package fl256 // import "go.gridfinity.dev/fl256"

import (
	"math/big"
)

// Zero returns a new float, with '256' bits of precision.
// This is a particularly a good way to generate a new variable
// with minimal typing for input conversion functions.
func Zero() *big.Float {
	r := big.NewFloat(0.0)
	r.SetPrec(256)
	return r
}

// New returns a new floating point number with '256' bits
// of precision, containing the desired initial value.
func New(f float64) *big.Float {
	r := big.NewFloat(f)
	r.SetPrec(256)
	return r
}

// FromString accepts a properly formed string of integers or decimal
// values and returns a big.Float to the caller. Negative values, and
// values greater than '4,398,046,511,103' are invalid inputs and will
// be rejected with the function returning 'nil'.
func FromString(s string) *big.Float {
	r, success := Zero().SetString(s)
	if r.Sign() < 0 {
		return nil
	}
	if r.MantExp(nil) > 42 {
		return nil
	}
	if success {
		return r
	}
	return nil
}

// Int returns the truncated value of the parameter as a *big.Int.
func Int(a *big.Float) *big.Int {
	i, _ := a.Int(nil)
	return i
}

// FromInt returns a *big.Float from a *big.Int. Negative values and
// values greater than '4,398,046,511,103' are invalid inputs and will
// be rejected by the function, returning 'nil'.
func FromInt(i *big.Int) *big.Float {
	r := Zero().SetInt(i)
	if r.Sign() < 0 {
		return nil
	}
	if r.MantExp(nil) > 42 {
		return nil
	}
	return r
}

// Int64 returns the value passed as int64, with the decimal truncated.
func Int64(a *big.Float) int64 {
	r := Zero()
	r.Copy(a)
	i, _ := r.Int64()
	return i
}

// FromInt64 returns a *big.Float from an int64. Negative values and
// values greater than '4,398,046,511,103' are invalid inputs, and will
// be rejected by the function, returning 'nil'.
func FromInt64(i int64) *big.Float {
	r := Zero().SetInt64(i)
	if r.Sign() < 0 {
		return nil
	}
	if r.MantExp(nil) > 42 {
		return nil
	}
	return r
}

// Uint64 returns the value passed as a uint64, with the decimal truncated.
func Uint64(a *big.Float) uint64 {
	r := Zero()
	r.Copy(a)
	i, _ := r.Uint64()
	return i
}

// FromUint64 returns a *big.Float from a uint64. Values greater than
// '4,398,046,511,103' are invalid inputs, and will and rejected by the
// functions, returning 'nil', as they exceed the precision of "fl256".
func FromUint64(i uint64) *big.Float {
	r := Zero().SetUint64(i)
	if r.Sign() < 0 {
		return nil
	}
	if r.MantExp(nil) > 42 {
		return nil
	}
	return r
}

// Equal returns true, if 'x' and 'y' are equal.
func Equal(x, y *big.Int) bool {
	return x.Cmp(y) == 0
}

// Greater returns true, if ('x' > 'y')
func Greater(x, y *big.Float) bool {
	return x.Cmp(y) == 1
}

// Lesser returns true, if ('x' < 'y')
func Lesser(x, y *big.Float) bool {
	return x.Cmp(y) == -1
}

// Exp returns ('a' ** 'b')
func Exp(a *big.Float, e uint64) *big.Float {
	result := Zero().Copy(a)
	for i := uint64(0); i < e-1; i++ {
		result = Mul(result, a)
	}
	return result
}

// Root returns the 'n'th root of 'a' to '255' significant bits.
// XXX(jhj): If it is larger, the function gets stuck in an infinite loop.
func Root(a *big.Float, n uint64) *big.Float {
	limit := Exp(New(2), 255)
	n1 := n - 1
	n1f, rn := New(float64(n1)), Div(New(1.0), New(float64(n)))
	x, x0 := New(1.0), Zero()
	_ = x0
	for {
		potx, t2 := Div(New(1.0), x), a
		for b := n1; b > 0; b >>= 1 {
			if b&1 == 1 {
				t2 = Mul(t2, potx)
			}
			potx = Mul(potx, potx)
		}
		x0, x = x, Mul(rn, Add(Mul(n1f, x), t2))
		if Lesser(Mul(Abs(Sub(x, x0)), limit), x) {
			break
		}
	}
	return x
}

// Abs returns 'a', with it's sign set to positive.
func Abs(a *big.Float) *big.Float {
	return Zero().Abs(a)
}

// Neg flips the sign of the parameter Float.
// Negative values are rejected by the Encode function.
// (They are invalid in the context of a crypto-currency ledger.)
func Neg(a *big.Float) *big.Float {
	return a.Neg(a)
}

// Sign returns '1' for positive,
// '-1' for negative, and '0' for zeros.
func Sign(a *big.Float) int {
	return a.Sign()
}

// Sqrt returns the square root of the given parameter.
func Sqrt(a *big.Float) *big.Float {
	return Root(a, 2)
}

// Add returns ('a' + 'b')
func Add(a, b *big.Float) *big.Float {
	return Zero().Add(a, b)
}

// Sub returns ('a' - 'b')
func Sub(a, b *big.Float) *big.Float {
	return Zero().Sub(a, b)
}

// Mul returns ('a' * 'b')
func Mul(a, b *big.Float) *big.Float {
	return Zero().Mul(a, b)
}

// Div returns ('a' / 'b')
func Div(a, b *big.Float) *big.Float {
	return Zero().Quo(a, b)
}

// Mod returns ('a' % 'b')
func Mod(a, b *big.Float) *big.Float {
	q := Div(a, b)
	i := Int(q)
	fi := Zero().SetInt(i)
	rem := Sub(a, Mul(b, fi))
	return rem
}

// Encode takes a big.Float input, bit-shifts left, and truncates the
// result at '214' bits of decimal precision, or a maximum of '42' bits
// of integer precision, returning a byte slice. Negative values, and
// values greater than '4,398,046,511,103' are invalid inputs, and will
// be rejected by the function, returning 'nil'.
func Encode(a *big.Float) []byte {
	if a.Sign() < 0 {
		return nil
	}
	mantissa := New(0)
	a.MantExp(mantissa)
	exp := a.MantExp(nil)
	if exp > 42 {
		return nil
	}
	bytes := Int(Mul(Exp(New(2), uint64(214+exp)), mantissa)).Bytes()
	return append(make([]byte, 32-len(bytes)), bytes...)
}

// Decode takes a value (created by the Encode function), loads it into
// a big.Float, and changes the exponent, shifting the point to '214'
// bits of decimal precision and '42' bits of integer precision (exponent).
func Decode(b []byte) *big.Float {
	decoded := big.NewInt(0).SetBytes(b)
	decodedF := Zero().SetInt(decoded)
	return Zero().SetMantExp(decodedF, -214)
}
