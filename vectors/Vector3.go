package vectors3d

import (
	"math"
)

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

/*
	Constructs a new 3D Vector from cartesian coordinates.
*/
func Cartesian(x float64, y float64, z float64) Vector3 {
	return Vector3{
		X: x,
		Y: y,
		Z: z,
	}
}

func (this *Vector3) Clear() {
	this.X = 0
	this.Y = 0
	this.Z = 0
}

/*
	Returns true if the vector is equal to another vector, false if not.
*/
func (this *Vector3) Equals(other *Vector3) bool {
	return (this.X == other.X) && (this.Y == other.Y) && (this.Z == other.Z)
}

/*
	Returns true if the vector is different to another vector, false if not.
*/
func (this *Vector3) NotEqual(other *Vector3) bool {
	return !((this.X == other.X) && (this.Y == other.Y) && (this.Z == other.Z))
}

/*
	Returns the magnitude of the vector.
*/
func (this *Vector3) Magnitude() float64 {
	return math.Sqrt(math.Pow(this.X, 2) + math.Pow(this.Y, 2) + math.Pow(this.Z, 2))
}

/*
	Returns the norm of the vector.
*/
func (this *Vector3) Normalize() Vector3 {
	mag := this.Magnitude()
	if mag == 0 {
		return *this
	}

	return this.Divide_Float64(mag)
}

/*
	Divides the vector by n.
*/
func (this *Vector3) Divide_Float64(n float64) Vector3 {
	return Vector3{
		this.X / n,
		this.Y / n,
		this.Z / n,
	}
}

/*
	Divides the vector by n.
*/
func (this *Vector3) Divide_Int(n int) Vector3 {
	return Vector3{
		this.X / float64(n),
		this.Y / float64(n),
		this.Z / float64(n),
	}
}

/*
	Substracts the vector by another vector.
*/
func (this *Vector3) Substract_Vector(n float64) Vector3 {
	return Vector3{
		this.X - n,
		this.Y - n,
		this.Z - n,
	}
}

/*
	Substracts the vector by n.
*/
func (this *Vector3) Substract_Float64(n float64) Vector3 {
	return Vector3{
		this.X - n,
		this.Y - n,
		this.Z - n,
	}
}

/*
	Substracts the vector by n.
*/
func (this *Vector3) Substract_Int(n int) Vector3 {
	return Vector3{
		this.X - float64(n),
		this.Y - float64(n),
		this.Z - float64(n),
	}
}

/*
	Adds the vector by n.
*/
func (this *Vector3) Add_Float64(n float64) Vector3 {
	return Vector3{
		this.X + n,
		this.Y + n,
		this.Z + n,
	}
}

/*
	Adds the vector by n.
*/
func (this *Vector3) Add_Int(n int) Vector3 {
	return Vector3{
		this.X + float64(n),
		this.Y + float64(n),
		this.Z + float64(n),
	}
}

/*
	Adds the vector with another one.
*/
func (this *Vector3) Add_Vector(other *Vector3) Vector3 {
	return Vector3{
		this.X + other.X,
		this.Y + other.Y,
		this.Z + other.Z,
	}
}

/*
	Multiplies the vector by an integer n.
*/
func (this *Vector3) Multiply_Int(n int) Vector3 {
	return Vector3{
		this.X * float64(n),
		this.Y * float64(n),
		this.Z * float64(n),
	}
}

/*
	Sets the vector equal to an array of floats.
*/
func (this *Vector3) Copy_Float64Array(array []float64) {
	this.X = array[0]
	this.Y = array[1]
	this.Y = array[2]
}

/*
	Sets the vector equal to an array of integers.
*/
func (this *Vector3) Copy_IntArray(array []int) {
	this.X = float64(array[0])
	this.Y = float64(array[1])
	this.Z = float64(array[2])
}

/*
	Sets the vector equal to another vector.
*/
func (this *Vector3) Copy_Vector(other *Vector3) {
	this.X = other.X
	this.Y = other.Y
	this.Z = other.Z
}
