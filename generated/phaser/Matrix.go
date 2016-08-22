// Automatic generation for Phaser.Matrix
// generated file Matrix.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Matrix is a 3x3 matrix mostly used for display transforms within the renderer.
// 
// It is represented like so:
// 
// | a | b | tx |
// | c | d | ty |
// | 0 | 0 | 1 |
type Matrix struct {
    *js.Object
}


// 
func (self *Matrix) GetAA() int{
    return self.Object.Get("a").Int()
}

// 
func (self *Matrix) SetAA(member int) {
    self.Object.Set("a", member)
}

// 
func (self *Matrix) GetBA() int{
    return self.Object.Get("b").Int()
}

// 
func (self *Matrix) SetBA(member int) {
    self.Object.Set("b", member)
}

// 
func (self *Matrix) GetCA() int{
    return self.Object.Get("c").Int()
}

// 
func (self *Matrix) SetCA(member int) {
    self.Object.Set("c", member)
}

// 
func (self *Matrix) GetDA() int{
    return self.Object.Get("d").Int()
}

// 
func (self *Matrix) SetDA(member int) {
    self.Object.Set("d", member)
}

// 
func (self *Matrix) GetTxA() int{
    return self.Object.Get("tx").Int()
}

// 
func (self *Matrix) SetTxA(member int) {
    self.Object.Set("tx", member)
}

// 
func (self *Matrix) GetTyA() int{
    return self.Object.Get("ty").Int()
}

// 
func (self *Matrix) SetTyA(member int) {
    self.Object.Set("ty", member)
}

// The const type of this object.
func (self *Matrix) GetTypeA() int{
    return self.Object.Get("type").Int()
}

// The const type of this object.
func (self *Matrix) SetTypeA(member int) {
    self.Object.Set("type", member)
}



// Sets the values of this Matrix to the values in the given array.
// 
// The Array elements should be set as follows:
// 
// a = array[0]
// b = array[1]
// c = array[3]
// d = array[4]
// tx = array[2]
// ty = array[5]
func (self *Matrix) FromArray(array []interface{}) *Matrix{
    return &Matrix{self.Object.Call("fromArray", array)}
}

// Sets the values of this Matrix to the values in the given array.
// 
// The Array elements should be set as follows:
// 
// a = array[0]
// b = array[1]
// c = array[3]
// d = array[4]
// tx = array[2]
// ty = array[5]
func (self *Matrix) FromArrayI(args ...interface{}) *Matrix{
    return &Matrix{self.Object.Call("fromArray", args)}
}

// Sets the values of this Matrix to the given values.
func (self *Matrix) SetTo(a int, b int, c int, d int, tx int, ty int) *Matrix{
    return &Matrix{self.Object.Call("setTo", a, b, c, d, tx, ty)}
}

// Sets the values of this Matrix to the given values.
func (self *Matrix) SetToI(args ...interface{}) *Matrix{
    return &Matrix{self.Object.Call("setTo", args)}
}

// Creates a new Matrix object based on the values of this Matrix.
// If you provide the output parameter the values of this Matrix will be copied over to it.
// If the output parameter is blank a new Matrix object will be created.
func (self *Matrix) Clone() *Matrix{
    return &Matrix{self.Object.Call("clone")}
}

// Creates a new Matrix object based on the values of this Matrix.
// If you provide the output parameter the values of this Matrix will be copied over to it.
// If the output parameter is blank a new Matrix object will be created.
func (self *Matrix) Clone1O(output *Matrix) *Matrix{
    return &Matrix{self.Object.Call("clone", output)}
}

// Creates a new Matrix object based on the values of this Matrix.
// If you provide the output parameter the values of this Matrix will be copied over to it.
// If the output parameter is blank a new Matrix object will be created.
func (self *Matrix) CloneI(args ...interface{}) *Matrix{
    return &Matrix{self.Object.Call("clone", args)}
}

// Copies the properties from this Matrix to the given Matrix.
func (self *Matrix) CopyTo(matrix *Matrix) *Matrix{
    return &Matrix{self.Object.Call("copyTo", matrix)}
}

// Copies the properties from this Matrix to the given Matrix.
func (self *Matrix) CopyToI(args ...interface{}) *Matrix{
    return &Matrix{self.Object.Call("copyTo", args)}
}

// Copies the properties from the given Matrix into this Matrix.
func (self *Matrix) CopyFrom(matrix *Matrix) *Matrix{
    return &Matrix{self.Object.Call("copyFrom", matrix)}
}

