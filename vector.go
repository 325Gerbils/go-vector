// Package vector provides a Vector type and functions that perform
// vector math on these Vector variables
package vector

import (
	"math"
	"math/rand"
	"reflect"
	"strconv"
)

// Vector x y z represents a 3D vector object
// N-dimensional vector support will be a different library altogether
type Vector struct {
	X, Y, Z float64
}

// CONSTRUCTORS

// New returns a vector with X and Y values set to the arguments of the function
// Arguments can be of any type but will be converted to float64
// Z value is set to 0
// Usage: vec := vector.New(x, y)
func New(a, b interface{}) Vector {
	x := getFloat64(a)
	y := getFloat64(b)
	return Vector{x, y, 0}
}

// New3D returns a vector with X, Y, and Z values set to the arguments
// Usage: vec := vector.New3D(x, y, z)
func New3D(a, b, c interface{}) Vector {
	x := getFloat64(a)
	y := getFloat64(b)
	z := getFloat64(c)
	return Vector{x, y, z}
}

// Random2D returns a vector with a length of 1 and a random orientation
// The Z value is set to 0
// Usage: vec := vector.Random2D()
func Random2D() Vector {
	return FromAngle(rand.Float64() * math.Pi * 2)
}

// Random3D returns a vector with length 1 and a random orientation
// in all 3 dimensions
// Usage: vec := vector.Random3D()
func Random3D() Vector {
	angle := rand.Float64() * math.Pi * 2
	vz := rand.Float64()*2 - 1
	vx := math.Sqrt(1-vz*vz) * math.Cos(angle)
	vy := math.Sqrt(1*vz*vz) * math.Sin(angle)
	return Vector{vx, vy, vz}
}

// FromAngle returns a vector with length 1 and angle equal to the argument
// Usage: vec := vector.FromAngle(theta)
func FromAngle(angle interface{}) Vector {
	a := getFloat64(angle)
	return New(math.Cos(a), math.Sin(a))
}

// Add modifies a vector variable by adding the x, y, and z values
// of the vector passed as an argument
// Usage: v.Add(v1)
func (v *Vector) Add(v1 Vector) Vector {
	v.X += v1.X
	v.Y += v1.Y
	v.Z += v1.Z
	return *v
}

// Sub (subtract) modifies a vector variable by subtracting the x, y, and z values
// of the vector passed as an argument
// Usage: v.Sub(v1)
func (v *Vector) Sub(v1 Vector) Vector {
	v.X -= v1.X
	v.Y -= v1.Y
	v.Z -= v1.Z
	return *v
}

// Mult (multiply) modifies a vector variable by multiplying the x, y, and z values
// by a scalar amount passed as an argument
// Usage: v.Mult(val)
func (v *Vector) Mult(s interface{}) Vector {
	m := getFloat64(s)
	v.X *= m
	v.Y *= m
	v.Z *= m
	return *v
}

// Div (divide) modifies a vector variable by dividing the x, y, and z values
// by a scalar amount passed as an argument
// Usage: v.Div(val)
func (v *Vector) Div(s interface{}) Vector {
	m := getFloat64(s)
	v.X /= m
	v.Y /= m
	v.Z /= m
	return *v
}

// Add returns a vector equivalent to v1.Copy().Add(v2)
// Usage: vec := vector.Add(v1, v2)
func Add(v1, v2 Vector) Vector {
	r := v1.Copy()
	r.Add(v2)
	return Vector{r.X, r.Y, r.Z}
}

// Sub (subtract) returns a vector equivalent to v1.Copy().Sub(v2)
// Usage: vec := vector.Sub(v1, v2)
func Sub(v1, v2 Vector) Vector {
	r := v1.Copy()
	r.Sub(v2)
	return Vector{r.X, r.Y, r.Z}
}

// Mag (magnitude) returns the length of the vector
// Aka distance from origin
// Usage: val := v.Mag()
func (v Vector) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// MagSq returns the length of the vector squared
// Faster than Mag() because there's no square root calculation
// Usage: val := v.MagSq()
func (v Vector) MagSq() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Copy returns a copy of a vector without modifying the initial vector
// Usage: vec := v.Copy()
func (v Vector) Copy() Vector {
	return Vector{v.X, v.Y, v.Z}
}

// Get returns a copy of a vector without modifying the initial vector
// Usage: vec := v.Get()
func (v Vector) Get() Vector {
	return v.Copy()
}

// Normalize sets the length (magnitude) of a vector to 1
// By dividing the vector by its magnitude
// Usage: v.Normalize()
func (v *Vector) Normalize() Vector {
	v.Div(v.Mag())
	return *v
}

