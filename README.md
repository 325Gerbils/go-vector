# go-vector
go-vector provides an easy to use library for creating and modifying vectors in Golang. 

[Skip to Examples](#Examples)

[Skip to Documentation](#Documentation)

A Vector is a struct type with associated methods and functions that either operate on existing Vector structs, return new Vector structs, or calculate properties of Vector structs.

```go
type Vector struct {
	X, Y, Z float64
}
```

Basic usage looks something like this:

```go
v1 := vector.New(4,5)    // Creates new Vector{4,5,0}
v2 := vector.Random2D()  // Creates new unit Vector with random orientation
v2.Mult(5)               // Multiplies v2's magnitude by 5
v1.Add(v2)               // Adds v2 to v1, modifying v1
fmt.Println(v1.X, v1.Y)  // Prints v1's X and Y values
fmt.Println(v2.Mag())    // Prints v2's magnitude
```


## Intro to Vectors

Todo.

## Examples

A simple physics simulator that adds up the forces (in this case, just gravity) on an object and calculates its location, velocity, and acceleration
```go
package main

import (
	"fmt"

	vector "github.com/325Gerbils/go-vector"
)

func main() {
	gravity := vector.New(0, -9.8) // force of gravity
	acc := vector.New(0, 0)        // acceleration
	vel := vector.New(0, 0)        // velocity
	loc := vector.New(0, 0)        // location
	for {
		acc.Add(gravity)          // Apply Gravity to acceleration
		vel.Add(acc)              // Add acceleration to velocity
		loc.Add(vel)              // Add velocity to location
		fmt.Println(loc.X, loc.Y) // Print the new location
		acc.Mult(0)               // Reset acceleration every frame
	}
}

```

## Documentation

Most documentation is built in to the code. Here is a list of the types and functions:

```
type Vector struct

func New(a, b interface{}) Vector

func New3D(a, b, c interface{}) Vector

func Random2D() Vector

func Random3D() Vector

func FromAngle(angle interface{}) Vector

func (v *Vector) Add(v1 Vector)

func (v *Vector) Sub(v1 Vector)

func (v *Vector) Mult(s interface{})

func (v *Vector) Div(s interface{}) 

func Add(v1, v2 Vector) Vector

func Sub(v1, v2 Vector) Vector

func (v Vector) Mag() float64

func (v Vector) MagSq() float64

func (v Vector) Copy() Vector

func (v Vector) Get() Vector

func (v *Vector) Normalize()

func (v Vector) Dist(v1 Vector) float64

func (v Vector) DistSq(v1 Vector) float64

func (v Vector) Dot(v1 Vector) float64

func (v Vector) Cross(v1 Vector) Vector

func (v *Vector) Limit(max interface{})

func (v *Vector) SetMag(mag interface{}) 

func (v Vector) Heading() float64

func (v *Vector) Rotate(amt interface{})

func Lerp(v1, v2 Vector, amt interface{}) Vector

func AngleBetween(v1, v2 Vector) float64
```

Todo. 

## Planned features

Return a zero vector
```go
vector.Zero() // Returns Vector{0,0,0}
```

Print vector values
```go
vector.Print()   // prints {x, y, z}
vector.Print2D() // prints {x, y}
```
