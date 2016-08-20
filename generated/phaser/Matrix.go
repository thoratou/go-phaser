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
func (self *Matrix) GetA() float64{
    return self.Get("a").Float()
}

// 
func (self *Matrix) SetA(member float64) {
    self.Set("a", member)
}

// 
func (self *Matrix) GetB() float64{
    return self.Get("b").Float()
}

// 
func (self *Matrix) SetB(member float64) {
    self.Set("b", member)
}

// 
func (self *Matrix) GetC() float64{
    return self.Get("c").Float()
}

// 
func (self *Matrix) SetC(member float64) {
    self.Set("c", member)
}

// 
func (self *Matrix) GetD() float64{
    return self.Get("d").Float()
}

// 
func (self *Matrix) SetD(member float64) {
    self.Set("d", member)
}

// 
func (self *Matrix) GetTx() float64{
    return self.Get("tx").Float()
}

// 
func (self *Matrix) SetTx(member float64) {
    self.Set("tx", member)
}

// 
func (self *Matrix) GetTy() float64{
    return self.Get("ty").Float()
}

// 
func (self *Matrix) SetTy(member float64) {
    self.Set("ty", member)
}

// The const type of this object.
func (self *Matrix) GetType() float64{
    return self.Get("type").Float()
}

// The const type of this object.
func (self *Matrix) SetType(member float64) {
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
func (self *Matrix) FromArrayI(args ...interface{}) Matrix{
    return Matrix{self.Call("fromArray", args)}
}

// Sets the values of this Matrix to the given values.
func (self *Matrix) SetToI(args ...interface{}) Matrix{
    return Matrix{self.Call("setTo", args)}
}

// Creates a new Matrix object based on the values of this Matrix.
// If you provide the output parameter the values of this Matrix will be copied over to it.
// If the output parameter is blank a new Matrix object will be created.
func (self *Matrix) CloneI(args ...interface{}) Matrix{
    return Matrix{self.Call("clone", args)}
}

// Copies the properties from this Matrix to the given Matrix.
func (self *Matrix) CopyToI(args ...interface{}) Matrix{
    return Matrix{self.Call("copyTo", args)}
}

// Copies the properties from the given Matrix into this Matrix.
func (self *Matrix) CopyFromI(args ...interface{}) Matrix{
    return Matrix{self.Call("copyFrom", args)}
}

// Creates a Float32 Array with values populated from this Matrix object.
func (self *Matrix) ToArrayI(args ...interface{}) Float32Array{
    return Float32Array{self.Call("toArray", args)}
}

// Get a new position with the current transformation applied.
// 
// Can be used to go from a childs coordinate space to the world coordinate space (e.g. rendering)
func (self *Matrix) ApplyI(args ...interface{}) Point{
    return Point{self.Call("apply", args)}
}

// Get a new position with the inverse of the current transformation applied.
// 
// Can be used to go from the world coordinate space to a childs coordinate space. (e.g. input)
func (self *Matrix) ApplyInverseI(args ...interface{}) Point{
    return Point{self.Call("applyInverse", args)}
}

// Translates the matrix on the x and y.
// This is the same as Matrix.tx += x.
func (self *Matrix) TranslateI(args ...interface{}) Matrix{
    return Matrix{self.Call("translate", args)}
}

// Applies a scale transformation to this matrix.
func (self *Matrix) ScaleI(args ...interface{}) Matrix{
    return Matrix{self.Call("scale", args)}
}

// Applies a rotation transformation to this matrix.
func (self *Matrix) RotateI(args ...interface{}) Matrix{
    return Matrix{self.Call("rotate", args)}
}

// Appends the given Matrix to this Matrix.
func (self *Matrix) AppendI(args ...interface{}) Matrix{
    return Matrix{self.Call("append", args)}
}

// Resets this Matrix to an identity (default) matrix.
func (self *Matrix) IdentityI(args ...interface{}) Matrix{
    return Matrix{self.Call("identity", args)}
}