// Dist (distance) returns a scalar distance from one vector to another
// Usage: val := v.Dist(v1)
func (v Vector) Dist(v1 Vector) float64 {
	dx := v.X - v1.X
	dy := v.Y - v1.Y
	dz := v.Z - v1.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// DistSq returns the square of the distance from one vector to another
// Faster than Dist() becaose there's no square root calculation
// Usage: val := v.DistSq(v1)
func (v Vector) DistSq(v1 Vector) float64 {
	dx := v.X - v1.X
	dy := v.Y - v1.Y
	dz := v.Z - v1.Z
	return dx*dx + dy*dy + dz*dz
}

// Dot evaluates and returns the dot product of two vectors
// Usage: val := v.Dot(v1)
func (v Vector) Dot(v1 Vector) float64 {
	return v.X*v1.X + v.Y*v1.Y + v.Z*v1.Z
}

// Cross evaluates and returns the vector cross product of two vectors in 3D space
// Usage: vec := v.Cross(v1)
func (v Vector) Cross(v1 Vector) Vector {
	cx := v.Y*v1.Z - v1.Y*v.Z
	cy := v.Z*v1.X - v1.Z*v.X
	cz := v.X*v1.Y - v1.X*v.Y
	return Vector{cx, cy, cz}
}

// Limit limits the magnitude of a vector to the argument value
// If the magnitude of the vector is over the limit, set it to the limit
// Usage: v.Limit(val)
func (v *Vector) Limit(max interface{}) Vector {
	m := getFloat64(max)
	if v.MagSq() > m*m {
		v.Normalize()
		v.Mult(m)
	}
	return *v
}

// SetMag sets the magnitude of a vector to the argument value
// Usage: v.SetMag(val)
func (v *Vector) SetMag(mag interface{}) Vector {
	m := getFloat64(mag)
	v.Normalize()
	v.Mult(m)
	return *v
}

// Heading returns the angle of a vector in 2D space
// Opposite sister function of FromAngle()
// Usage: val := v.Heading()
func (v Vector) Heading() float64 {
	return math.Atan2(v.Y, v.X)
}

// Rotate rotates a vector by a given angle
// Usage: v.Rotate(amt)
func (v *Vector) Rotate(amt interface{}) Vector {
	t := v.X
	a := getFloat64(amt)
	v.X = v.X*math.Cos(a) - v.Y*math.Sin(a)
	v.Y = t*math.Sin(a) - v.X*math.Cos(a)
	return *v
}

// Lerp returns a vector that is the interpolation between two vectors
// at a certain percentage (amt)
// Usage: vec := vector.Lerp(v1, v2, amt)
func Lerp(v1, v2 Vector, amt interface{}) Vector {
	a := getFloat64(amt)
	x := lerp(v1.X, v2.X, a)
	y := lerp(v1.Y, v2.Y, a)
	z := lerp(v1.Z, v2.Z, a)
	return Vector{x, y, z}
}

// AngleBetween calculates the angle between two vectors in 3D space and
// returns that angle
// If the magnitude of any vector is 0, it returns 0 because there is no defined angle
// Constrains the angle between 0 and PI
// Usage: val := vector.AngleBetween(v1, v2)
func AngleBetween(v1, v2 Vector) float64 {
	if v1.X == 0 && v1.Y == 0 && v1.Z == 0 {
		return 0
	}
	if v2.X == 0 && v2.Y == 0 && v2.Z == 0 {
		return 0
	}
	dot := v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
	v1mag := math.Sqrt(v1.X*v1.X + v1.Y*v1.Y + v1.Z*v1.Z)
	v2mag := math.Sqrt(v2.X*v2.X + v2.Y*v2.Y + v2.Z*v2.Z)
	amt := dot / (v1mag * v2mag)
	if amt <= -1 {
		return 0
	} else if amt >= 1 {
		return math.Pi
	}
	return math.Acos(amt)
}

// Converts to float64 from an unknown interface{} type
// Non-exported as it's only needed within the library and doesn't directly
// interface with vectors
func getFloat64(unknown interface{}) float64 {
	switch i := unknown.(type) {
	case float64:
		return i
	case float32:
		return float64(i)
	case int64:
		return float64(i)
	case int32:
		return float64(i)
	case int:
		return float64(i)
	case uint64:
		return float64(i)
	case uint32:
		return float64(i)
	case uint:
		return float64(i)
	case string:
		f, _ := strconv.ParseFloat(i, 64)
		return f
	default:
		v := reflect.ValueOf(unknown)
		v = reflect.Indirect(v)
		if v.Type().ConvertibleTo(reflect.TypeOf(float64(0))) {
			fv := v.Convert(reflect.TypeOf(float64(0)))
			return fv.Float()
		} else if v.Type().ConvertibleTo(reflect.TypeOf("")) {
			sv := v.Convert(reflect.TypeOf(""))
			s := sv.String()
			f, _ := strconv.ParseFloat(s, 64)
			return f
		} else {
			return math.NaN()
		}
	}
}

// Lerp interpolates between two floating points at a percentage amt
// Needed for Lerp(v1, v2, amt)
func lerp(a, b, amt float64) float64 {
	return a + (b-a)*amt
}
