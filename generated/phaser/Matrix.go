// Package phaser Automatic generation for Phaser.Matrix
// generated file Matrix.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Matrix The Matrix is a 3x3 matrix mostly used for display transforms within the renderer.
// 
// It is represented like so:
// 
// | a | b | tx |
// | c | d | ty |
// | 0 | 0 | 1 |
type Matrix struct {
    *js.Object
}

// NewMatrix The Matrix is a 3x3 matrix mostly used for display transforms within the renderer.
// 
// It is represented like so:
// 
// | a | b | tx |
// | c | d | ty |
// | 0 | 0 | 1 |
func NewMatrix() *Matrix {
    return &Matrix{js.Global.Get("Phaser").Get("Matrix").New()}
}
// NewMatrix1O The Matrix is a 3x3 matrix mostly used for display transforms within the renderer.
// 
// It is represented like so:
// 
// | a | b | tx |
// | c | d | ty |
// | 0 | 0 | 1 |
func NewMatrix1O(a int) *Matrix {
    return &Matrix{js.Global.Get("Phaser").Get("Matrix").New(a)}
}
// NewMatrix2O The Matrix is a 3x3 matrix mostly used for display transforms within the renderer.
// 
// It is represented like so:
// 
// | a | b | tx |
// | c | d | ty |
// | 0 | 0 | 1 |
func NewMatrix2O(a int, b int) *Matrix {
    return &Matrix{js.Global.Get("Phaser").Get("Matrix").New(a, b)}
}
// NewMatrix3O The Matrix is a 3x3 matrix mostly used for display transforms within the renderer.
// 
// It is represented like so:
// 
// | a | b | tx |
// | c | d | ty |
// | 0 | 0 | 1 |
func NewMatrix3O(a int, b int, c int) *Matrix {
    return &Matrix{js.Global.Get("Phaser").Get("Matrix").New(a, b, c)}
}
// NewMatrix4O The Matrix is a 3x3 matrix mostly used for display transforms within the renderer.
// 
// It is represented like so:
// 
// | a | b | tx |
// | c | d | ty |
// | 0 | 0 | 1 |
func NewMatrix4O(a int, b int, c int, d int) *Matrix {
    return &Matrix{js.Global.Get("Phaser").Get("Matrix").New(a, b, c, d)}
}
// NewMatrix5O The Matrix is a 3x3 matrix mostly used for display transforms within the renderer.
// 
// It is represented like so:
// 
// | a | b | tx |
// | c | d | ty |
// | 0 | 0 | 1 |
func NewMatrix5O(a int, b int, c int, d int, tx int) *Matrix {
    return &Matrix{js.Global.Get("Phaser").Get("Matrix").New(a, b, c, d, tx)}
}
// NewMatrix6O The Matrix is a 3x3 matrix mostly used for display transforms within the renderer.
// 
// It is represented like so:
// 
// | a | b | tx |
// | c | d | ty |
// | 0 | 0 | 1 |
func NewMatrix6O(a int, b int, c int, d int, tx int, ty int) *Matrix {
    return &Matrix{js.Global.Get("Phaser").Get("Matrix").New(a, b, c, d, tx, ty)}
}
// NewMatrixI The Matrix is a 3x3 matrix mostly used for display transforms within the renderer.
// 
// It is represented like so:
// 
// | a | b | tx |
// | c | d | ty |
// | 0 | 0 | 1 |
func NewMatrixI(args ...interface{}) *Matrix {
    return &Matrix{js.Global.Get("Phaser").Get("Matrix").New(args)}
}



// Matrix Binding conversion method to Matrix point 
func ToMatrix(jsStruct interface{}) *Matrix {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Matrix{Object: object}
	}
	return nil
}



// A empty description
func (self *Matrix) A() int{
    return self.Object.Get("a").Int()
}

// SetAA empty description
func (self *Matrix) SetAA(member int) {
    self.Object.Set("a", member)
}

// B empty description
func (self *Matrix) B() int{
    return self.Object.Get("b").Int()
}

// SetBA empty description
func (self *Matrix) SetBA(member int) {
    self.Object.Set("b", member)
}

