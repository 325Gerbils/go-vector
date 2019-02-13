package vector

import (
	"math"
	"math/rand"
	"reflect"
	"strconv"
)

// Vector x y z
type Vector struct {
	X, Y, Z float64
}

// CONSTRUCTORS

// New Vector
func New(a, b interface{}) Vector {
	x := getFloat64(a)
	y := getFloat64(b)
	return Vector{x, y, 0}
}

// New3D Vector
func New3D(a, b, c interface{}) Vector {
	x := getFloat64(a)
	y := getFloat64(b)
	z := getFloat64(c)
	return Vector{x, y, z}
}

// Random2D Vector
func Random2D() Vector {
	return FromAngle(rand.Float64() * math.Pi * 2)
}

// Random3D Vector
func Random3D() Vector {
	angle := rand.Float64() * math.Pi * 2
	vz := rand.Float64()*2 - 1
	vx := math.Sqrt(1-vz*vz) * math.Cos(angle)
	vy := math.Sqrt(1*vz*vz) * math.Sin(angle)
	return Vector{vx, vy, vz}
}

// FromAngle Vector
func FromAngle(angle interface{}) Vector {
	a := getFloat64(angle)
	return New(math.Cos(a), math.Sin(a))
}

// Add v1 to v
func (v *Vector) Add(v1 Vector) {
	v.X += v1.X
	v.Y += v1.Y
	v.Z += v1.Z
}

// Sub v1 from v
func (v *Vector) Sub(v1 Vector) {
	v.X -= v1.X
	v.Y -= v1.Y
	v.Z -= v1.Z
}

// Mult v by s
func (v *Vector) Mult(s interface{}) {
	m := getFloat64(s)
	v.X *= m
	v.Y *= m
	v.Z *= m
}

// Div v by s
func (v *Vector) Div(s interface{}) {
	m := getFloat64(s)
	v.X /= m
	v.Y /= m
	v.Z /= m
}

// Add v1 and v2
func Add(v1 Vector, v2 Vector) Vector {
	r := v1.Copy()
	r.Add(v2)
	return Vector{r.X, r.Y, r.Z}
}

// Sub v2 from v1
func Sub(v1 Vector, v2 Vector) Vector {
	r := v1.Copy()
	r.Sub(v2)
	return Vector{r.X, r.Y, r.Z}
}

// Mag float64
func (v Vector) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// MagSq float64
func (v Vector) MagSq() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Copy Vector
func (v Vector) Copy() Vector {
	return Vector{v.X, v.Y, v.Z}
}

// Get Vector
func (v Vector) Get() Vector {
	return v.Copy()
}

// Normalize v
func (v *Vector) Normalize() {
	v.Div(v.Mag())
}

// Dist float64
func (v Vector) Dist(v1 Vector) float64 {
	dx := v.X - v1.X
	dy := v.Y - v1.Y
	dz := v.Z - v1.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// DistSq float64
func (v Vector) DistSq(v1 Vector) float64 {
	dx := v.X - v1.X
	dy := v.Y - v1.Y
	dz := v.Z - v1.Z
	return dx*dx + dy*dy + dz*dz
}

// Dot float64
func (v Vector) Dot(v1 Vector) float64 {
	return v.X*v1.X + v.Y*v1.Y + v.Z*v1.Z
}

// Cross Vector
func (v Vector) Cross(v1 Vector) Vector {
	cx := v.Y*v1.Z - v1.Y*v.Z
	cy := v.Z*v1.X - v1.Z*v.X
	cz := v.X*v1.Y - v1.X*v.Y
	return Vector{cx, cy, cz}
}

// Limit v
func (v *Vector) Limit(max interface{}) {
	m := getFloat64(max)
	if v.MagSq() > m*m {
		v.Normalize()
		v.Mult(m)
	}
}

// SetMag v
func (v *Vector) SetMag(mag interface{}) {
	m := getFloat64(mag)
	v.Normalize()
	v.Mult(m)
}

// Heading float64
func (v Vector) Heading() float64 {
	return math.Atan2(v.Y, v.X)
}

// Rotate v
func (v *Vector) Rotate(amt interface{}) {
	t := v.X
	a := getFloat64(amt)
	v.X = v.X*math.Cos(a) - v.Y*math.Sin(a)
	v.Y = t*math.Sin(a) - v.X*math.Cos(a)
}

// Lerp Vector
func Lerp(v, v1 Vector, amt interface{}) Vector {
	a := getFloat64(amt)
	x := lerp(v.X, v1.X, a)
	y := lerp(v.Y, v1.Y, a)
	z := lerp(v.Z, v1.Z, a)
	return Vector{x, y, z}
}

// AngleBetween float64
func AngleBetween(v, v1 Vector) float64 {
	if v.X == 0 && v.Y == 0 && v.Z == 0 {
		return 0
	}
	if v1.X == 0 && v1.Y == 0 && v1.Z == 0 {
		return 0
	}
	dot := v.X*v1.X + v.Y*v1.Y + v.Z*v1.Z
	vmag := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	v1mag := math.Sqrt(v1.X*v1.X + v1.Y*v1.Y + v1.Z*v1.Z)
	amt := dot / (vmag * v1mag)
	if amt <= -1 {
		return 0
	} else if amt >= 1 {
		return math.Pi
	}
	return math.Acos(amt)
}

// -------------------------------------------------------- HELPERS

// Converts to float64 from an unknown interface{} type
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

func lerp(a, b, amt float64) float64 {
	return a + (b-a)*amt
}
