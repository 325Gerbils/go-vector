package vector

import (
	"math"
	"math/rand"
	"reflect"
	"strconv"
)

// Vector x y z
type Vector struct {
	x, y, z float64
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
	v.x += v1.x
	v.y += v1.y
	v.z += v1.z
}

// Sub v1 from v
func (v *Vector) Sub(v1 Vector) {
	v.x -= v1.x
	v.y -= v1.y
	v.z -= v1.z
}

// Mult v by s
func (v *Vector) Mult(s interface{}) {
	m := getFloat64(s)
	v.x *= m
	v.y *= m
	v.z *= m
}

// Div v by s
func (v *Vector) Div(s interface{}) {
	m := getFloat64(s)
	v.x /= m
	v.y /= m
	v.z /= m
}

// Add v1 and v2
func Add(v1 Vector, v2 Vector) Vector {
	r := v1.Copy()
	r.Add(v2)
	return Vector{r.x, r.y, r.z}
}

// Sub v2 from v1
func Sub(v1 Vector, v2 Vector) Vector {
	r := v1.Copy()
	r.Sub(v2)
	return Vector{r.x, r.y, r.z}
}

// Mag float64
func (v Vector) Mag() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

// MagSq float64
func (v Vector) MagSq() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

// Copy Vector
func (v Vector) Copy() Vector {
	return Vector{v.x, v.y, v.z}
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
	dx := v.x - v1.x
	dy := v.y - v1.y
	dz := v.z - v1.z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// DistSq float64
func (v Vector) DistSq(v1 Vector) float64 {
	dx := v.x - v1.x
	dy := v.y - v1.y
	dz := v.z - v1.z
	return dx*dx + dy*dy + dz*dz
}

// Dot float64
func (v Vector) Dot(v1 Vector) float64 {
	return v.x*v1.x + v.y*v1.y + v.z*v1.z
}

// Cross Vector
func (v Vector) Cross(v1 Vector) Vector {
	cx := v.y*v1.z - v1.y*v.z
	cy := v.z*v1.x - v1.z*v.x
	cz := v.x*v1.y - v1.x*v.y
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
	return math.Atan2(v.y, v.x)
}

// Rotate v
func (v *Vector) Rotate(amt interface{}) {
	t := v.x
	a := getFloat64(amt)
	v.x = v.x*math.Cos(a) - v.y*math.Sin(a)
	v.y = t*math.Sin(a) - v.x*math.Cos(a)
}

// Lerp Vector
func Lerp(v, v1 Vector, amt interface{}) Vector {
	a := getFloat64(amt)
	x := lerp(v.x, v1.x, a)
	y := lerp(v.y, v1.y, a)
	z := lerp(v.z, v1.z, a)
	return Vector{x, y, z}
}

// AngleBetween float64
func AngleBetween(v, v1 Vector) float64 {
	if v.x == 0 && v.y == 0 && v.z == 0 {
		return 0
	}
	if v1.x == 0 && v1.y == 0 && v1.z == 0 {
		return 0
	}
	dot := v.x*v1.x + v.y*v1.y + v.z*v1.z
	vmag := math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
	v1mag := math.Sqrt(v1.x*v1.x + v1.y*v1.y + v1.z*v1.z)
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
