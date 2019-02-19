# go-vector
go-vector provides an easy to use library for creating and modifying vectors in Golang. 

[Skip to Examples](#Examples)

[Skip to Documentation](#Documentation)

A Vector is a struct type with associated methods and functions that either operate on existing Vector structs, return new Vector structs, or calculate properties of Vector structs

```go
type Vector struct {
	X, Y, Z float64
}
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

Todo. 