// Copies the properties from the given Matrix into this Matrix.
func (self *Matrix) CopyFromI(args ...interface{}) *Matrix{
    return &Matrix{self.Object.Call("copyFrom", args)}
}

// Creates a Float32 Array with values populated from this Matrix object.
func (self *Matrix) ToArray() *Float32Array{
    return &Float32Array{self.Object.Call("toArray")}
}

// Creates a Float32 Array with values populated from this Matrix object.
func (self *Matrix) ToArray1O(transpose bool) *Float32Array{
    return &Float32Array{self.Object.Call("toArray", transpose)}
}

// Creates a Float32 Array with values populated from this Matrix object.
func (self *Matrix) ToArray2O(transpose bool, array *Float32Array) *Float32Array{
    return &Float32Array{self.Object.Call("toArray", transpose, array)}
}

// Creates a Float32 Array with values populated from this Matrix object.
func (self *Matrix) ToArrayI(args ...interface{}) *Float32Array{
    return &Float32Array{self.Object.Call("toArray", args)}
}

// Get a new position with the current transformation applied.
// 
// Can be used to go from a childs coordinate space to the world coordinate space (e.g. rendering)
func (self *Matrix) Apply(pos *Point) *Point{
    return &Point{self.Object.Call("apply", pos)}
}

// Get a new position with the current transformation applied.
// 
// Can be used to go from a childs coordinate space to the world coordinate space (e.g. rendering)
func (self *Matrix) Apply1O(pos *Point, newPos *Point) *Point{
    return &Point{self.Object.Call("apply", pos, newPos)}
}

// Get a new position with the current transformation applied.
// 
// Can be used to go from a childs coordinate space to the world coordinate space (e.g. rendering)
func (self *Matrix) ApplyI(args ...interface{}) *Point{
    return &Point{self.Object.Call("apply", args)}
}

// Get a new position with the inverse of the current transformation applied.
// 
// Can be used to go from the world coordinate space to a childs coordinate space. (e.g. input)
func (self *Matrix) ApplyInverse(pos *Point) *Point{
    return &Point{self.Object.Call("applyInverse", pos)}
}

// Get a new position with the inverse of the current transformation applied.
// 
// Can be used to go from the world coordinate space to a childs coordinate space. (e.g. input)
func (self *Matrix) ApplyInverse1O(pos *Point, newPos *Point) *Point{
    return &Point{self.Object.Call("applyInverse", pos, newPos)}
}

// Get a new position with the inverse of the current transformation applied.
// 
// Can be used to go from the world coordinate space to a childs coordinate space. (e.g. input)
func (self *Matrix) ApplyInverseI(args ...interface{}) *Point{
    return &Point{self.Object.Call("applyInverse", args)}
}

// Translates the matrix on the x and y.
// This is the same as Matrix.tx += x.
func (self *Matrix) Translate(x int, y int) *Matrix{
    return &Matrix{self.Object.Call("translate", x, y)}
}

// Translates the matrix on the x and y.
// This is the same as Matrix.tx += x.
func (self *Matrix) TranslateI(args ...interface{}) *Matrix{
    return &Matrix{self.Object.Call("translate", args)}
}

// Applies a scale transformation to this matrix.
func (self *Matrix) Scale(x int, y int) *Matrix{
    return &Matrix{self.Object.Call("scale", x, y)}
}

// Applies a scale transformation to this matrix.
func (self *Matrix) ScaleI(args ...interface{}) *Matrix{
    return &Matrix{self.Object.Call("scale", args)}
}

// Applies a rotation transformation to this matrix.
func (self *Matrix) Rotate(angle int) *Matrix{
    return &Matrix{self.Object.Call("rotate", angle)}
}

// Applies a rotation transformation to this matrix.
func (self *Matrix) RotateI(args ...interface{}) *Matrix{
    return &Matrix{self.Object.Call("rotate", args)}
}

// Appends the given Matrix to this Matrix.
func (self *Matrix) Append(matrix *Matrix) *Matrix{
    return &Matrix{self.Object.Call("append", matrix)}
}

// Appends the given Matrix to this Matrix.
func (self *Matrix) AppendI(args ...interface{}) *Matrix{
    return &Matrix{self.Object.Call("append", args)}
}

// Resets this Matrix to an identity (default) matrix.
func (self *Matrix) Identity() *Matrix{
    return &Matrix{self.Object.Call("identity")}
}

// Resets this Matrix to an identity (default) matrix.
func (self *Matrix) IdentityI(args ...interface{}) *Matrix{
    return &Matrix{self.Object.Call("identity", args)}
}
