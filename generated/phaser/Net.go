// Automatic generation for Phaser.Net
// generated file Net.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Phaser.Net handles browser URL related tasks such as checking host names, domain names and query string manipulation.
type Net struct {
    *js.Object
}




// Returns the hostname given by the browser.
func (self *Net) GetHostNameI(args ...interface{}) string{
    return self.Call("getHostName", args).String()
}

// Compares the given domain name against the hostname of the browser containing the game.
// If the domain name is found it returns true.
// You can specify a part of a domain, for example 'google' would match 'google.com', 'google.co.uk', etc.
// Do not include 'http://' at the start.
func (self *Net) CheckDomainNameI(args ...interface{}) bool{
    return self.Call("checkDomainName", args).Bool()
}

// Updates a value on the Query String and returns it in full.
// If the value doesn't already exist it is set.
// If the value exists it is replaced with the new value given. If you don't provide a new value it is removed from the query string.
// Optionally you can redirect to the new url, or just return it as a string.
func (self *Net) UpdateQueryStringI(args ...interface{}) string{
    return self.Call("updateQueryString", args).String()
}

// Returns the Query String as an object.
// If you specify a parameter it will return just the value of that parameter, should it exist.
func (self *Net) GetQueryStringI(args ...interface{}) interface{}{
    return self.Call("getQueryString", args)
}

// Takes a Uniform Resource Identifier (URI) component (previously created by encodeURIComponent or by a similar routine) and
// decodes it, replacing \ with spaces in the return. Used internally by the Net classes.
func (self *Net) DecodeURII(args ...interface{}) string{
    return self.Call("decodeURI", args).String()
}