// C empty description
func (self *Matrix) C() int{
    return self.Object.Get("c").Int()
}

// SetCA empty description
func (self *Matrix) SetCA(member int) {
    self.Object.Set("c", member)
}

// D empty description
func (self *Matrix) D() int{
    return self.Object.Get("d").Int()
}

// SetDA empty description
func (self *Matrix) SetDA(member int) {
    self.Object.Set("d", member)
}

// Tx empty description
func (self *Matrix) Tx() int{
    return self.Object.Get("tx").Int()
}

// SetTxA empty description
func (self *Matrix) SetTxA(member int) {
    self.Object.Set("tx", member)
}

// Ty empty description
func (self *Matrix) Ty() int{
    return self.Object.Get("ty").Int()
}

// SetTyA empty description
func (self *Matrix) SetTyA(member int) {
    self.Object.Set("ty", member)
}

// Type The const type of this object.
func (self *Matrix) Type() int{
    return self.Object.Get("type").Int()
}

// SetTypeA The const type of this object.
func (self *Matrix) SetTypeA(member int) {
    self.Object.Set("type", member)
}


// FromArray Sets the values of this Matrix to the values in the given array.
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

// FromArrayI Sets the values of this Matrix to the values in the given array.
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

// SetTo Sets the values of this Matrix to the given values.
func (self *Matrix) SetTo(a int, b int, c int, d int, tx int, ty int) *Matrix{
    return &Matrix{self.Object.Call("setTo", a, b, c, d, tx, ty)}
}

// SetToI Sets the values of this Matrix to the given values.
func (self *Matrix) SetToI(args ...interface{}) *Matrix{
    return &Matrix{self.Object.Call("setTo", args)}
}

// Clone Creates a new Matrix object based on the values of this Matrix.
// If you provide the output parameter the values of this Matrix will be copied over to it.
// If the output parameter is blank a new Matrix object will be created.
func (self *Matrix) Clone() *Matrix{
    return &Matrix{self.Object.Call("clone")}
}

// Clone1O Creates a new Matrix object based on the values of this Matrix.
// If you provide the output parameter the values of this Matrix will be copied over to it.
// If the output parameter is blank a new Matrix object will be created.
func (self *Matrix) Clone1O(output *Matrix) *Matrix{
    return &Matrix{self.Object.Call("clone", output)}
}

// CloneI Creates a new Matrix object based on the values of this Matrix.
// If you provide the output parameter the values of this Matrix will be copied over to it.
// If the output parameter is blank a new Matrix object will be created.
func (self *Matrix) CloneI(args ...interface{}) *Matrix{
    return &Matrix{self.Object.Call("clone", args)}
}

// CopyTo Copies the properties from this Matrix to the given Matrix.
func (self *Matrix) CopyTo(matrix *Matrix) *Matrix{
    return &Matrix{self.Object.Call("copyTo", matrix)}
}

// CopyToI Copies the properties from this Matrix to the given Matrix.
func (self *Matrix) CopyToI(args ...interface{}) *Matrix{
    return &Matrix{self.Object.Call("copyTo", args)}
}

// CopyFrom Copies the properties from the given Matrix into this Matrix.
func (self *Matrix) CopyFrom(matrix *Matrix) *Matrix{
    return &Matrix{self.Object.Call("copyFrom", matrix)}
}

// CopyFromI Copies the properties from the given Matrix into this Matrix.
func (self *Matrix) CopyFromI(args ...interface{}) *Matrix{
    return &Matrix{self.Object.Call("copyFrom", args)}
}

// ToArray Creates a Float32 Array with values populated from this Matrix object.
func (self *Matrix) ToArray() *Float32Array{
    return &Float32Array{self.Object.Call("toArray")}
}

// ToArray1O Creates a Float32 Array with values populated from this Matrix object.
func (self *Matrix) ToArray1O(transpose bool) *Float32Array{
    return &Float32Array{self.Object.Call("toArray", transpose)}
}

