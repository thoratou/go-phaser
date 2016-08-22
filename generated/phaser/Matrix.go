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
func (self *Matrix) GetA() int{
    return self.Get("a").Int()
}

// 
func (self *Matrix) SetA(member int) {
    self.Set("a", member)
}

// 
func (self *Matrix) GetB() int{
    return self.Get("b").Int()
}

// 
func (self *Matrix) SetB(member int) {
    self.Set("b", member)
}

// 
func (self *Matrix) GetC() int{
    return self.Get("c").Int()
}

// 
func (self *Matrix) SetC(member int) {
    self.Set("c", member)
}

// 
func (self *Matrix) GetD() int{
    return self.Get("d").Int()
}

// 
func (self *Matrix) SetD(member int) {
    self.Set("d", member)
}

// 
func (self *Matrix) GetTx() int{
    return self.Get("tx").Int()
}

// 
func (self *Matrix) SetTx(member int) {
    self.Set("tx", member)
}

// 
func (self *Matrix) GetTy() int{
    return self.Get("ty").Int()
}

// 
func (self *Matrix) SetTy(member int) {
    self.Set("ty", member)
}

// The const type of this object.
func (self *Matrix) GetType() int{
    return self.Get("type").Int()
}

// The const type of this object.
func (self *Matrix) SetType(member int) {
    self.Set("type", member)
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
    return &Matrix{self.Call("fromArray", args)}
}

// Sets the values of this Matrix to the given values.
func (self *Matrix) SetToI(args ...interface{}) *Matrix{
    return &Matrix{self.Call("setTo", args)}
}

// Creates a new Matrix object based on the values of this Matrix.
// If you provide the output parameter the values of this Matrix will be copied over to it.
// If the output parameter is blank a new Matrix object will be created.
func (self *Matrix) CloneI(args ...interface{}) *Matrix{
    return &Matrix{self.Call("clone", args)}
}

// Copies the properties from this Matrix to the given Matrix.
func (self *Matrix) CopyToI(args ...interface{}) *Matrix{
    return &Matrix{self.Call("copyTo", args)}
}

// Copies the properties from the given Matrix into this Matrix.
func (self *Matrix) CopyFromI(args ...interface{}) *Matrix{
    return &Matrix{self.Call("copyFrom", args)}
}

// Creates a Float32 Array with values populated from this Matrix object.
func (self *Matrix) ToArrayI(args ...interface{}) *Float32Array{
    return &Float32Array{self.Call("toArray", args)}
}

// Get a new position with the current transformation applied.
// 
// Can be used to go from a childs coordinate space to the world coordinate space (e.g. rendering)
func (self *Matrix) ApplyI(args ...interface{}) *Point{
    return &Point{self.Call("apply", args)}
}

// Get a new position with the inverse of the current transformation applied.
// 
// Can be used to go from the world coordinate space to a childs coordinate space. (e.g. input)
func (self *Matrix) ApplyInverseI(args ...interface{}) *Point{
    return &Point{self.Call("applyInverse", args)}
}

// Translates the matrix on the x and y.
// This is the same as Matrix.tx += x.
func (self *Matrix) TranslateI(args ...interface{}) *Matrix{
    return &Matrix{self.Call("translate", args)}
}

// Applies a scale transformation to this matrix.
func (self *Matrix) ScaleI(args ...interface{}) *Matrix{
    return &Matrix{self.Call("scale", args)}
}

// Applies a rotation transformation to this matrix.
func (self *Matrix) RotateI(args ...interface{}) *Matrix{
    return &Matrix{self.Call("rotate", args)}
}

// Appends the given Matrix to this Matrix.
func (self *Matrix) AppendI(args ...interface{}) *Matrix{
    return &Matrix{self.Call("append", args)}
}

// Resets this Matrix to an identity (default) matrix.
func (self *Matrix) IdentityI(args ...interface{}) *Matrix{
    return &Matrix{self.Call("identity", args)}
}
