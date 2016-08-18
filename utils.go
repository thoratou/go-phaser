package main

import (
	"fmt"
	"strings"
	"unicode"
)

func LowerInitial(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func UpperInitial(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func TypeNoNamespace(str string) string {
	str = strings.TrimPrefix(str, "Phaser.")
	str = strings.TrimPrefix(str, "PIXI.")
	/*lastPoint := strings.LastIndex(str, ".")
	if lastPoint != -1 {
		return str[lastPoint+1 : len(str)]
	}*/
	str = strings.Replace(str, ".", "", -1)
	str = UpperInitial(str)
	return str
}

func SplitMultilines(str string) []string {
	lines := strings.Replace(str, "\r", "\n", -1)
	return strings.Split(lines, "\n")
}

func GoNativeType(str string) string {
	switch str {
	case "boolean":
		return "bool"
	case "Boolean":
		return "bool"
	case "bool":
		return "bool"
	case "integer":
		return "int"
	case "Integer":
		return "int"
	case "int":
		return "int"
	case "number":
		return "float64"
	case "Number":
		return "float64"
	case "float64":
		return "float64"
	case "Bounds-like":
		return "float64"
	case "string":
		return "string"
	case "String":
		return "string"
	default:
		return ""
	}
}

func GoType(str string) string {
	if native := GoNativeType(str); native != "" {
		return native
	}
	switch str {
	case "any":
		return "interface{}"
	case "object":
		return "interface{}"
	case "*":
		return "interface{}"
	case "function":
		return "func(...interface{})"
	case "array":
		return "[]interface{}"
	case "Array":
		return "[]interface{}"
	case "Array.<any>":
		return "[]interface{}"
	case "Array.<Array.<any>>":
		return "[][]interface{}"
	default:
		if strings.HasPrefix(str, "Array.<") {
			typeInArray := strings.TrimPrefix(str, "Array.<")
			typeInArray = strings.TrimSuffix(typeInArray, ">")
			if native := GoNativeType(typeInArray); native != "" {
				return "[]" + native
			}
			return "[]" + TypeNoNamespace(typeInArray)
		}
		if strings.HasPrefix(str, "array.<") {
			typeInArray := strings.TrimPrefix(str, "array.<")
			typeInArray = strings.TrimSuffix(typeInArray, ">")
			if native := GoNativeType(typeInArray); native != "" {
				return "[]" + native
			}
			return "[]" + TypeNoNamespace(typeInArray)
		}
		return TypeNoNamespace(str)
	}
}

func GoTypeInArray(str string) string {
	if fullType := GoType(str); strings.HasPrefix(fullType, "[]") {
		return strings.TrimPrefix(fullType, "[]")
	}
	return ""
}

func GoTypeNativeInArray(str string) string {
	if fullType := GoType(str); strings.HasPrefix(fullType, "[]") {
		typeInArray := strings.TrimPrefix(fullType, "[]")
		return GoNativeType(typeInArray)
	}
	return ""
}

func RemoveDuplicateMembers(elements []Member, className string) []Member {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []Member{}

	for _, v := range elements {
		if encountered[v.Name] == true {
			// Do not add duplicate.
			fmt.Println("warning: remove duplicate member: " + v.Name + " in class: " + className)
		} else {
			// Record this element as an encountered element.
			encountered[v.Name] = true
			// Append to result slice.
			result = append(result, v)
		}
	}
	// Return the new slice.
	return result
}

func RemoveDuplicateFunctions(elements []Function, className string) []Function {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []Function{}

	for _, v := range elements {
		if encountered[v.Name] == true {
			// Do not add duplicate.
			fmt.Println("warning: remove duplicate function: " + v.Name + " in class: " + className)
		} else {
			// Record this element as an encountered element.
			encountered[v.Name] = true
			// Append to result slice.
			result = append(result, v)
		}
	}
	// Return the new slice.
	return result
}