// ToArray2O Creates a Float32 Array with values populated from this Matrix object.
func (self *Matrix) ToArray2O(transpose bool, array *Float32Array) *Float32Array{
    return &Float32Array{self.Object.Call("toArray", transpose, array)}
}

// ToArrayI Creates a Float32 Array with values populated from this Matrix object.
func (self *Matrix) ToArrayI(args ...interface{}) *Float32Array{
    return &Float32Array{self.Object.Call("toArray", args)}
}

// Apply Get a new position with the current transformation applied.
// 
// Can be used to go from a childs coordinate space to the world coordinate space (e.g. rendering)
func (self *Matrix) Apply(pos *Point) *Point{
    return &Point{self.Object.Call("apply", pos)}
}

// Apply1O Get a new position with the current transformation applied.
// 
// Can be used to go from a childs coordinate space to the world coordinate space (e.g. rendering)
func (self *Matrix) Apply1O(pos *Point, newPos *Point) *Point{
    return &Point{self.Object.Call("apply", pos, newPos)}
}

// ApplyI Get a new position with the current transformation applied.
// 
// Can be used to go from a childs coordinate space to the world coordinate space (e.g. rendering)
func (self *Matrix) ApplyI(args ...interface{}) *Point{
    return &Point{self.Object.Call("apply", args)}
}

// ApplyInverse Get a new position with the inverse of the current transformation applied.
// 
// Can be used to go from the world coordinate space to a childs coordinate space. (e.g. input)
func (self *Matrix) ApplyInverse(pos *Point) *Point{
    return &Point{self.Object.Call("applyInverse", pos)}
}

// ApplyInverse1O Get a new position with the inverse of the current transformation applied.
// 
// Can be used to go from the world coordinate space to a childs coordinate space. (e.g. input)
func (self *Matrix) ApplyInverse1O(pos *Point, newPos *Point) *Point{
    return &Point{self.Object.Call("applyInverse", pos, newPos)}
}

// ApplyInverseI Get a new position with the inverse of the current transformation applied.
// 
// Can be used to go from the world coordinate space to a childs coordinate space. (e.g. input)
func (self *Matrix) ApplyInverseI(args ...interface{}) *Point{
    return &Point{self.Object.Call("applyInverse", args)}
}

// Translate Translates the matrix on the x and y.
// This is the same as Matrix.tx += x.
func (self *Matrix) Translate(x int, y int) *Matrix{
    return &Matrix{self.Object.Call("translate", x, y)}
}

// TranslateI Translates the matrix on the x and y.
// This is the same as Matrix.tx += x.
func (self *Matrix) TranslateI(args ...interface{}) *Matrix{
    return &Matrix{self.Object.Call("translate", args)}
}

// Scale Applies a scale transformation to this matrix.
func (self *Matrix) Scale(x int, y int) *Matrix{
    return &Matrix{self.Object.Call("scale", x, y)}
}

// ScaleI Applies a scale transformation to this matrix.
func (self *Matrix) ScaleI(args ...interface{}) *Matrix{
    return &Matrix{self.Object.Call("scale", args)}
}

// Rotate Applies a rotation transformation to this matrix.
func (self *Matrix) Rotate(angle int) *Matrix{
    return &Matrix{self.Object.Call("rotate", angle)}
}

// RotateI Applies a rotation transformation to this matrix.
func (self *Matrix) RotateI(args ...interface{}) *Matrix{
    return &Matrix{self.Object.Call("rotate", args)}
}

// Append Appends the given Matrix to this Matrix.
func (self *Matrix) Append(matrix *Matrix) *Matrix{
    return &Matrix{self.Object.Call("append", matrix)}
}

// AppendI Appends the given Matrix to this Matrix.
func (self *Matrix) AppendI(args ...interface{}) *Matrix{
    return &Matrix{self.Object.Call("append", args)}
}

// Identity Resets this Matrix to an identity (default) matrix.
func (self *Matrix) Identity() *Matrix{
    return &Matrix{self.Object.Call("identity")}
}

// IdentityI Resets this Matrix to an identity (default) matrix.
func (self *Matrix) IdentityI(args ...interface{}) *Matrix{
    return &Matrix{self.Object.Call("identity", args)}
}

